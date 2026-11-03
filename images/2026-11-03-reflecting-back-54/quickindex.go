// Command quickindex is a tiny, dependency-free full-text search engine for a
// collection of markdown files. It has two modes:
//
//	# build an index from a list of files on stdin (or given as arguments)
//	fd -e md | go run quickindex.go -o my.idx
//	go run quickindex.go -o my.idx content/posts/*.md
//
//	# query it
//	go run quickindex.go -f my.idx "agentic coding"
//	go run quickindex.go -f my.idx -json 'go* wasm'
//
// The index is a JSON document (optionally gzipped, when the file name ends in
// ".gz") holding an inverted index with term positions, so phrase queries and
// result snippets work without the original files.
//
// Relevance is BM25 (k1=1.2, b=0.75) summed over fields, with title matches
// weighted higher than tag matches, and tag matches higher than body matches.
//
// Query syntax: plain words (all of them must match, case insensitive),
// "quoted phrases", and trailing-* prefix terms. Words with no exact match are
// expanded as prefixes automatically (disable with -no-prefix). If the AND of
// all terms matches nothing, the results degrade to the OR of the terms.
//
// Portability / WASM: everything above the "CLI" section at the bottom is pure
// Go - no os, no filesystem, no cgo. AddDoc and Search take bytes and return
// plain structs, so a browser build only needs a thin wrapper, for example:
//
//	//go:build js
//	func main() {
//		js.Global().Set("quickindexSearch", js.FuncOf(func(this js.Value, a []js.Value) any {
//			idx := &Index{}
//			if err := json.Unmarshal([]byte(a[0].String()), idx); err != nil {
//				return "{}"
//			}
//			out, _ := json.Marshal(Search(idx, a[1].String(), SearchOpts{Limit: 10, Prefix: true, RequireAll: true}))
//			return string(out)
//		}))
//		select {}
//	}
//
// Built with:  GOOS=js GOARCH=wasm go build -o quickindex.wasm quickindex.go
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"
)

const (
	indexVersion = 1

	// tokens shorter than this are not indexed (keeps "a", "i", "x" out)
	minIndexToken = 2

	// BM25 parameters
	bm25K1 = 1.2
	bm25B  = 0.75

	// how many index terms a single prefix term may expand to
	defaultMaxExpand = 24
)

// ---------------------------------------------------------------------------
// markdown -> plain text
// ---------------------------------------------------------------------------

var (
	reFenceLine = regexp.MustCompile("(?m)^[ \\t]*(?:\x60{3}|~{3})[^\\n]*$")
	reComment   = regexp.MustCompile(`(?s)<!--.*?-->`)
	reShortcode = regexp.MustCompile(`{{<[^>]*>}}|{{%[^%]*%}}`)
	reImage     = regexp.MustCompile(`!\[([^\]]*)\]\([^)]*\)`)
	reLink      = regexp.MustCompile(`\[([^\]]*)\]\[[^\]]*\]|!\[([^\]]*)\]\([^)]*\)|\[([^\]]*)\]\([^)]*\)`)
	reAngle     = regexp.MustCompile(`</?[a-zA-Z!/][^>\n]*>`)
	reURL       = regexp.MustCompile(`(?i)\b(?:https?|ftp)://[^\s<>)'"\]]+`)
	reHeading   = regexp.MustCompile(`(?m)^ {0,3}#{1,6}[ \t]*`)
	reQuote     = regexp.MustCompile(`(?m)^ {0,3}>+ ?`)
	reBullet    = regexp.MustCompile(`(?m)^[ \t]*(?:[-*+]|\d+[.)])[ \t]+`)
	reTableSep  = regexp.MustCompile(`(?m)^[ \t]*\|?[ \t]*:?-+[ \t|:-]*$`)
	rePlain     = strings.NewReplacer("|", " ", "`", " ", "*", " ", "#", " ", "\\", " ", "—", " ", "–", " ")
)

// markdownToText strips the markdown/Hugo noise that should not be searchable:
// link targets, HTML tags, bare URLs, code fence banners, table rulers. The
// prose - and the body of code samples - stays.
func markdownToText(body string) string {
	s := reComment.ReplaceAllString(body, " ")
	s = reShortcode.ReplaceAllString(s, " ")
	s = reFenceLine.ReplaceAllString(s, "\n")
	s = reLink.ReplaceAllString(s, " $1$2$3 ")
	s = reImage.ReplaceAllString(s, " $1 ")
	s = reAngle.ReplaceAllString(s, " ")
	s = reURL.ReplaceAllString(s, " ")
	s = reHeading.ReplaceAllString(s, "")
	s = reQuote.ReplaceAllString(s, "")
	s = reBullet.ReplaceAllString(s, " ")
	s = reTableSep.ReplaceAllString(s, "\n")
	return rePlain.Replace(s)
}

