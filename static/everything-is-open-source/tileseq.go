// tileseq generates an animated GIF on a grid of colored tiles.
//
// The canvas is subdivided into equally sized square tiles. The animation runs
// in three phases:
//
//	1. fluctuate - every tile flickers through random palette colors
//	2. reveal    - the requested text is progressively drawn onto the grid
//	3. freeze    - the finished text is held still long enough to read
//
// A built-in 5x7 ASCII font (uppercase letters, digits and a little
// punctuation) is mapped onto the grid, optionally scaled up so each font
// pixel covers several tiles.
//
// Usage:
//
//	go run tileseq.go -text "GO" -scale 2 -bold 1 -textcolor "#ffffff" -output out.gif
//
// Don't worry much about GIF size here; control it via -width/-height/-tile or
// post-process with tools like gifsicle.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
)

// ----------------------------------------------------------------------------
// Palettes
// ----------------------------------------------------------------------------

// palettes maps a name to a set of "background" colors that tiles flicker
// through. The text and background-base colors are handled separately.
var palettes = map[string][]color.RGBA{
	"warm": {
		{255, 99, 71, 255},   // tomato
		{255, 140, 0, 255},   // dark orange
		{255, 215, 0, 255},   // gold
		{220, 20, 60, 255},   // crimson
		{255, 165, 0, 255},   // orange
		{210, 105, 30, 255},  // chocolate
		{255, 127, 80, 255},  // coral
	},
	"cool": {
		{0, 191, 255, 255},   // deep sky blue
		{30, 144, 255, 255},  // dodger blue
		{0, 206, 209, 255},   // dark turquoise
		{72, 209, 204, 255},  // medium turquoise
		{65, 105, 225, 255},  // royal blue
		{106, 90, 205, 255},  // slate blue
		{0, 139, 139, 255},   // dark cyan
	},
	"neon": {
		{255, 0, 255, 255},   // magenta
		{0, 255, 255, 255},   // cyan
		{57, 255, 20, 255},   // neon green
		{255, 49, 49, 255},   // neon red
		{255, 255, 0, 255},   // yellow
		{188, 19, 254, 255},  // electric purple
		{255, 105, 180, 255}, // hot pink
	},
	"pastel": {
		{255, 209, 220, 255}, // pastel pink
		{174, 198, 207, 255}, // pastel blue
		{193, 225, 193, 255}, // pastel green
		{253, 253, 150, 255}, // pastel yellow
		{222, 165, 233, 255}, // pastel purple
		{255, 179, 153, 255}, // pastel orange
		{176, 224, 230, 255}, // powder blue
	},
	"mono": {
		{40, 40, 40, 255},
		{80, 80, 80, 255},
		{120, 120, 120, 255},
		{160, 160, 160, 255},
		{200, 200, 200, 255},
		{230, 230, 230, 255},
	},
	"rainbow": {
		{255, 0, 0, 255},
		{255, 127, 0, 255},
		{255, 255, 0, 255},
		{0, 255, 0, 255},
		{0, 0, 255, 255},
		{75, 0, 130, 255},
		{148, 0, 211, 255},
	},
	"forest": {
		{34, 139, 34, 255},   // forest green
		{85, 107, 47, 255},   // dark olive
		{107, 142, 35, 255},  // olive drab
		{154, 205, 50, 255},  // yellow green
		{46, 139, 87, 255},   // sea green
		{0, 100, 0, 255},     // dark green
		{189, 183, 107, 255}, // dark khaki
	},
	"ocean": {
		{0, 105, 148, 255},   // deep ocean
		{0, 150, 199, 255},   // cerulean
		{72, 202, 228, 255},  // sky
		{144, 224, 239, 255}, // foam
		{2, 62, 138, 255},    // navy
		{0, 180, 216, 255},   // bright cyan-blue
		{173, 232, 244, 255}, // pale aqua
	},
	"sunset": {
		{255, 94, 77, 255},   // coral red
		{255, 149, 5, 255},   // amber
		{255, 195, 0, 255},   // golden
		{199, 0, 57, 255},    // deep rose
		{144, 12, 63, 255},   // wine
		{255, 127, 80, 255},  // coral
		{88, 24, 69, 255},    // dusk purple
	},
	"candy": {
		{255, 105, 180, 255}, // hot pink
		{255, 182, 193, 255}, // light pink
		{152, 245, 255, 255}, // baby cyan
		{186, 156, 255, 255}, // lavender
		{255, 250, 150, 255}, // lemon
		{170, 255, 195, 255}, // mint
		{255, 160, 200, 255}, // bubblegum
	},
	"earth": {
		{139, 90, 43, 255},   // saddle brown
		{160, 120, 80, 255},  // clay
		{205, 170, 125, 255}, // tan
		{120, 100, 60, 255},  // umber
		{222, 184, 135, 255}, // burlywood
		{101, 67, 33, 255},   // dark brown
		{188, 158, 130, 255}, // sand
	},
	"miami": {
		{0, 255, 209, 255},   // turquoise
		{255, 0, 144, 255},   // magenta-pink
		{255, 209, 0, 255},   // yellow
		{120, 0, 255, 255},   // violet
		{0, 209, 255, 255},   // cyan
		{255, 110, 199, 255}, // pink
		{72, 255, 0, 255},    // lime
	},
}

