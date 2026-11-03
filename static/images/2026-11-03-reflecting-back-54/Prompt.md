Task: Write a Go program "quickindex" that can be compiled to WASM later that
would allow us to provide search over the content of this site.

Write a Go program that can be fed a list of of markdown files; it builds an
inverted index and allows to use a query to get back a list of ranked results.
This should work on the command line first. The list of files can be provided
through stdin, e.g. through a pipe:

    fd -e md | quickindex -o my.idx

Then query

    quickindex -f my.idx QUERY

Keep it simple and use stdlib only, so we could run this simply by using `go
run quickindex.go` - also try to put all code into a single file.

Search should be case insensitive.