// splitFrontMatter separates Hugo front matter (--- or +++) from the body.
func splitFrontMatter(s string) (fm, body string) {
	t := strings.TrimLeft(strings.TrimLeft(s, "\ufeff"), " \t\r\n")
	var marker string
	switch {
	case strings.HasPrefix(t, "---"):
		marker = "---"
	case strings.HasPrefix(t, "+++"):
		marker = "+++"
	default:
		return "", s
	}
	rest := t[len(marker):]
	if i := strings.IndexByte(rest, '\n'); i >= 0 {
		rest = rest[i+1:]
	} else {
		return "", s
	}
	var sb strings.Builder
	for len(rest) > 0 {
		line := rest
		rest = ""
		if i := strings.IndexByte(line, '\n'); i >= 0 {
			line, rest = line[:i], line[i+1:]
		}
		if c := strings.TrimSpace(line); c == marker || c == "..." {
			return sb.String(), rest
		}
		sb.WriteString(line)
		sb.WriteByte('\n')
	}
	return "", s
}

func unquote(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
			s = strings.TrimSpace(s[1 : len(s)-1])
		}
	}
	return s
}

// parseMeta understands just enough YAML/TOML front matter for the handful of
// keys we care about: title, date, tags, categories, description, draft, slug.
func parseMeta(fm string) map[string][]string {
	meta := map[string][]string{}
	var key string
	for _, line := range strings.Split(fm, "\n") {
		t := strings.TrimSpace(line)
		if t == "" || strings.HasPrefix(t, "#") {
			continue
		}
		if strings.HasPrefix(t, "-") && key != "" {
			if v := unquote(strings.TrimLeft(t, "- ")); v != "" {
				meta[key] = append(meta[key], v)
			}
			continue
		}
		// front matter is YAML (key: value) or TOML (key = value)
		i, j := strings.IndexByte(t, ':'), strings.IndexByte(t, '=')
		var sep int
		switch {
		case i < 0 && j < 0:
			key = ""
			continue
		case j < 0 || (i >= 0 && i < j):
			sep = i
		default:
			sep = j
		}
		key = strings.ToLower(strings.TrimSpace(t[:sep]))
		val := strings.TrimSpace(t[sep+1:])
		switch {
		case val == "":
			meta[key] = nil
		case strings.HasPrefix(val, "[") && strings.HasSuffix(val, "]"):
			for _, p := range strings.Split(strings.Trim(val, "[]"), ",") {
				if p = unquote(p); p != "" {
					meta[key] = append(meta[key], p)
				}
			}
		default:
			meta[key] = append(meta[key], unquote(val))
		}
	}
	return meta
}

func first(meta map[string][]string, keys ...string) string {
	for _, k := range keys {
		if v := meta[k]; len(v) > 0 && v[0] != "" {
			return v[0]
		}
	}
	return ""
}

func normDate(s string) string {
	if s == "" {
		return ""
	}
	layouts := []string{
		time.RFC3339Nano, "2006-01-02T15:04:05", "2006-01-02 15:04:05",
		"2006-01-02", "2006/01/02", "02.01.2006", "January 2, 2006", "2 January 2006",
	}
	for _, l := range layouts {
		if t, err := time.Parse(l, s); err == nil {
			return t.Format("2006-01-02")
		}
	}
	return s
}

// ---------------------------------------------------------------------------
// tokenizer
// ---------------------------------------------------------------------------

// token is a word plus its byte range in the source text, so snippets can be
// cut out of the original text after the fact.
type token struct {
	text       string
	start, end int
}

// tokenize lowercases and splits on anything that is not a letter or digit.
// Short tokens are kept here (positions must stay consistent) and dropped when
// indexing.
func tokenize(s string) []token {
	var out []token
	var sb strings.Builder
	start := -1
	for i, r := range s {
		lr := unicode.ToLower(r)
		if unicode.IsLetter(lr) || unicode.IsDigit(lr) {
			if start < 0 {
				start = i
			}
			sb.WriteRune(lr)
			continue
		}
		if start >= 0 {
			out = append(out, token{sb.String(), start, i})
			sb.Reset()
			start = -1
		}
	}
	if start >= 0 {
		out = append(out, token{sb.String(), start, len(s)})
	}
	return out
}

func words(toks []token) []string {
	out := make([]string, 0, len(toks))
	for _, t := range toks {
		if len([]rune(t.text)) >= minIndexToken {
			out = append(out, t.text)
		}
	}
	return out
}

// ---------------------------------------------------------------------------
// index
// ---------------------------------------------------------------------------

// Doc is one indexed file. Text holds the cleaned body so that the browser can
// render snippets without fetching the page.
type Doc struct {
	Path    string   `json:"path"`
	URL     string   `json:"url"`
	Title   string   `json:"title"`
	Date    string   `json:"date,omitempty"`
	Tags    []string `json:"tags,omitempty"`
	Summary string   `json:"summary,omitempty"`
	Text    string   `json:"text,omitempty"`
}