func paletteNames() string {
	names := make([]string, 0, len(palettes))
	for n := range palettes {
		names = append(names, n)
	}
	sort.Strings(names)
	return strings.Join(names, ", ")
}

// ----------------------------------------------------------------------------
// Font: 5x7 bitmaps for a usable subset of ASCII.
// ----------------------------------------------------------------------------

const (
	glyphW = 5
	glyphH = 7
)

// font holds glyphs as 7 rows of 5 characters; '#' (or any non-space) is "on".
var font = map[rune][]string{
	'A': {"01110", "10001", "10001", "11111", "10001", "10001", "10001"},
	'B': {"11110", "10001", "10001", "11110", "10001", "10001", "11110"},
	'C': {"01110", "10001", "10000", "10000", "10000", "10001", "01110"},
	'D': {"11110", "10001", "10001", "10001", "10001", "10001", "11110"},
	'E': {"11111", "10000", "10000", "11110", "10000", "10000", "11111"},
	'F': {"11111", "10000", "10000", "11110", "10000", "10000", "10000"},
	'G': {"01110", "10001", "10000", "10111", "10001", "10001", "01111"},
	'H': {"10001", "10001", "10001", "11111", "10001", "10001", "10001"},
	'I': {"11111", "00100", "00100", "00100", "00100", "00100", "11111"},
	'J': {"00111", "00010", "00010", "00010", "00010", "10010", "01100"},
	'K': {"10001", "10010", "10100", "11000", "10100", "10010", "10001"},
	'L': {"10000", "10000", "10000", "10000", "10000", "10000", "11111"},
	'M': {"10001", "11011", "10101", "10101", "10001", "10001", "10001"},
	'N': {"10001", "10001", "11001", "10101", "10011", "10001", "10001"},
	'O': {"01110", "10001", "10001", "10001", "10001", "10001", "01110"},
	'P': {"11110", "10001", "10001", "11110", "10000", "10000", "10000"},
	'Q': {"01110", "10001", "10001", "10001", "10101", "10010", "01101"},
	'R': {"11110", "10001", "10001", "11110", "10100", "10010", "10001"},
	'S': {"01111", "10000", "10000", "01110", "00001", "00001", "11110"},
	'T': {"11111", "00100", "00100", "00100", "00100", "00100", "00100"},
	'U': {"10001", "10001", "10001", "10001", "10001", "10001", "01110"},
	'V': {"10001", "10001", "10001", "10001", "10001", "01010", "00100"},
	'W': {"10001", "10001", "10001", "10101", "10101", "11011", "10001"},
	'X': {"10001", "10001", "01010", "00100", "01010", "10001", "10001"},
	'Y': {"10001", "10001", "01010", "00100", "00100", "00100", "00100"},
	'Z': {"11111", "00001", "00010", "00100", "01000", "10000", "11111"},
	'0': {"01110", "10001", "10011", "10101", "11001", "10001", "01110"},
	'1': {"00100", "01100", "00100", "00100", "00100", "00100", "01110"},
	'2': {"01110", "10001", "00001", "00110", "01000", "10000", "11111"},
	'3': {"11111", "00010", "00100", "00010", "00001", "10001", "01110"},
	'4': {"00010", "00110", "01010", "10010", "11111", "00010", "00010"},
	'5': {"11111", "10000", "11110", "00001", "00001", "10001", "01110"},
	'6': {"00110", "01000", "10000", "11110", "10001", "10001", "01110"},
	'7': {"11111", "00001", "00010", "00100", "01000", "01000", "01000"},
	'8': {"01110", "10001", "10001", "01110", "10001", "10001", "01110"},
	'9': {"01110", "10001", "10001", "01111", "00001", "00010", "01100"},
	' ': {"00000", "00000", "00000", "00000", "00000", "00000", "00000"},
	'.': {"00000", "00000", "00000", "00000", "00000", "00110", "00110"},
	',': {"00000", "00000", "00000", "00000", "00110", "00100", "01000"},
	'!': {"00100", "00100", "00100", "00100", "00100", "00000", "00100"},
	'?': {"01110", "10001", "00001", "00110", "00100", "00000", "00100"},
	'-': {"00000", "00000", "00000", "11111", "00000", "00000", "00000"},
	':': {"00000", "00110", "00110", "00000", "00110", "00110", "00000"},
	'\'': {"00100", "00100", "01000", "00000", "00000", "00000", "00000"},
}

