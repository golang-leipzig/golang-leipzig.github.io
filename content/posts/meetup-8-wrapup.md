---
title: "Meetup #8 wrap-up"
date: 2019-11-19T21:00:00+02:00
draft: false
tags:
- summary
- meetup
---

{{% h2 %}}Anatomy of a Go module proxy{{% /h2 %}}

The Go packaging story started many years ago with `GO15VENDOREXPERIMENT` and
a myriad of tools for managing dependencies. And, annoyingly `GOPATH` was one
of the main obstactles for people getting started with Go. Since [Go
1.11](https://golang.org/doc/go1.11#modules), we have experimental support for
Go modules.

An new component related to modules is the [module
proxy](https://proxy.golang.org/). A bit of history, internals and pitfalls has
been compiled by [klingtnet](https://github.com/klingtnet) into a great talk.

[Go Module Proxy - Internals and
Pitfalls](https://rawcdn.githack.com/golang-leipzig/module-proxy-november-meetup/c58770084e025fd4a1ea0f6b7f69c5adeae1817a/slides.html#/title-slide).
The repository for can be found at
[golang-leipzig/module-proxy-november-meetup](https://github.com/golang-leipzig/module-proxy-november-meetup).

Proxies are great, as they lower the load on source repository hosts and can
prevent events like [left-pad](http://left-pad.io/) (2016). However, there is
a slight potential information leak when the Go tools contact
[proxy.golang.org](https://proxy.golang.org/) - the `GOPRIVATE` [environment
variable](https://golang.org/doc/go1.13#modules) [can
help](https://github.com/golang/go/issues/33796).

Advantages, implementations and the proxy protocal details can be found in the
[slides](https://rawcdn.githack.com/golang-leipzig/module-proxy-november-meetup/c58770084e025fd4a1ea0f6b7f69c5adeae1817a/slides.html#/title-slide).

![](/proxy-protocol.png)

Also, thanks [klingtnet](https://github.com/klingtnet) for liberating an slightly
strange zip implementation from the depths of the Go module proxy
implementation!

* This issue started the conversation: [golang/go/issues/34953](https://github.com/golang/go/issues/34953)
* And a few weeks later, we have: [golang.org/x/mod/zip](https://godoc.org/golang.org/x/mod/zip)

{{% h2 %}}Pull Request Deployment{{% /h2 %}}

Everyone needs this. When a change to a codebase is made, how great would it be
to just see that single change, e.g. for a given pull request? That's what this
pull request deployment tool does. While the tool itself is tied to company
context, [panzerdev](https://github.com/panzerdev) condensed the workings of the tool into these slides:

* [Pull_Request_Deployment.pdf](/downloads/Pull_Request_Deployment.pdf)

Motivation:

> * Code changes in a branch should be tested in isolation based on latest master
> * Testing of not yet approved changes shouldn't change shared datasource (DB)
> * Multiple code changes shouldn't share a datasource (DB) simultaniously
> * Should be gone after merge without leftover junk

Read the slides, get inspired, join us next time - and: automate all the things!

{{% h2 %}}Misc{{% /h2 %}}

* Go turned 10 this month - [Happy Birthday](https://blog.golang.org/10years) and thanks for bringing fun back into programming!
* A new site emerged: [pkg.go.dev/](https://pkg.go.dev/)

> Go.dev is a companion website to golang.org. Golang.org is the home of the
open source project and distribution, while go.dev is the hub for Go users
providing centralized and curated resources from across the Go ecosystem.

* [GOLAB](https://golab.io/) 2019 was great - a truly international conference with over 2/3 of the people coming from all over the world!
* For upcoming conferences, consult the wiki: [Go Conferences and Major Events](https://github.com/golang/go/wiki/Conferences)

