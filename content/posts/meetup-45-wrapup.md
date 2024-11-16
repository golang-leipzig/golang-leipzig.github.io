---
title: "Hybrid Meetup #45 wrap-up"
date: 2024-11-14T17:00:00+01:00
draft: false
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
[#65653](https://github.com/golang/go/issues/65653), ...

Also, there is a currently open proposal:

* [proposal: cmd/cover: support branch coverage](https://github.com/golang/go/issues/70306)

> Adding **branch coverage** would be beneficial, as it would allow developers to
> better understand which branches (conditional paths) in their code are being
> executed during tests, rather than just which lines.

The [gocoverageplus](https://github.com/Fabianexe/gocoverageplus) addresses
some of these shortcomings and also supports
[cobertura](https://gcovr.com/en/stable/output/cobertura.html) output format, as well as complexity metrics.

A short usage demo:

```shell
$ go install github.com/Fabianexe/gocoverageplus@latest
$ go test -cover -coverprofile=c.out
```

Put a [config
file](https://github.com/Fabianexe/gocoverageplus/?tab=readme-ov-file#config)
into your projects folder, then:

```
$ gocoverageplus -i c.out -o p.out
```

Then render the report, e.g. as HTML or to stdout. A nice little TUI for
coverage reports is: [gocovsh](https://github.com/orlangure/gocovsh).

[![](/images/gocovsh.png)](https://github.com/orlangure/gocovsh)

### Go scheduler overlay

We took another look on testing, through the lens of [timestone](https://github.com/Metamogul/timestone),

>  a library to create deterministic and easy-to-understand unit tests for time-dependent, concurrent go code.

It requires modification of existing code (replace go routine invocations with
timestone), but after that, tests can be run in either with the system
scheduler, i.e. [using
goroutines](https://github.com/Metamogul/timestone/blob/7411decd9b3e1e28ef539e2bbb0ebb67b9e059d7/system/scheduler.go#L20-L29)
or a [simulation
scheduler](https://github.com/Metamogul/timestone/blob/7411decd9b3e1e28ef539e2bbb0ebb67b9e059d7/simulation/scheduler.go#L149-L155),
that gives the caller control over the process execution.

By default, Go programs can exhibit non-determinism in a few places, among others in:

* randomized work stealing in the [scheduler](https://github.com/golang/go/blob/8e714281e441f93d2865adb3c5a507fd161314e9/src/runtime/proc.go#L7198)
* randomized map [iteration](https://github.com/golang/go/blob/8b0ac33da8574b74ba50ad727b59fa8679d93e4b/src/internal/runtime/maps/map.go#L141-L142)
* randomized [select clause](https://github.com/golang/go/blob/8e714281e441f93d2865adb3c5a507fd161314e9/src/runtime/select.go#L181)

With [timestone](https://github.com/Metamogul/timestone), it is possible to get
around the scheduling randomness, when this interferes with test results. Check
out some of the examples included in the library, here:
[timestone/examples](https://github.com/Metamogul/timestone/tree/main/examples).


### Misc

* for integration testing, [testcontainers](https://golang.testcontainers.org/)
  can come in handy; you can simulate the actual services your code interfaces
  with, like key-value stores, object stores, search engines, databases, etc. -- a handy, albeit slightly slower, alternative to mock objects
* for a more unusual test setup, see this lightning talk about [running go tests in a VM w/ qemu](https://www.youtube.com/watch?v=69Zy77O-BUM)