// blockSize returns the width of the widest line and the total height of all
// lines stacked vertically, both in tiles.
func blockSize(lines []string, scale, gap, lineGap int) (w, h int) {
	advance := glyphW*scale + gap*scale
	for _, ln := range lines {
		n := len([]rune(ln))
		lw := 0
		if n > 0 {
			lw = n*advance - gap*scale
		}
		if lw > w {
			w = lw
		}
	}
	h = len(lines)*glyphH*scale + (len(lines)-1)*lineGap*scale
	if h < 0 {
		h = 0
	}
	return w, h
}

// textMask builds a boolean grid (rows x cols) where true marks a tile that
// belongs to the rendered text. Each line is scaled by scale (font pixel ->
// scale x scale tiles), centered horizontally, and the whole block of lines is
// centered vertically with lineGap font-pixels between lines. Tiles that fall
// outside the grid are clipped with a warning. bold thickens the glyph strokes
// by dilating the resulting mask bold times.
func textMask(lines []string, cols, rows, scale, gap, lineGap, bold int) [][]bool {
	mask := make([][]bool, rows)
	for r := range mask {
		mask[r] = make([]bool, cols)
	}

	advance := glyphW*scale + gap*scale
	lineH := glyphH * scale

	totalW, totalH := blockSize(lines, scale, gap, lineGap)
	if totalW > cols || totalH > rows {
		log.Printf("warning: text (%dx%d tiles) does not fully fit the grid (%dx%d); reduce -scale or grow the canvas",
			totalW, totalH, cols, rows)
	}

	originY := (rows - totalH) / 2

	for li, ln := range lines {
		runes := []rune(strings.ToUpper(ln))
		lineW := 0
		if len(runes) > 0 {
			lineW = len(runes)*advance - gap*scale
		}
		originX := (cols - lineW) / 2
		lineY := originY + li*(lineH+lineGap*scale)

		for i, ch := range runes {
			glyph, ok := font[ch]
			if !ok {
				glyph = font['?']
			}
			gx := originX + i*advance
			for fy := 0; fy < glyphH; fy++ {
				row := glyph[fy]
				for fx := 0; fx < glyphW; fx++ {
					if row[fx] == '0' || row[fx] == ' ' {
						continue
					}
					// paint a scale x scale block of tiles
					for sy := 0; sy < scale; sy++ {
						for sx := 0; sx < scale; sx++ {
							tx := gx + fx*scale + sx
							ty := lineY + fy*scale + sy
							if tx >= 0 && tx < cols && ty >= 0 && ty < rows {
								mask[ty][tx] = true
							}
						}
					}
				}
			}
		}
	}

	for i := 0; i < bold; i++ {
		mask = dilate(mask)
	}
	return mask
}