// Field is one searchable field (title, tag, body).
//
// Post maps a term to a flat posting list laid out as
//
//	[doc, n, pos0, ... posN-1, doc, n, pos0, ...]
//
// which keeps the JSON small while still supporting phrase queries.
type Field struct {
	Name  string             `json:"name"`
	Boost float64            `json:"boost"`
	Len   []int32            `json:"len"`
	Total int64              `json:"total"`
	Post  map[string][]int32 `json:"post"`
}

type Index struct {
	Version   int      `json:"version"`
	Generated string   `json:"generated,omitempty"`
	Base      string   `json:"base,omitempty"`
	Docs      []Doc    `json:"docs"`
	Fields    []*Field `json:"fields"`
}

type BuildOpts struct {
	Base          string // URL prefix prepended to result URLs
	MaxText       int    // cap indexed/stored body runes per doc (0 = no cap)
	IncludeDrafts bool
}

func NewIndex(opts BuildOpts) *Index {
	return &Index{
		Version: indexVersion,
		Base:    strings.TrimRight(opts.Base, "/"),
		Fields: []*Field{
			{Name: "title", Boost: 5, Post: map[string][]int32{}},
			{Name: "tag", Boost: 3, Post: map[string][]int32{}},
			{Name: "body", Boost: 1, Post: map[string][]int32{}},
		},
	}
}

func (idx *Index) field(name string) *Field {
	for _, f := range idx.Fields {
		if f.Name == name {
			return f
		}
	}
	return nil
}

// addDoc indexes one document. Positions refer to the full token stream (short
// tokens are skipped but still counted) so snippets can be cut out of the text.
func (f *Field) addDoc(docID int32, toks []token) {
	positions := map[string][]int32{}
	n := 0
	for i, t := range toks {
		if len([]rune(t.text)) < minIndexToken {
			continue
		}
		positions[t.text] = append(positions[t.text], int32(i))
		n++
	}
	f.Len = append(f.Len, int32(n))
	f.Total += int64(n)
	for term, pos := range positions {
		p := f.Post[term]
		p = append(p, docID, int32(len(pos)))
		f.Post[term] = append(p, pos...)
	}
}

// eachPosting walks a flat posting list and returns the number of documents.
func eachPosting(p []int32, fn func(doc int32, pos []int32)) int {
	n, i := 0, 0
	for i+2 <= len(p) {
		doc, cnt := p[i], int(p[i+1])
		if cnt < 0 || i+2+cnt > len(p) {
			break
		}
		fn(doc, p[i+2:i+2+cnt])
		n++
		i += 2 + cnt
	}
	return n
}

// pageURL maps a source path to the URL Hugo will publish it under.
func pageURL(name string) string {
	p := filepath.ToSlash(name)
	p = strings.TrimPrefix(p, "./")
	if i := strings.Index(p, "content/"); i == 0 || (i > 0 && p[i-1] == '/') {
		p = p[i+len("content/"):]
	}
	dir, base := path.Split(p)
	if base == "_index.md" || base == "index.md" || base == "_index.markdown" {
		base = ""
	} else {
		base = strings.TrimSuffix(base, path.Ext(base))
	}
	u := path.Join("/", dir, base)
	if !strings.HasSuffix(u, "/") {
		u += "/"
	}
	return u
}

var reFirstHeading = regexp.MustCompile(`(?m)^ {0,3}#{1,6}[ \t]*(.+?)[ \t]*#*$`)

