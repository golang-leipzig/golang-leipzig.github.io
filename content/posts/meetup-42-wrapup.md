---
title: "Hybrid Meetup #42 wrap-up"
date: 2024-05-03T10:00:00+01:00
draft: false
tags:
- summary
- meetup
---

## Fast Things First

Hybrid Meetup #42 took place
[2024-04-30](https://www.meetup.com/leipzig-golang/events/298066357) at
[CHECK24](https://check24.de) [Leipzig Office](https://www.linkedin.com/feed/update/urn:li:activity:7193158068761706496/) at Altes Landratsamt. We had two
talks that revolved around Go and performance: [Patrick Vahl](https://github.com/Crash129) gave a great
overview about SIMD in Go and what is possible (and not possible), today. Examples from the talk
can be found here:
[Crash129/go-simd-example](https://github.com/Crash129/go-simd-example). A few points from the talk:

* performance gains can be significant
* try to isolate SIMD as much as possible

> The difficulty with a general purpose approach to SIMD [...] is that the performance can be dramatically different on
> different processors. -- [go/issues/53171](https://github.com/golang/go/issues/53171#issuecomment-1142953120)

* Go uses its own assembly syntax as an intermediate representation, more on that in [The Design of the Go Assembler](https://go.dev/talks/2016/asm.slide#1) (2016)
* there are various approaches to SIMD and Go including [avo](https://github.com/mmcloughlin/avo), [gocc](https://github.com/kelindar/gocc), and more - see: [Examples of methods to use SIMD in Go](https://github.com/Crash129/go-simd-example?tab=readme-ov-file#examples-of-methods-to-use-simd-in-go)
* explore the assembly generated for Go programs via [compiler explorer](https://godbolt.org/z/oKvWhxqqn)

While SIMD support is still a topic of discussion, e.g. in [#53171](https://github.com/golang/go/issues/53171), there exist usable approaches today.

[Martin Czygan](https://de.linkedin.com/in/martin-czygan-58348842) took a Go
implementation of the [1BRC](https://web.archive.org/web/20241116164008/https://1brc.dev/) from 4 minutes to 4s, using
fewer allocations, parallel processing and mmap.

![](/images/657582.gif)

Notes: [1brc-in-go](https://github.com/golang-leipzig/1brc-in-go). Please
submit a PR with a variant you found!

Thanks a lot to [CHECK24](https://check24.de) for hosting the event in their
nice office (with a view). Looking forward to our next event together which is
planned for October 2024.



## Misc

* A good use case for [mmap](https://man7.org/linux/man-pages/man2/mmap.2.html) is a parallel processing of read only data (that can be bigger than the physical RAM);
  there is an experimental package in the
[x/exp/mmap](https://pkg.go.dev/golang.org/x/exp/mmap) and a library [edsrzf/mmap-go](https://github.com/edsrzf/mmap-go). The concept of virtual
memory goes back to one of the first super computer
[Atlas](https://en.wikipedia.org/wiki/Atlas_(computer)).
* The [simdjson](https://github.com/simdjson/simdjson) project explores the use of SIMD for record breaking, high performance JSON processing

<!--
https://www.linkedin.com/feed/update/urn:li:activity:7193158068761706496?updateEntityUrn=urn%3Ali%3Afs_feedUpdate%3A%28V2%2Curn%3Ali%3Aactivity%3A7193158068761706496%29
-->