// dilate grows every true tile into its 8-connected neighbors. Repeated
// application thickens glyph strokes (the "bold" factor). Note that dilation
// also eats into letter counters, so for bold > 0 a scale of 2+ and a slightly
// larger -gap keep the text legible and letters separated.
func dilate(in [][]bool) [][]bool {
	rows := len(in)
	if rows == 0 {
		return in
	}
	cols := len(in[0])
	out := make([][]bool, rows)
	for y := range out {
		out[y] = make([]bool, cols)
	}
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if !in[y][x] {
				continue
			}
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					ny, nx := y+dy, x+dx
					if ny >= 0 && ny < rows && nx >= 0 && nx < cols {
						out[ny][nx] = true
					}
				}
			}
		}
	}
	return out
}

// ----------------------------------------------------------------------------
// Rendering
// ----------------------------------------------------------------------------

type config struct {
	text      string
	file      string
	width     int
	height    int
	tile      int
	scale     int
	gap       int
	lineGap   int
	bold      int
	palette   string
	textColor string
	dim       float64
	dimPad    int
	output    string
	seed      int64
	delay     int
	fluxFr    int
	revealFr  int
	freezeFr  int
}

func main() {
	cfg := config{}
	flag.StringVar(&cfg.text, "text", "HELLO", "text to render; use \\n for line breaks (ASCII letters, digits and . , ! ? - : ')")
	flag.StringVar(&cfg.file, "file", "", "read text from this file (one animation line per file line); overrides -text")
	flag.IntVar(&cfg.width, "width", 640, "canvas width in pixels")
	flag.IntVar(&cfg.height, "height", 360, "canvas height in pixels")
	flag.IntVar(&cfg.tile, "tile", 16, "tile size in pixels")
	flag.IntVar(&cfg.scale, "scale", 0, "font scale in tiles per font-pixel (0 = auto-fit)")
	flag.IntVar(&cfg.gap, "gap", 1, "gap between letters in font-pixels")
	flag.IntVar(&cfg.lineGap, "linegap", 1, "vertical gap between text lines in font-pixels")
	flag.IntVar(&cfg.bold, "bold", 0, "stroke thickness: number of dilation passes (0 = thin)")
	flag.StringVar(&cfg.palette, "palette", "neon", "palette: "+paletteNames())
	flag.StringVar(&cfg.textColor, "textcolor", "", "revealed-text color: #rrggbb hex or a name (white, black, red, green, blue, yellow, cyan, magenta); empty = palette default")
	flag.Float64Var(&cfg.dim, "dim", 1.0, "brightness of the background behind/around the text (0=black .. 1=no dimming)")
	flag.IntVar(&cfg.dimPad, "dimpad", 1, "padding in tiles around the text included in the dimmed region")
	flag.StringVar(&cfg.output, "output", "tileseq.gif", "output GIF filename")
	flag.Int64Var(&cfg.seed, "seed", 42, "random seed")
	flag.IntVar(&cfg.delay, "delay", 6, "frame delay in 1/100s")
	flag.IntVar(&cfg.fluxFr, "flux", 18, "number of fluctuation frames")
	flag.IntVar(&cfg.revealFr, "reveal", 28, "number of reveal frames")
	flag.IntVar(&cfg.freezeFr, "freeze", 30, "number of freeze frames")
	flag.Parse()

	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg config) error {
	rng := rand.New(rand.NewSource(cfg.seed))

	pal, ok := palettes[cfg.palette]
	if !ok {
		return fmt.Errorf("unknown palette %q; choose one of: %s", cfg.palette, paletteNames())
	}
	if cfg.tile <= 0 {
		return fmt.Errorf("tile size must be > 0")
	}

	cols := cfg.width / cfg.tile
	rows := cfg.height / cfg.tile
	if cols == 0 || rows == 0 {
		return fmt.Errorf("canvas too small for tile size: %dx%d / %d", cfg.width, cfg.height, cfg.tile)
	}

	lines, err := textLines(cfg)
	if err != nil {
		return err
	}

	// Auto-fit the font scale if the user did not pick one.
	scale := cfg.scale
	if scale <= 0 {
		scale = autoScale(lines, cols, rows, cfg.gap, cfg.lineGap)
	}

	mask := textMask(lines, cols, rows, scale, cfg.gap, cfg.lineGap, cfg.bold)

	// dimRegion marks tiles behind and around the text that get dimmed to make
	// the letters stand out. It is the text mask grown by dimPad tiles.
	var dimRegion [][]bool
	if cfg.dim < 1.0 {
		dimRegion = mask
		for i := 0; i < cfg.dimPad; i++ {
			dimRegion = dilate(dimRegion)
		}
	}

	// Color indices in the GIF palette:
	//   0                     -> background base (calm dark)
	//   1                     -> text color
	//   2 .. 2+N-1            -> flicker colors (bright)
	//   2+N .. 2+2N-1         -> dimmed flicker colors (behind the text)
	bgBase := color.RGBA{18, 18, 22, 255}
	textColor := pickTextColor(cfg.palette)
	if cfg.textColor != "" {
		c, err := parseColor(cfg.textColor)
		if err != nil {
			return err
		}
		textColor = c
	}

	gifPalette := color.Palette{bgBase, textColor}
	for _, c := range pal {
		gifPalette = append(gifPalette, c)
	}
	flickerStart := 2
	flickerN := len(pal)
	// Dimmed variants of each flicker color, indexed at settled+flickerN.
	for _, c := range pal {
		gifPalette = append(gifPalette, dimColor(c, cfg.dim))
	}

	anim := gif.GIF{LoopCount: 0}

	// Per-tile reveal order: a tile becomes "locked" to its final color once
	// the reveal progress passes its threshold. Text tiles reveal as text;
	// background tiles settle onto a fixed flicker color.
	threshold := make([][]float64, rows)
	settled := make([][]int, rows) // settled flicker color index per tile
	for y := 0; y < rows; y++ {
		threshold[y] = make([]float64, cols)
		settled[y] = make([]int, cols)
		for x := 0; x < cols; x++ {
			threshold[y][x] = rng.Float64()
			settled[y][x] = flickerStart + rng.Intn(flickerN)
		}
	}

	totalFrames := cfg.fluxFr + cfg.revealFr + cfg.freezeFr

	for f := 0; f < totalFrames; f++ {
		// progress in [0,1] across the reveal phase
		var progress float64
		switch {
		case f < cfg.fluxFr:
			progress = 0
		case f < cfg.fluxFr+cfg.revealFr:
			progress = float64(f-cfg.fluxFr+1) / float64(cfg.revealFr)
		default:
			progress = 1
		}

		img := image.NewPaletted(image.Rect(0, 0, cols*cfg.tile, rows*cfg.tile), gifPalette)

		for y := 0; y < rows; y++ {
			for x := 0; x < cols; x++ {
				var idx int
				locked := progress >= threshold[y][x]
				switch {
				case locked && mask[y][x]:
					idx = 1 // text
				case locked && dimRegion != nil && dimRegion[y][x]:
					// settled, dimmed halo behind the text
					idx = settled[y][x] + flickerN
				case locked:
					idx = settled[y][x] // settled background flicker color
				default:
					// still fluctuating
					idx = flickerStart + rng.Intn(flickerN)
				}
				fillTile(img, x, y, cfg.tile, uint8(idx))
			}
		}

		anim.Image = append(anim.Image, img)
		anim.Delay = append(anim.Delay, cfg.delay)
	}

	out, err := os.Create(cfg.output)
	if err != nil {
		return fmt.Errorf("create output: %w", err)
	}
	defer out.Close()

	if err := gif.EncodeAll(out, &anim); err != nil {
		return fmt.Errorf("encode gif: %w", err)
	}

	log.Printf("wrote %s: %d frames, %d line(s), grid %dx%d tiles, tile %dpx, scale %d, palette %q",
		cfg.output, totalFrames, len(lines), cols, rows, cfg.tile, scale, cfg.palette)
	return nil
}

