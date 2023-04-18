---
title: "Hybrid Meetup #35 wrap-up"
date: 2023-04-19T01:00:00+02:00
draft: false
tags:
- summary
- meetup
---

## Cgo experiments

Hybrid meetup #35 took place [2023-04-18
19:00](https://www.meetup.com/leipzig-golang/events/290666173/) at [Basislager
Leipzig](https://www.basislager.co/) and was all about [Cgo](https://pkg.go.dev/cmd/cgo).

Cgo allows to bridge C and Go. From C, we can connect to other languages, like
C++, too. Use cases are legacy code or performance. Implementing a
[pseudo-random
walk](https://github.com/miku/cgosamples/blob/06f9ab34f0aef71f865872bfcfd391a01a626a0d/x/xcgoloop/main.go#L1-L95)
in both Go and C shows that C can be up to 4x faster for this particular task.

```
 C               1              -1 29.553µs
GO               1              -1 84.303µs
 C            1000             -12 41.799µs
GO            1000              26 107.295µs
 C         1000000           -1810 22.837387ms
GO         1000000              -2 83.182797ms
 C      1000000000           15894 6.25975724s
GO      1000000000          -44632 25.620427056s
```

One example was concerned with wrapping
[ggml](https://github.com/ggerganov/ggml), a lightweight *tensor library for
machine learning* -
[xcgoggml](https://github.com/miku/cgosamples/tree/main/x/xcgoggml). A
threading [mystery
remains](https://github.com/miku/cgosamples/tree/main/x/xcgothreads).

Experiment details can be found here: [https://github.com/miku/cgosamples](https://github.com/miku/cgosamples)

## Zig cross-compilation

The zig compiler (LLVM) can build real staticly linked binaries from a Go
project with C dependencies, such as sqlite3. Example repo:
[https://github.com/klingtnet/cross-compile-with-zigcc](https://github.com/klingtnet/cross-compile-with-zigcc).


## Misc

* Cgo is a not Go, many guarantees are given up, maintenance of mixed C and Go
  code is much more difficult - rather than use CGO, maybe a grpc or other
interface may keep C and Go world better separated
* To wrap C++, we need to create a C [façade](https://en.wikipedia.org/wiki/Facade_pattern)
* Go has a nice [ast](https://pkg.go.dev/go/ast)
* GRCP performance in Go may suffer, if interfaces are used (due to overhead of reflection)
* [oapi-codegen](https://github.com/deepmap/oapi-codegen) is an openapi code generator for Go (and an alternative to the [swagger tool](https://github.com/OpenAPITools/openapi-generator))
* Go's type system can be limiting, when working with Web APIs or SDKs
* [Simple Markup Language](https://www.simpleml.com/)
* [Parsec](https://hackage.haskell.org/package/parsec), [Participle](https://github.com/alecthomas/participle)
* [YAML/HELL](https://ruudvanasseldonk.com/2023/01/11/the-yaml-document-from-hell)
* Using symbols straight from a shared object via [dlopen](https://stackoverflow.com/a/27510758/89391)
* A rundown of a benchmark across Python, C, and BLAS, for a HPC and competetive programming blog: [https://en.algorithmica.org/hpc/complexity/languages/](https://en.algorithmica.org/hpc/complexity/languages/)

We also briefly discussed
[LLMs](https://en.wikipedia.org/wiki/Large_language_model), as we ran
[alpaca](https://crfm.stanford.edu/2023/03/13/alpaca.html) [model
ggml-alpaca-7b-q4.bin](https://news.ycombinator.com/item?id=35191137) in the
background, which wrote some simple SQL and thought of a haiku about the Go programming language:

![](/images/haiku.png)

Did you find a nice haiku?

[Join our meetup](https://www.meetup.com/Leipzig-Golang/) and let us know!

