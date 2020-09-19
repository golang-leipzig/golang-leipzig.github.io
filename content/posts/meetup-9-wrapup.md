---
title: "Meetup #9 wrap-up"
date: 2020-02-22T01:00:00+02:00
draft: false
tags:
- summary
- meetup
---

{{< h2 >}}Go 1.14{{< /h2 >}}

[Meetup #9](https://www.meetup.com/Leipzig-Golang/events/268785494/)
took place on Friday, February 21, 2020, 19:00 CEST at
[Basislager](https://www.basislager.co/).

We looked at changes coming in the upcoming [Go
1.14](https://tip.golang.org/doc/go1.14) release, which is is expected to be
released in February 2020.

The slides can be found here:

* [Slides](https://rawcdn.githack.com/golang-leipzig/go-1.14-and-beyond/e430eef8ac39a2a3eb15a4d422bebe2069e05960/slides.html), [Repo](https://github.com/golang-leipzig/go-1.14-and-beyond)

[![](/images/go1.14andbeyond.png)](https://rawcdn.githack.com/golang-leipzig/go-1.14-and-beyond/750e27aec11e1fa49064173fc6512d7d79515e84/slides.html#/title-slide)

There are many improvements, just one of them is lower-cost [defer
statements](https://golang.org/ref/spec#Defer_statements).

The proposal can be found here:

* [Proposal: Low-cost defers through inline code, and extra funcdata to manage
  the panic
  case](https://go.googlesource.com/proposal/+/refs/heads/master/design/34481-opencoded-defers.md)

Recommended podcast: [gotime/112](https://changelog.com/gotime/112).

----

{{< h2 >}}HTTP Getaway{{< /h2 >}}

The [net/http](https://golang.org/pkg/net/http/) package offers many extension points, using interfaces or callbacks - as well as alternative implementations. A few examples are sketched out in [HTTP Getaway](https://github.com/miku/httpgetaway).

![](https://raw.githubusercontent.com/miku/httpgetaway/master/static/hijack.gif)

Retry is a common pattern. Interestingly, application layer retries might not
always be enough, sometimes only a connection reset will help.

![](https://raw.githubusercontent.com/miku/httpgetaway/master/static/levels.png)

----

{{< h2 >}}Misc{{< /h2 >}}

* We discussed various ways to handle errors, there has been some updates
  starting with Go 1.13: [Working with Errors in Go
  1.13](https://blog.golang.org/go1.13-errors)
* A popular testing library is [stretchr/testify](https://github.com/stretchr/testify), e.g. helps to test code that panics. Maybe more lightweight - [is](https://github.com/matryer/is).
* Speaking of panics: Shall libraries panic? In which cases? According to
  [Effective Go](https://golang.org/doc/effective_go.html#panic), `panic`
  should be an exception, rather the rule.
* Pop quiz on modules: Is
  [go.sum](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more)
immutable? The [docs
say](https://golang.org/cmd/go/#hdr-Module_authentication_using_go_sum): "In
day-to-day development, the checksum of a given module version should never
change." - however it *can* change, here's an example commit:
[766e40eb](https://github.com/kubernetes-sigs/kind/commit/766e40ebe4c6ca4cf8c845cba3a4a6d51e94fc9b#diff-f949e2d81c8076ebbf8af38fcbb72c1f).
More details can be found a blog series on modules, e.g. [Part 4: Mirrors,
Checksums and
Athens](https://www.ardanlabs.com/blog/2020/02/modules-04-mirros-checksums-athens.html).
* Three HTTP router libraries that seem to be popular:
  [mux](https://github.com/gorilla/mux), [chi](https://github.com/go-chi/chi),
  [httprouter](https://github.com/julienschmidt/httprouter)
* Combining [io.Reader](https://golang.org/pkg/io/#Reader) and
  [Context](https://golang.org/pkg/context/#Context) can be useful to cancel an expensive read (e.g. large file) in an HTTP handler, here's a description of [context-aware IO](https://pace.dev/blog/2020/02/03/context-aware-ioreader-for-golang-by-mat-ryer).
* One cool project: [kind](https://github.com/kubernetes-sigs/kind), providing local clusters for testing Kubernetes. Given a docker with at least Go 1.11, starting a cluster becomes a one-liner: `GO111MODULE="on" go get sigs.k8s.io/kind@v0.7.0 && kind create cluster` - incredible.
* 16 days ago, a new W3C recommendation was published, [Trace Context Level
  1](https://www.w3.org/TR/trace-context/), which *defines
  standard HTTP headers and a value format to propagate context information
  that enables distributed tracing scenarios. The specification standardizes
  how context information is sent and modified between services. Context
  information uniquely identifies individual requests in a distributed system
  and also defines a means to add and propagate provider-specific context
  information.*

And much more.

----

{{< h2 >}}Contributing{{< /h2 >}}

We want to make it simpler to contribute and to stay up to date with our meetup
and we may create a mailing list in the future.

Until then, please feel invited to:

* follow our [RSS feed](https://golangleipzig.space/posts/index.xml)
* [PR a topic](https://github.com/golang-leipzig/topics)
* contact us via [Meetup](https://www.meetup.com/Leipzig-Golang/) or [E-mail](mailto:martin.czygan@gmail.com)
* join us at GitHub at [Leipzig Gophers](https://github.com/golang-leipzig) org
