---
title: "Virtual Meetup #15 wrap-up"
date: 2020-12-19T14:00:00+02:00
draft: false
tags:
- summary
- meetup
---

## Concurrent woes

We had two input presentations, both involving concurrent contructs.

First, we heard from the [author](https://twitter.com/embano1) of [package
waitgroup](https://github.com/embano1/waitgroup) about problems with concurrent
code, waitgroups and shutdowns. The basic problem: At shutdown time, you want
to be graceful, but at the same time enforce a hard timeout on a number of
processes running (and potentially joining in a `wg.Wait`).

Enter [package waitgroup](https://github.com/embano1/waitgroup) which allows to
preempt running goroutines. The general problem of goroutine termination is
hard, and even harder, if you have a large code base and a tree of concurrent
routines in flight, that need some kind of cancellation (modern packages use
[context](https://golang.org/pkg/context/), or maybe manual
[timeouts](https://github.com/golang/go/wiki/Timeouts)).

The solution of [waitgroup](https://github.com/embano1/waitgroup) is to embed
the standard library waitgroup and combine it with timeout. The package comes
with, examples too:

* [https://pkg.go.dev/github.com/embano1/waitgroup#example-Await](https://pkg.go.dev/github.com/embano1/waitgroup#example-Await)

The second lightning talk summarized chapter 2 on "Filters" of [Software
Tools](https://openlibrary.org/works/OL4617639W/) by Kernighan and Plauger
(1976) and highlighted some timeless ideas on software style (today, you might
say clean code). Go has great facilities for streaming data processing in the
[io](https://golang.org/pkg/io/) package. Furthermore, you can spice up filters
with parallel processing, without adding much complexity to the program.

That's what [parallel](https://github.com/miku/parallel) is about - it allow to
write parallel filters fast. The package is best used as a little utility for
one off filters.

* [Fast parallel filters in Go](https://gist.github.com/miku/738f361c8156264626c74f9b717927ff#fast-parallel-filters-in-go)

## Misc

----

[Join our meetup](https://www.meetup.com/Leipzig-Golang) to get notified of upcoming events!

