---
title: "Virtual Meetup #23 wrap-up"
date: 2021-11-26T14:00:00+01:00
draft: false
tags:
- summary
- meetup
---

## Summary

Meetup #23, Nov 23, 2021, originally planned on-site at
[Sciendis](https://www.sciendis.de/), moved online again, due to the [pandemic
regulations](https://www.coronavirus.sachsen.de/download/SaechsCoronaNotVO-2021-11-19.pdf).

We had two input presentations, one about *Go in a startup environment* at
Sciendis and reasons why Go may be a good choice. And one lightning talk about
[a data web service](https://github.com/miku/dwstalk).

## Go in a startup environment

Sciendis develops a React Native application called
[Wundera](https://wundera.health/) with the backend service being written in
Go, in a classic cloud setup. Moving from six to two services (user, patient)
proved beneficial, reduced code duplication and also lowered the data
serialization overhead.

A few points why Go has been attractive:

* easier language with fewer constructs results in less (cognitive) overhead; faster onboarding, maybe less code
* Go compiles to native Code, minimal requirements on deploy target
* Go garbage collector is minimalistic as well (see e.g. [Go GC
  Settings](https://archive.fosdem.org/2019/schedule/event/gogc/attachments/slides/3134/export/events/attachments/gogc/slides/3134/Go_GC_Settings_Bryan_Boreham.pdf))
* Go is still a fast language
* approachable concurrency
* community agreement over code style (no more code reviews discussion
  important, yet shallow problems of style)

Interestingly, with Go you learn and use more the language itself (and its
standard library) than a specific framework.

## A data web service

A quick overview on a small web service built on [sqlite](https://sqlite.org),
[net/http](https://pkg.go.dev/net/http) and
[gorilla](https://www.gorillatoolkit.org/) handlers.

[![](https://github.com/miku/dwstalk/raw/main/static/Untitled-2020-06-15-1740.png)](https://github.com/miku/dwstalk)

Notes: [A data web service](https://github.com/miku/dwstalk).

## Misc

* Java is drafting a lightweight concurrency model as an alternative to
  threads, see:
[https://openjdk.java.net/jeps/8277131](https://openjdk.java.net/jeps/8277131)

> Drastically reduce the effort of writing, maintaining, and observing
> high-throughput concurrent applications that make the best use of available
> hardware through virtual threads, a lightweight user-mode thread
> implementation with dramatically reduced costs.

Interestingly, the argument of mismatch between application and language view on what the concurrent unit is, pops up:

> This results in the asynchronous style of programming, that not only requires
> a separate and incompatible set of APIs, but breaks the connection between
> the logical application unit (transaction) and the platform's unit (thread),
> which makes the platform unaware of the application's logical units.

This argument appears also in [Concurrency in
Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/)
(which we discussed in [meetup #7](https://golangleipzig.space/posts/meetup-7-invitation/)).

* For ideas on idiomatic Go, see: [Effective Go](https://go.dev/doc/effective_go)
* Go has it dark corners, too - we are collecting some of them in [Go Gotchas](https://github.com/golang-leipzig/gotchas)
* End-to-end CI pipeline with GH
  [actions](https://github.com/features/actions),
[goreleaser](https://github.com/goreleaser/goreleaser) and
[ko](https://github.com/google/ko):
[https://github.com/embano1/ci-demo-app](https://github.com/embano1/ci-demo-app)
* Test options:
  [https://github.com/stretchr/testify](https://github.com/stretchr/testify),
for integration tests [test
containers](https://github.com/testcontainers/testcontainers-go) can be used
(example test suite [starting elasticsearch
servers](https://github.com/miku/esbulk/blob/69db0fab576ff53ef40653f8696b3942e6e26734/run_test.go#L78-L106)
to test an indexing tool).
* In general, having you environment wrapped in containers (and a
  docker-compose to set them up) is a popular technique
([discussion](https://news.ycombinator.com/item?id=25859588)- VS Code supports
it, too:
[https://code.visualstudio.com/docs/remote/containers](https://code.visualstudio.com/docs/remote/containers)
- especially via a `devcontainers.json` file in your repo:

> A devcontainer.json file in your project tells VS Code how to access (or
> create) a development container with a well-defined tool and runtime stack.

* Productive Postgres with Go: [GopherCon 2020: Johan Brandhorst-Satzkorn - A Journey to Postgres Productivity with Go](https://www.youtube.com/watch?v=AgHdVPSty7k)
* ORM tools seem to be a bit less used with Go, but options are [GORM](https://gorm.io/index.html), or also [ent](https://entgo.io/)
* [sqlite](https://sqlite.org/) is an awesome database, and since it's a stable platform, all kinds of
  project develop around or on top of it, like [streaming
replication](https://github.com/benbjohnson/litestream), or
[trough](https://github.com/internetarchive/trough); video recommendation: [David Crawshaw SQLite and Go](https://www.youtube.com/watch?v=RqubKSF3wig)
* For mocking time in tests: [https://github.com/benbjohnson/clock](https://github.com/benbjohnson/clock)
* Domain-driven design is a popular technique, originating from Eric Evans 2003
  book [Domain Driven Design](https://openlibrary.org/works/OL4464385W), with a
good summary being: [DDD
Distilled](https://openlibrary.org/works/OL19546213W/)
* If projects get bigger, put an `ARCHITECTURE.md` into the repo, e.g. [like
  this](https://github.com/kstenerud/go-concise-encoding/blob/master/ARCHITECTURE.md).
* Tabular test strategy to decouple language from *domain* content ([example](https://gitlab.com/internetarchive/refcat/-/blob/master/skate/testdata/verify.csv) dealing with JSON document comparisons).
* Test augmentation library for Go: [https://pkg.go.dev/gotest.tools/assert](https://pkg.go.dev/gotest.tools/assert); also useful in tests to compare compound values: [google/go-cmp](https://github.com/google/go-cmp)
* Writeup on generics in Go: [Go Generics - A Quick Overview](https://www.klingt.net/articles/go-generics-a-quick-overview.html)

Video recommendations:

* [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)
* [dotGo 2019 - Kat Zień - Achieving maintainability with hexagonal architecture](https://www.youtube.com/watch?v=vKbVrsMnhDc)
* [David Crawshaw SQLite and Go](https://www.youtube.com/watch?v=RqubKSF3wig)
* Channel: [https://www.youtube.com/c/MarcelDempers](https://www.youtube.com/c/MarcelDempers)

----

[Join our meetup](https://www.meetup.com/Leipzig-Golang) to get notified of
upcoming events!

