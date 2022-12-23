---
title: "Meetup #3 wrap-up"
date: 2019-04-14T23:15:00+02:00
draft: false
tags:
- meetup
- summary
---

[Meetup #3](https://www.meetup.com/Leipzig-Golang/events/260338152/)
took place on Friday, April 12, 2019, 19:00 CEST at
[Basislager](https://www.basislager.co/). We are now officially part of the [Go
Developer Network](https://www.meetup.com/pro/go/) (GDN, currently 103 groups with 47418 members in 37 countries), which was [announced
in March 2019](https://blog.golang.org/go-developer-network).

> The GDN is a collection of Go user groups working together with a shared
> mission to empower developer communities with the knowledge, experience, and
> wisdom to build the next generation of software in Go.

{{< h2 >}}Meetup topics{{< /h2 >}}

We took a closer look at [Go
modules](https://github.com/golang/go/wiki/Modules) for dependency management
by highlighting a few passages from the [Go and versioning design
documents](https://research.swtch.com/vgo). The
[repository](https://github.com/golang-leipzig/gomodintro) and
[Slides](https://github.com/golang-leipzig/gomodintro/blob/master/Slides.md)
([PDF](https://github.com/golang-leipzig/gomodintro/blob/master/Slides.pdf))
are online.

Furthermore, we had a demonstration of some strange behaviour in the
[gorilla/handlers](https://github.com/gorilla/handlers) package:
[gorilla/handlers.CompressHandler](https://github.com/gorilla/handlers/blob/ac6d24f88de4584385a0cb3a88f953d08a2f7a05/compress.go#L57-L64)
might gzip twice! If you want to see the bug in action,
[klingtnet](https://github.com/klingtnet) got you covered:

* [spreadshirt/gorilla-handlers-double-gzip-bug](https://github.com/spreadshirt/gorilla-handlers-double-gzip-bug)

Interestingly, the
[nytimes/gziphandler](https://github.com/nytimes/gziphandler) does not exhibit
the double compression behavior. The solution becomes a bit clearer, when you take into account,
that each HTTP middleware will act on its own and has limited means to pass
information to another handler (e.g. via HTTP headers). If the middleware acts on a
header (e.g.
[Accept-Encoding](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Encoding)),
then it can choose to modify the request (it's [a
pointer](https://golang.org/pkg/net/http/#HandlerFunc.ServeHTTP), in contrast
to the response writer).

{{< h2 >}}From Go to Cloud to Linux{{< /h2 >}}

Clouds and Cloud-Native technologies would be nothing without Linux. The [Linux
Meetup Leipzig](https://www.meetup.com/Linux-Meetup-Leipzig/) hosts interesting
talks, e.g. a series on container technologies (by [Sascha
Grunert](https://github.com/saschagrunert)) or on [tunnels and
honeypots](https://www.meetup.com/Linux-Meetup-Leipzig/events/260563903/).
Since there is overlap in topics, we might organize a joint event in the
future.

{{< h2 >}}Postponed topics{{< /h2 >}}

* worker pool implementation benchmarks (from [#2](https://golangleipzig.space/posts/second-meetup-wrapup/))
* Ramen Linux lightning talk (Update: `https://ramenlinux.com` is no more)

{{< h2 >}}References{{< /h2 >}}

* [github.com/golang-leipzig/gomodintro](https://github.com/golang-leipzig/gomodintro)
* [github.com/spreadshirt/gorilla-handlers-double-gzip-bug](https://github.com/spreadshirt/gorilla-handlers-double-gzip-bug)
* The [gorilla toolkit](https://www.gorillatoolkit.org/)
* [github.com/nytimes/gziphandler](https://github.com/nytimes/gziphandler)

The tool for inspecting Go binaries and list their deps is called [goversion](https://github.com/rsc/goversion):


    $ go get -u rsc.io/goversion


### Misc

* [Lightning Talk: Brad Fitzpatrick - The nuclear option, go test -run=InQemu](https://www.youtube.com/watch?v=69Zy77O-BUM) (2018)
* A blog post by Russ Cox on the [software dependency
  problem](https://research.swtch.com/deps), which reflect on the many
different ways, languages manage dependencies today.

> Dependency managers now exist for essentially every programming language.
> Maven Central (Java), Nuget (.NET), Packagist (PHP), PyPI (Python), and
> RubyGems (Ruby) each host over 100,000 packages. The arrival of this kind of
> fine-grained, widespread software reuse is one of the most consequential
> shifts in software development over the past two decades. And if we’re not
> more careful, it will lead to serious problems.

