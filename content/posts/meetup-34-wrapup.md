---
title: "Hybrid Meetup #34 wrap-up"
date: 2023-02-22T12:00:00+02:00
draft: false
tags:
- summary
- meetup
---

## Go 1.20

Hybrid meetup #34 took place [2023-02-21
19:00](https://www.meetup.com/leipzig-golang/events/290666161/) at [Basislager
Leipzig](https://www.basislager.co/) and was all about [Go
1.20](https://tip.golang.org/doc/go1.20) (similar to previous years, where we
looked at Go 1.14, 1.16, 1.18 - [#9](https://golangleipzig.space/posts/meetup-9-wrapup/), [#16](https://golangleipzig.space/posts/meetup-16-wrapup/), [#25](https://golangleipzig.space/posts/meetup-25-wrapup/)).

![](/images/go1.20latest.png)

We had two great presentations:

* A rundown of most novelties in Go 1.20: [Go 1.20
  updates](https://gist.github.com/klingtnet/0b5fd6cd742fe030115d3cc10776531f), including many examples -- thanks, [Andreas](https://www.klingt.net/)! ... and
* [Go 1.20 Coverage Profiling Support for Kubernetes Apps](https://www.mgasch.com/2023/02/go-e2e/), a walkthrough on how to use Go 1.20 binary profiling with containerized workloads (with inception style isolation layer qualities) -- thanks, [Michael](https://twitter.com/embano1)!

Some highlights from the release notes:

* [context.WithCancelCause](https://pkg.go.dev/context@master#WithCancelCause) allows to pass a custom error along the cancellation line
* package http gained a wrapper around the `http.ResponseWriter`, called [`http.ResponseController`](https://pkg.go.dev/net/http#ResponseController) allowing for per request timeout controls

Go started to [add support](https://go.dev/doc/pgo) for [Profile-guided optimization](https://en.wikipedia.org/wiki/Profile-guided_optimization), a technique already in place in various technologies.

> As of Go 1.20, benchmarks for a representative set of Go programs show that
> **building with PGO improves performance by around 2-4%**. We expect performance
> gains to generally increase over time as additional optimizations take
> advantage of PGO in future versions of Go. -- [https://go.dev/doc/pgo](https://go.dev/doc/pgo)

A challenge in containerized environments (e.g. [k8s](https://kubernetes.io/))
is that profiling information may be written to emphemeral storage, hence lost,
once the container is torn down.

At the same time, end-to-end (e2e) tests are an important technique to test
software, especially when running in complex enviroments with many interactions
(e.g. with other services, backends, ...). You do not want to mock everything (too tedious), and
[containers](https://cloud.google.com/learn/what-are-containers) are a popular way to isolate components, anyway.

The example project we looked at was a [vSphere](https://github.com/embano1/vsphere) client
library, that uses [e2e
tests](https://github.com/embano1/vsphere/blob/main/.github/workflows/e2e.yaml)
running in a local cluster using [kind](https://kind.sigs.k8s.io/), and
[ko](https://github.com/ko-build/ko) to turn Go programs into images,
[automagically](https://ko.build/) (excellent developer tools, btw).

If you [mount a local
path](https://github.com/embano1/vsphere/blob/545c3a4658b945fcaa1fef4b815cf7e92079fb9f/.github/workflows/e2e.yaml#L57-L68)
into kind, and mount that again [into the
pod](https://github.com/embano1/vsphere/blob/545c3a4658b945fcaa1fef4b815cf7e92079fb9f/test/client_test.go#L221-L231)
running the tests and let the Go binary know, [where to write the cover
information](https://github.com/embano1/vsphere/blob/545c3a4658b945fcaa1fef4b815cf7e92079fb9f/test/client_test.go#L195)
(via [GOCOVERDIR](https://go.dev/testing/coverage/#running)) - you can cover
any environment that has a local filesystem exposed (e.g. your laptop, or CI,
...).

It's not easy, but it works - and impressive as for the [fidelity](https://abseil.io/resources/swe-book/html/ch14.html) of the test
environment.

Quick Github Actions tip: To not exhausting CI limits on GH actions, cancel previous runs on a
PR branch, if you push new commits - see [this
example](https://github.com/embano1/vsphere/blob/545c3a4658b945fcaa1fef4b815cf7e92079fb9f/.github/workflows/e2e.yaml#L9-L12).

Thanks to [Andreas](https://klingt.net) and
[Michael](https://twitter.com/embano1) for the insights.

Go 1.20 continues to add small improvements to the language and ecosystem and
more interesting things are in the pipeline, like the experimental structured
logging [slog](https://pkg.go.dev/golang.org/x/exp/slog).

What's your favorite feature?

[Join our meetup](https://www.meetup.com/Leipzig-Golang/) and let us know!

