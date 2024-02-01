---
title: "Hybrid Meetup #40 wrap-up"
date: 2024-01-31T10:00:00+01:00
draft: false
tags:
- summary
- meetup
---

## Domain Driven Data Oriented Design

Hybrid Meetup #40 took place
[2024-01-30](https://www.meetup.com/leipzig-golang/events/298481328/) at
[Basislager Leipzig](https://www.basislager.co/) and we had a great
presentation by [Bill Kennedy](https://twitter.com/goinggodotnet) from [Ardan Labs](https://www.ardanlabs.com/) about Domain Driven Data
Oriented Design. A recording of the talk is available at our [YouTube Channel](https://www.youtube.com/@golangleipzig1131).

[![](/images/meetup-40-youtube-thumb.png)](https://youtu.be/zuLsdP3i8sQ)

All code shown in the talk is available at:
[ardanlabs/service](https://github.com/ardanlabs/service).

Just a few notes from the talk:

* [CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) is boilerplate
* [domain-driven design](https://en.wikipedia.org/wiki/Domain-driven_design) comes with a cost but can help to keep larger projects structured and maintainable
* one way to limit data exposure per domain is to create [a relational view](https://en.wikipedia.org/wiki/View_(SQL)) per domain
* example of an opaque type from the repo is
  [`user.Role`](https://github.com/ardanlabs/service/blob/387cc15defe9c9e01ca130118c8e01ed1a164844/business/core/user/role.go#L17-L20) - it could be a string, but we limit its usage by using a parse function to
ensure validity
* code that reads code to generate docs: [webapi.go](https://github.com/ardanlabs/service/blob/387cc15defe9c9e01ca130118c8e01ed1a164844/app/tooling/docs/webapi/webapi.go)

Thanks again to Bill for sharing his perspective on domain driven design and
how it could be applied in Go.

## Misc

* Go has first-class self-parsing support in the
  [ast](https://pkg.go.dev/go/ast) package, with the canonical example of code
working on code being [go
fmt](https://cs.opensource.google/go/go/+/refs/tags/go1.21.6:src/cmd/gofmt/gofmt.go);
in [#36](https://golangleipzig.space/posts/meetup-36-wrapup/) we learned about
[encore.dev](https://github.com/encoredev/encore), which makes extensive use of
the Go AST to infer and validate a complete service topology
* [Parse, don't
  validate](https://lexi-lambda.github.io/blog/2019/11/05/parse-don-t-validate/)
(Haskell ahead) - how to use data structures to make illegal state
unrepresentatable and more
* an [opaque type](https://en.wikipedia.org/wiki/Opaque_data_type) for [time](https://github.com/golang/go/blob/ae457e811d44261801bda261731b5006d629930d/src/time/time.go#L135-L156) in the
  standard library helped with switching the underlying implementation with the 2017 proposal for [*Monotonic
Elapsed Time Measurements in
Go*](https://go.googlesource.com/proposal/+/master/design/12914-monotonic.md) - an issue that had grave effects: [Other reported software problems associated
with the leap second](https://en.wikipedia.org/wiki/Leap_second#Other_reported_software_problems_associated_with_the_leap_second)

> Its [Cloudflare's] DNS resolver implementation incorrectly calculated a
> negative number when subtracting two timestamps obtained from the Go
> programming language's [time.Now()](https://pkg.go.dev/time#Now) function,
> which then used only a real-time clock source. This could have been avoided
> by using a monotonic clock source, which has since been added to Go [1.9](https://go.dev/doc/go1.9#monotonic-time).

## [CGMR](https://en.wikipedia.org/wiki/Central_German_Metropolitan_Region) and beyond

* on [2024-02-26](https://www.meetup.com/webwirtschaft/events/298927799), there will be an *Intro to Go* (golang) talk by [Karl](https://karlbreuer.com/) at [webmontag Halle](https://webwirtschaft.net/webmontag/), at [Mitteldeutsches Multimediazentrum Halle](https://www.mmz-halle.de/)
* if you are in Hannover, Germany, check out [Hannover Gophers](https://www.meetup.com/de-DE/hannover-gophers/)

[![](/images/clean_475668308_30.webp)](https://www.meetup.com/de-DE/hannover-gophers/)