// AddDoc parses one markdown file and adds it to the index. It reports whether
// the file was indexed (drafts are skipped unless asked for).
func AddDoc(idx *Index, name string, content []byte, modTime time.Time, opts BuildOpts) bool {
	fmStr, body := splitFrontMatter(string(content))
	meta := parseMeta(fmStr)
	if !opts.IncludeDrafts {
		if strings.EqualFold(first(meta, "draft"), "true") {
			return false
		}
	}

	text := markdownToText(body)
	if opts.MaxText > 0 {
		if r := []rune(text); len(r) > opts.MaxText {
			text = string(r[:opts.MaxText])
		}
	}

	title := first(meta, "title")
	if title == "" {
		if m := reFirstHeading.FindStringSubmatch(body); len(m) > 1 {
			title = strings.TrimSpace(m[1])
		}
	}
	if title == "" {
		base := strings.TrimSuffix(filepath.Base(name), filepath.Ext(name))
		title = strings.NewReplacer("-", " ", "_", " ").Replace(base)
	}

	date := normDate(first(meta, "date", "lastmod", "publishdate"))
	if date == "" && !modTime.IsZero() {
		date = modTime.Format("2006-01-02")
	}

	var tags []string
	seen := map[string]bool{}
	for _, t := range append(append([]string{}, meta["tags"]...), meta["categories"]...) {
		if t != "" && !seen[strings.ToLower(t)] {
			seen[strings.ToLower(t)] = true
			tags = append(tags, t)
		}
	}

	url := first(meta, "url")
	if url == "" {
		base := filepath.Base(name)
		if s := first(meta, "slug"); s != "" {
			base = s + ".md"
		}
		dir, _ := path.Split(filepath.ToSlash(name))
		url = pageURL(path.Join(dir, base))
	}
	if !strings.HasPrefix(url, "/") {
		url = "/" + url
	}
	if idx.Base != "" {
		url = idx.Base + url
	}

	doc := Doc{
		Path:    filepath.ToSlash(name),
		URL:     url,
		Title:   title,
		Date:    date,
		Tags:    tags,
		Summary: first(meta, "description", "summary", "excerpt"),
		Text:    text,
	}
	id := int32(len(idx.Docs))
	idx.Docs = append(idx.Docs, doc)

	idx.field("title").addDoc(id, tokenize(title))
	idx.field("tag").addDoc(id, tokenize(strings.Join(tags, " ")))
	idx.field("body").addDoc(id, tokenize(text))
	return true
}

// ---------------------------------------------------------------------------
// index I/O
// ---------------------------------------------------------------------------

func WriteIndex(w io.Writer, idx *Index) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", " ")
	return enc.Encode(idx)
}