// fillTile paints a single grid tile with the given palette index.
func fillTile(img *image.Paletted, col, row, tile int, idx uint8) {
	x0 := col * tile
	y0 := row * tile
	for dy := 0; dy < tile; dy++ {
		for dx := 0; dx < tile; dx++ {
			img.SetColorIndex(x0+dx, y0+dy, idx)
		}
	}
}

// autoScale picks the largest font scale such that all lines fit the grid.
func autoScale(lines []string, cols, rows, gap, lineGap int) int {
	for s := 12; s >= 1; s-- {
		w, h := blockSize(lines, s, gap, lineGap)
		if w <= cols && h <= rows {
			return s
		}
	}
	return 1
}

// textLines resolves the lines to render: from -file if given (one animation
// line per file line), otherwise from -text split on literal "\n". A trailing
// blank line from a file's final newline is dropped, but interior blank lines
// are kept so they appear as vertical spacing in the animation.
func textLines(cfg config) ([]string, error) {
	var raw string
	if cfg.file != "" {
		b, err := os.ReadFile(cfg.file)
		if err != nil {
			return nil, fmt.Errorf("read text file: %w", err)
		}
		raw = string(b)
	} else {
		raw = cfg.text
	}

	raw = strings.ReplaceAll(raw, "\r\n", "\n")
	raw = strings.ReplaceAll(raw, "\\n", "\n") // allow literal \n in -text
	raw = strings.TrimSuffix(raw, "\n")        // ignore one trailing newline
	lines := strings.Split(raw, "\n")
	if len(lines) == 0 {
		lines = []string{""}
	}
	return lines, nil
}

