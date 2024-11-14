---
title: "Hybrid Meetup #45 wrap-up"
date: 2024-11-12T10:00:00+01:00
draft: true
tags:
- summary
- meetup
---

## Testing in Go

Hybrid Meetup #45 took place
[2024-10-29](https://www.meetup.com/leipzig-golang/events/298481354/) 19:00 at
[CHECK24](https://check24.de) Leipzig Office at Altes Landratsamt and we were
thrilled to see people joining from Leipzig
[and](https://en.wikipedia.org/wiki/Saxony-Anhalt)
[beyond](https://en.wikipedia.org/wiki/Thuringia), both on site and online. We
had two great input presentations about open source projects in the testing
domain.

### Coverage metrics

[Fabian](https://www.linkedin.com/in/fabian-g%C3%A4rtner-913584141/) analyzed
the Go coverage tool [cmd/cover](https://pkg.go.dev/cmd/cover), listed some of
its shortcomings and demonstrated a tool to address them and to improve some
metrics: [gocoverageplus](https://github.com/Fabianexe/gocoverageplus).

The coverage tool shipped with Go operates on the source code and uses
instrumentation to estimate coverage:

> When generated instrumented code, the cover tool computes approximate
basic block information by studying the source. It is thus more portable than
binary-rewriting coverage tools, but also a little less capable. For instance,
it does not probe inside && and || expressions, and can be mildly confused by
single statements with multiple function literals. -- [cmd/cover](https://pkg.go.dev/cmd/cover)

In the past, go cover had a couple of issues, among them:
[#23883](https://github.com/golang/go/issues/23883),
[#51430](https://github.com/golang/go/issues/51430),
[#58770](https://github.com/golang/go/issues/58770),
[#65570](https://github.com/golang/go/issues/65570),
[#65653](https://github.com/golang/go/issues/65653)

Also, there is a currently open proposal:

* [proposal: cmd/cover: support branch coverage](https://github.com/golang/go/issues/70306)

> Adding **branch coverage** would be beneficial, as it would allow developers to
> better understand which branches (conditional paths) in their code are being
> executed during tests, rather than just which lines.

The [gocoverageplus](https://github.com/Fabianexe/gocoverageplus) addresses
some of these shortcomings and also supports
[cobertura](https://gcovr.com/en/stable/output/cobertura.html) output format.

A short usage demo:

```shell
$ go install github.com/Fabianexe/gocoverageplus@latest
$ go test -cover -coverprofile=cover.out
```

Put a [config
file](https://github.com/Fabianexe/gocoverageplus/?tab=readme-ov-file#config)
into your projects folder, then:

```
$ gocoverageplus -i cover.out -o plus.out
```

### Go scheduler overlay