func ReadIndex(r io.Reader) (*Index, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if len(data) > 2 && data[0] == 0x1f && data[1] == 0x8b {
		zr, err := gzip.NewReader(bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
		defer zr.Close()
		if data, err = io.ReadAll(zr); err != nil {
			return nil, err
		}
	}
	idx := &Index{}
	if err := json.Unmarshal(data, idx); err != nil {
		return nil, err
	}
	if idx.Version != indexVersion {
		return nil, fmt.Errorf("index version %d, this build understands %d", idx.Version, indexVersion)
	}
	for _, f := range idx.Fields {
		if f.Post == nil {
			f.Post = map[string][]int32{}
		}
	}
	return idx, nil
}

// ---------------------------------------------------------------------------
// search
// ---------------------------------------------------------------------------

// SearchOpts tunes ranking behaviour.
type SearchOpts struct {
	Limit      int
	Prefix     bool // expand terms without an exact match as prefixes
	MaxExpand  int  // cap for prefix expansion
	RequireAll bool // default true: every term must match, fall back to OR
}

// Hit is one ranked result.
type Hit struct {
	Rank    int      `json:"rank"`
	Doc     int      `json:"doc"`
	Score   float64  `json:"score"`
	Title   string   `json:"title"`
	URL     string   `json:"url"`
	Date    string   `json:"date,omitempty"`
	Tags    []string `json:"tags,omitempty"`
	Snippet string   `json:"snippet"`
	Terms   []string `json:"terms,omitempty"`
}

type clause struct {
	phrase bool // "quoted" or derived from a punctuated word
	loose  bool // derived from one query word; may degrade to AND of its words
	text   string
	prefix bool
}

var stopwords = map[string]bool{
	"a": true, "an": true, "and": true, "are": true, "as": true, "at": true,
	"be": true, "but": true, "by": true, "do": true, "does": true, "for": true,
	"from": true, "had": true, "has": true, "have": true, "how": true, "i": true,
	"in": true, "is": true, "it": true, "not": true, "of": true, "on": true,
	"or": true, "our": true, "so": true, "that": true, "the": true, "their": true,
	"there": true, "these": true, "they": true, "this": true, "to": true,
	"too": true, "us": true, "was": true, "were": true, "what": true,
	"when": true, "where": true, "which": true, "who": true, "why": true,
	"will": true, "with": true, "you": true, "your": true,
}

// parseQuery splits a query into term and phrase clauses.
func parseQuery(q string) []clause {
	var out []clause
	rs := []rune(q)
	i := 0
	for i < len(rs) {
		c := rs[i]
		switch {
		case c == '"' || c == '\u201c' || c == '\u201d':
			j := i + 1
			for j < len(rs) && rs[j] != '"' && rs[j] != '\u201c' && rs[j] != '\u201d' {
				j++
			}
			if ph := strings.TrimSpace(string(rs[i+1 : min(j, len(rs))])); ph != "" {
				out = append(out, clause{phrase: true, text: ph})
			}
			i = j + 1
		case unicode.IsSpace(c):
			i++
		default:
			j := i
			for j < len(rs) && !unicode.IsSpace(rs[j]) && rs[j] != '"' {
				j++
			}
			w := strings.ToLower(strings.TrimSpace(string(rs[i:j])))
			cl := clause{}
			if strings.HasSuffix(w, "*") {
				cl.prefix = true
				w = strings.TrimRight(w, "*")
			}
			if w != "" {
				// A word like meetup-59 or cloud-native is indexed as two
				// tokens; search for them next to each other, falling back to
				// "both words anywhere".
				if ws := words(tokenize(w)); len(ws) > 1 {
					out = append(out, clause{phrase: true, loose: true, text: strings.Join(ws, " ")})
				} else {
					cl.text = w
					out = append(out, cl)
				}
			}
			i = j
		}
	}
	return out
}

// avgDocLen returns the mean field length used by BM25's length normalisation.
func (f *Field) avgDocLen() float64 {
	if len(f.Len) == 0 {
		return 1
	}
	return math.Max(1, float64(f.Total)/float64(len(f.Len)))
}

func bm25(tf, dl int, avgdl float64, df, n int) float64 {
	idf := math.Log(1 + (float64(n-df)+0.5)/(float64(df)+0.5))
	return idf * (float64(tf) * (bm25K1 + 1)) / (float64(tf) + bm25K1*(1-bm25B+bm25B*float64(dl)/avgdl))
}

type matchInfo struct {
	score     float64
	positions []int32 // body positions, used to cut the snippet
}

// expand returns the index terms a prefix matches, most frequent first.
func (f *Field) expand(prefix string, max int) []string {
	var cands []string
	for t := range f.Post {
		if t != prefix && strings.HasPrefix(t, prefix) {
			cands = append(cands, t)
		}
	}
	if len(cands) <= 1 {
		return cands
	}
	sort.Slice(cands, func(i, j int) bool {
		li, lj := len(f.Post[cands[i]]), len(f.Post[cands[j]])
		if li != lj {
			return li > lj
		}
		return cands[i] < cands[j]
	})
	if max > 0 && len(cands) > max {
		cands = cands[:max]
	}
	return cands
}

// matchTerm scores every document containing the term (or a prefix expansion).
func (idx *Index) matchTerm(term string, prefix bool, opts SearchOpts) map[int32]*matchInfo {
	res := map[int32]*matchInfo{}
	n := len(idx.Docs)
	for _, f := range idx.Fields {
		terms := []string{term}
		if _, ok := f.Post[term]; !ok && prefix && opts.MaxExpand > 0 {
			terms = f.expand(term, opts.MaxExpand)
		}
		avgdl := f.avgDocLen()
		for _, t := range terms {
			post := f.Post[t]
			df := 0
			eachPosting(post, func(int32, []int32) { df++ })
			if df == 0 {
				continue
			}
			eachPosting(post, func(doc int32, pos []int32) {
				dl := 1
				if int(doc) < len(f.Len) && f.Len[doc] > 0 {
					dl = int(f.Len[doc])
				}
				m := res[doc]
				if m == nil {
					m = &matchInfo{}
					res[doc] = m
				}
				m.score += f.Boost * bm25(len(pos), dl, avgdl, df, n)
				if f.Name == "body" {
					m.positions = append(m.positions, pos...)
				}
			})
		}
	}
	return res
}

// matchPhrase finds documents where the words appear adjacent, in order.
func (idx *Index) matchPhrase(phrase string) map[int32]*matchInfo {
	res := map[int32]*matchInfo{}
	phraseWords := words(tokenize(phrase))
	if len(phraseWords) == 0 {
		return res
	}
	if len(phraseWords) == 1 {
		return idx.matchTerm(phraseWords[0], false, SearchOpts{Prefix: true, MaxExpand: defaultMaxExpand})
	}
	n := len(idx.Docs)
	for _, f := range idx.Fields {
		avgdl := f.avgDocLen()
		per := make([]map[int32][]int32, len(phraseWords))
		df := make([]int, len(phraseWords))
		ok := true
		for i, w := range phraseWords {
			per[i] = map[int32][]int32{}
			eachPosting(f.Post[w], func(doc int32, pos []int32) {
				per[i][doc] = pos
			})
			df[i] = len(per[i])
			if df[i] == 0 {
				ok = false
			}
		}
		if !ok {
			continue
		}
		for doc, firstPos := range per[0] {
			for _, p := range firstPos {
				cur, adjacent := p, true
				for i := 1; i < len(per); i++ {
					found := false
					for _, q := range per[i][doc] {
						if q == cur+1 {
							cur, found = q, true
							break
						}
					}
					if !found {
						adjacent = false
						break
					}
				}
				if !adjacent {
					continue
				}
				dl := 1
				if int(doc) < len(f.Len) && f.Len[doc] > 0 {
					dl = int(f.Len[doc])
				}
				score := 0.0
				for i := range phraseWords {
					score += bm25(len(per[i][doc]), dl, avgdl, df[i], n)
				}
				m := res[doc]
				if m == nil {
					m = &matchInfo{}
					res[doc] = m
				}
				m.score += f.Boost * score * 1.5 // phrase bonus
				if f.Name == "body" {
					m.positions = append(m.positions, firstPos...)
				}
				break
			}
		}
	}
	return res
}

// combine merges the per-clause result maps. With and=true only documents
// matched by every clause survive.
func combine(lists []map[int32]*matchInfo, and bool) map[int32]*matchInfo {
	out := map[int32]*matchInfo{}
	for _, l := range lists {
		for doc, m := range l {
			e, ok := out[doc]
			if !ok {
				out[doc] = &matchInfo{score: m.score, positions: append([]int32(nil), m.positions...)}
				continue
			}
			e.score += m.score
			e.positions = append(e.positions, m.positions...)
		}
	}
	if !and {
		return out
	}
	for doc := range out {
		for _, l := range lists {
			if _, ok := l[doc]; !ok {
				delete(out, doc)
				break
			}
		}
	}
	return out
}

// Search runs a query against the index and returns ranked hits.
func Search(idx *Index, q string, opts SearchOpts) []Hit {
	if len(idx.Docs) == 0 {
		return nil
	}
	if opts.MaxExpand == 0 {
		opts.MaxExpand = defaultMaxExpand
	}
	clauses := parseQuery(q)
	keep := []clause{}
	for _, cl := range clauses {
		if !cl.phrase && !cl.prefix && stopwords[cl.text] {
			continue
		}
		keep = append(keep, cl)
	}
	if len(keep) > 0 {
		clauses = keep
	}
	if len(clauses) == 0 {
		return nil
	}

	lists := make([]map[int32]*matchInfo, 0, len(clauses))
	terms := map[string]bool{}
	for _, cl := range clauses {
		var m map[int32]*matchInfo
		if cl.phrase {
			m = idx.matchPhrase(cl.text)
			if len(m) == 0 && cl.loose {
				// words of e.g. "meetup-59" do not appear adjacent; require
				// them anywhere in the document instead
				var subs []map[int32]*matchInfo
				for _, w := range words(tokenize(cl.text)) {
					subs = append(subs, idx.matchTerm(w, false, opts))
				}
				if m = combine(subs, true); len(m) == 0 {
					m = combine(subs, false)
				}
			}
		} else {
			m = idx.matchTerm(cl.text, cl.prefix || opts.Prefix, opts)
		}
		lists = append(lists, m)
		for _, w := range words(tokenize(cl.text)) {
			terms[w] = true
		}
	}

	matches := combine(lists, opts.RequireAll)
	if len(matches) == 0 && opts.RequireAll {
		matches = combine(lists, false) // degrade to OR
	}
	if len(matches) == 0 {
		return nil
	}

	docs := make([]int32, 0, len(matches))
	for doc := range matches {
		docs = append(docs, doc)
	}
	sort.Slice(docs, func(i, j int) bool {
		a, b := matches[docs[i]], matches[docs[j]]
		if a.score != b.score {
			return a.score > b.score
		}
		da, db := idx.Docs[docs[i]].Date, idx.Docs[docs[j]].Date
		if da != db {
			return da > db
		}
		return idx.Docs[docs[i]].Title < idx.Docs[docs[j]].Title
	})
	if opts.Limit > 0 && len(docs) > opts.Limit {
		docs = docs[:opts.Limit]
	}

	hits := make([]Hit, 0, len(docs))
	for i, doc := range docs {
		d := idx.Docs[doc]
		// centre the snippet on the rarest clause that matched this document,
		// that is the one most worth showing
		positions := matches[doc].positions
		best, bestCount := -1, 0
		for ci, l := range lists {
			if _, ok := l[doc]; ok && (best < 0 || len(l) < bestCount) {
				best, bestCount = ci, len(l)
			}
		}
		if best >= 0 {
			positions = lists[best][doc].positions
		}
		hits = append(hits, Hit{
			Rank:    i + 1,
			Doc:     int(doc),
			Score:   math.Round(matches[doc].score*1000) / 1000,
			Title:   d.Title,
			URL:     d.URL,
			Date:    d.Date,
			Tags:    d.Tags,
			Snippet: snippet(&d, positions),
			Terms:   sortedKeys(terms),
		})
	}
	return hits
}

// queryTerms returns the words and the prefix stems of a query; used to
// highlight matches in the results.
func queryTerms(q string) (exact, prefixes map[string]bool) {
	exact, prefixes = map[string]bool{}, map[string]bool{}
	for _, cl := range parseQuery(q) {
		if cl.prefix {
			prefixes[cl.text] = true
			continue
		}
		for _, w := range words(tokenize(cl.text)) {
			exact[w] = true
		}
	}
	return exact, prefixes
}

func sortedKeys(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

var reWS = regexp.MustCompile(`\s+`)

func oneLine(s string) string { return strings.TrimSpace(reWS.ReplaceAllString(s, " ")) }

// snippet cuts a window of text around the first matched body position.
func snippet(d *Doc, positions []int32) string {
	const before, after = 16, 44
	toks := tokenize(d.Text)
	if len(toks) == 0 {
		if s := oneLine(d.Summary); s != "" {
			return s
		}
		return oneLine(d.Title)
	}
	center, found := 0, false
	for _, p := range positions {
		if p >= 0 && int(p) < len(toks) && (!found || p < int32(center)) {
			center, found = int(p), true
		}
	}
	lo, hi := max(0, center-before), min(len(toks), center+after)
	var sb strings.Builder
	if lo > 0 {
		sb.WriteString("… ")
	}
	sb.WriteString(oneLine(d.Text[toks[lo].start:toks[hi-1].end]))
	if hi < len(toks) {
		sb.WriteString(" …")
	}
	return sb.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ---------------------------------------------------------------------------
// CLI
// ---------------------------------------------------------------------------

const usageText = `quickindex - tiny full text search for markdown files

build an index (file names on stdin or as arguments):
  fd -e md | quickindex -o my.idx
  fd -e md | quickindex -base https://golangleipzig.space -o search.idx
  quickindex -o my.idx content/posts/*.md

query the index:
  quickindex -f my.idx "agentic coding"
  quickindex -f my.idx go* wasm
  quickindex -f my.idx -json '"code of conduct"'

Query syntax: words (all must match, case insensitive), "phrases", trailing-*
prefix terms. Search is case insensitive and diacritic sensitive.

  -0                 file names on stdin are NUL separated (fd --print0)
  -all               print every match, ignoring -n
  -base url          URL prefix stored with every result
  -f file            index to search (query mode)
  -include-drafts    also index files whose front matter says draft: true
  -json              compact index output / results as JSON
  -max-text n        index at most n body runes per file (0 = no limit)
  -n int             maximum number of results (default 10)
  -no-color          never colourise
  -no-prefix         do not prefix-expand terms that have no exact match
  -o file            write the index here (default: stdout)
  -q                 do not print build statistics
`

func main() {
	var (
		outPath  = flag.String("o", "", "write the index here (build mode); default stdout")
		idxPath  = flag.String("f", "", "index file to search (query mode)")
		limit    = flag.Int("n", 10, "maximum number of results")
		all      = flag.Bool("all", false, "print every match, ignoring -n")
		asJSON   = flag.Bool("json", false, "compact index / JSON results")
		base     = flag.String("base", "", "URL prefix for results, e.g. https://golangleipzig.space")
		maxText  = flag.Int("max-text", 0, "index at most this many body runes per file (0 = no limit)")
		drafts   = flag.Bool("include-drafts", false, "also index files with draft: true")
		noPrefix = flag.Bool("no-prefix", false, "do not prefix-expand unknown terms")
		nul      = flag.Bool("0", false, "NUL separated file names on stdin (fd --print0)")
		quiet    = flag.Bool("q", false, "do not print build statistics")
		noColor  = flag.Bool("no-color", false, "never colourise output")
	)
	flag.Usage = func() { fmt.Fprint(os.Stderr, usageText) }
	flag.Parse()

	var err error
	if *idxPath != "" {
		err = runQuery(*idxPath, strings.Join(flag.Args(), " "), *limit, *all, *asJSON, !*noPrefix, !*noColor && stdoutIsTerminal())
	} else {
		err = runBuild(flag.Args(), *outPath, *nul, BuildOpts{Base: *base, MaxText: *maxText, IncludeDrafts: *drafts}, *quiet, *asJSON)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "quickindex:", err)
		os.Exit(1)
	}
}

func runBuild(args []string, outPath string, nul bool, opts BuildOpts, quiet, compact bool) error {
	names, err := collectFiles(args, nul)
	if err != nil {
		return err
	}
	if len(names) == 0 {
		flag.Usage()
		return fmt.Errorf("no input files (try: fd -e md | quickindex -o my.idx)")
	}
	idx := NewIndex(opts)
	idx.Generated = time.Now().UTC().Format(time.RFC3339)
	skipped, failed := 0, 0
	for _, name := range names {
		data, err := os.ReadFile(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "quickindex: %v\n", err)
			failed++
			continue
		}
		var mod time.Time
		if st, err := os.Stat(name); err == nil {
			mod = st.ModTime()
		}
		if !AddDoc(idx, name, data, mod, opts) {
			skipped++
		}
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if !compact {
		enc.SetIndent("", " ")
	}
	if err := enc.Encode(idx); err != nil {
		return err
	}

	out := io.Writer(os.Stdout)
	dest := "stdout"
	if outPath != "" {
		f, err := os.Create(outPath)
		if err != nil {
			return err
		}
		defer f.Close()
		out, dest = f, outPath
	}
	data := buf.Bytes()
	if strings.HasSuffix(dest, ".gz") {
		var gz bytes.Buffer
		gw := gzip.NewWriter(&gz)
		if _, err := gw.Write(data); err != nil {
			return err
		}
		if err := gw.Close(); err != nil {
			return err
		}
		data = gz.Bytes()
	}
	if _, err := out.Write(data); err != nil {
		return err
	}
	if !quiet {
		terms := 0
		for _, f := range idx.Fields {
			terms += len(f.Post)
		}
		fmt.Fprintf(os.Stderr, "indexed %d documents (%d skipped, %d unreadable), %d distinct terms, %s -> %s\n",
			len(idx.Docs), skipped, failed, terms, humanSize(len(data)), dest)
	}
	return nil
}

type queryResult struct {
	Query string `json:"query"`
	Total int    `json:"total"`
	Hits  []Hit  `json:"hits"`
}

func runQuery(idxPath, q string, limit int, all, asJSON, prefix, color bool) error {
	if strings.TrimSpace(q) == "" {
		return fmt.Errorf("missing query (quickindex -f %s QUERY)", idxPath)
	}
	data, err := os.ReadFile(idxPath)
	if err != nil {
		return err
	}
	idx, err := ReadIndex(bytes.NewReader(data))
	if err != nil {
		return err
	}
	if all {
		limit = 0
	}
	hits := Search(idx, q, SearchOpts{Limit: limit, Prefix: prefix, RequireAll: true})

	if asJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(queryResult{Query: q, Total: len(hits), Hits: hits})
	}

	terms, prefixes := queryTerms(q)
	if len(hits) == 0 {
		fmt.Printf("no match for %q\n", q)
		return nil
	}
	fmt.Printf("%d match(es) for %q\n\n", len(hits), q)
	for _, h := range hits {
		fmt.Printf("%s %s %s\n", paint(fmt.Sprintf("%2d.", h.Rank), "2", color),
			paint(h.Title, "1", color), paint(fmt.Sprintf("(%.2f)", h.Score), "2", color))
		meta := h.URL
		if h.Date != "" {
			meta += "  ·  " + h.Date
		}
		if len(h.Tags) > 0 {
			meta += "  ·  " + strings.Join(h.Tags, ", ")
		}
		fmt.Printf("    %s\n", paint(meta, "36", color))
		if h.Snippet != "" {
			fmt.Printf("    %s\n\n", highlight(h.Snippet, terms, prefixes, paint("$1", "1;36", color)))
		}
	}
	return nil
}

// highlight wraps every token that is a query term (or matches a prefix term),
// using paint's output, which must contain $1 where the match goes.
func highlight(text string, terms, prefixes map[string]bool, repl string) string {
	if len(terms) == 0 && len(prefixes) == 0 || !strings.Contains(repl, "$1") {
		return text
	}
	re := regexp.MustCompile(`\$1`)
	var sb strings.Builder
	last := 0
	for _, t := range tokenize(text) {
		hit := terms[t.text]
		if !hit {
			for p := range prefixes {
				if strings.HasPrefix(t.text, p) {
					hit = true
					break
				}
			}
		}
		if !hit {
			continue
		}
		sb.WriteString(text[last:t.start])
		sb.WriteString(re.ReplaceAllString(repl, text[t.start:t.end]))
		last = t.end
	}
	sb.WriteString(text[last:])
	return sb.String()
}

func paint(s, code string, on bool) string {
	if !on {
		return s
	}
	return "\x1b[" + code + "m" + s + "\x1b[0m"
}

func stdoutIsTerminal() bool {
	st, err := os.Stdout.Stat()
	return err == nil && st.Mode()&os.ModeCharDevice != 0
}

func stdinPiped() bool {
	st, err := os.Stdin.Stat()
	return err == nil && st.Mode()&os.ModeCharDevice == 0
}

// collectFiles merges positional arguments with the file names on stdin.
func collectFiles(args []string, nul bool) ([]string, error) {
	seen := map[string]bool{}
	var out []string
	add := func(s string) {
		s = strings.TrimSpace(strings.Trim(strings.TrimSpace(s), `"`))
		if s == "" || s == "-" || seen[s] {
			return
		}
		seen[s] = true
		out = append(out, s)
	}
	for _, a := range args {
		add(a)
	}
	if stdinPiped() {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
		sep := "\n"
		if nul {
			sep = "\x00"
		}
		for _, line := range strings.Split(string(data), sep) {
			add(line)
		}
	}
	return out, nil
}

func humanSize(n int) string {
	switch {
	case n < 1024:
		return fmt.Sprintf("%dB", n)
	case n < 1024*1024:
		return fmt.Sprintf("%.1fKB", float64(n)/1024)
	default:
		return fmt.Sprintf("%.1fMB", float64(n)/1024/1024)
	}
}