// dimColor scales an RGB color's brightness by f (clamped to [0,1]).
func dimColor(c color.RGBA, f float64) color.RGBA {
	if f < 0 {
		f = 0
	}
	if f > 1 {
		f = 1
	}
	return color.RGBA{
		R: uint8(float64(c.R) * f),
		G: uint8(float64(c.G) * f),
		B: uint8(float64(c.B) * f),
		A: 255,
	}
}

// parseColor accepts a "#rrggbb"/"rrggbb" hex value or one of a few common
// color names and returns it as an opaque RGBA.
func parseColor(s string) (color.RGBA, error) {
	named := map[string]color.RGBA{
		"white":   {255, 255, 255, 255},
		"black":   {0, 0, 0, 255},
		"red":     {255, 0, 0, 255},
		"green":   {0, 200, 0, 255},
		"blue":    {0, 0, 255, 255},
		"yellow":  {255, 255, 0, 255},
		"cyan":    {0, 255, 255, 255},
		"magenta": {255, 0, 255, 255},
	}
	if c, ok := named[strings.ToLower(s)]; ok {
		return c, nil
	}
	hex := strings.TrimPrefix(s, "#")
	if len(hex) != 6 {
		return color.RGBA{}, fmt.Errorf("invalid color %q: use #rrggbb or a name (white, black, red, green, blue, yellow, cyan, magenta)", s)
	}
	var r, g, b uint8
	if _, err := fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b); err != nil {
		return color.RGBA{}, fmt.Errorf("invalid hex color %q: %w", s, err)
	}
	return color.RGBA{r, g, b, 255}, nil
}

// pickTextColor returns a high-contrast text color for the given palette.
func pickTextColor(name string) color.RGBA {
	switch name {
	case "mono":
		return color.RGBA{255, 255, 255, 255}
	case "pastel":
		return color.RGBA{40, 40, 60, 255}
	default:
		return color.RGBA{245, 245, 245, 255}
	}
}
