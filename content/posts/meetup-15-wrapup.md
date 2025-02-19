---
title: "Virtual Meetup #15 wrap-up"
date: 2020-12-19T14:00:00+02:00
draft: false
tags:
- summary
- meetup
---

## Concurrent woes

We had two input presentations, both involving concurrent constructs.

First, we heard from the [author](https://twitter.com/embano1) of [package
waitgroup](https://github.com/embano1/waitgroup) about problems with concurrent
code, waitgroups and shutdowns (especially in cloud native and serverless
scenarios). The basic problem: At shutdown time, you want to be graceful, but
at the same time enforce a hard timeout on a number of processes running (and
potentially joining in a [`wg.Wait`](https://golang.org/pkg/sync/#WaitGroup.Wait)).

Enter [package waitgroup](https://github.com/embano1/waitgroup) which allows to
preempt running goroutines. The general problem of goroutine termination is
hard, and even harder, if you have a large code base and a tree of concurrent
routines in flight, that need some kind of cancellation (modern packages use
[context](https://golang.org/pkg/context/), or maybe manual
[timeouts](https://github.com/golang/go/wiki/Timeouts)).

The solution of [waitgroup](https://github.com/embano1/waitgroup) is to embed
the [standard library waitgroup](https://golang.org/pkg/sync/#WaitGroup) and
combine it with a [timeout](https://github.com/golang/go/wiki/Timeouts). The
package comes with examples too:

* [https://pkg.go.dev/github.com/embano1/waitgroup#example-Await](https://pkg.go.dev/github.com/embano1/waitgroup#example-Await)

Testing concurrent code is its own challange, but this package has examples for that as well.

----

The second lightning talk summarized chapter 2 on "Filters" from [Software
Tools](https://openlibrary.org/works/OL4617639W/) by Kernighan and Plauger
(1976) and highlighted some timeless ideas on software style (today, you might
say clean code). Go has great facilities for streaming data processing in the
[io](https://golang.org/pkg/io/) package. Furthermore, you can spice up filters
with parallel processing, without adding much complexity to the program.

![](/images/softwaretools76-s.png)

That's what [parallel](https://github.com/miku/parallel) is about - it allow to
write parallel filters fast. The package is best used as a little utility for
one off filters.

* [Fast parallel filters in Go](https://gist.github.com/miku/738f361c8156264626c74f9b717927ff#fast-parallel-filters-in-go)

## Misc

The best about meetups are things beyond slides and code: discussions, screen
shared setups, neat tools people use, and much more.

Here's a list of things to check out:

* The [ko](https://github.com/google/ko) project helps to get Go applications
  to run on [kubernetes](https://kubernetes.io/) fast, with minimal effort. If
you run go on k8s, you will need this.
* The [zap](https://github.com/uber-go/zap) logger is a fast structured logging library.
* Just as you can combine waitgroups with timeouts, you can blend context
  and readers, for details see:
[https://pace.dev/blog/2020/02/03/context-aware-ioreader-for-golang-by-mat-ryer.html](https://pace.dev/blog/2020/02/03/context-aware-ioreader-for-golang-by-mat-ryer.html)
* Behaviour driven development (BDD) uses a semi-formal language to describe test cases, in Go we have [ginkgo](https://github.com/onsi/ginkgo) - the terminology is language-agnostic.

Can you tame complexity with BDD style development? A VMware *Office of the CTO* tech deep dive addresses this question in a blog post:

* Taming Complexity in Software Development with Behavior Driven Development (BDD) (vanished as of 05/2024)

As for stream processing, there is a classic blog post, as well as a few interesting projects:

* [Go concurrency patterns: Pipelines and cancellation](https://blog.golang.org/pipelines)
* [Automi](https://github.com/vladimirvivien/automi) (prototype)
* [RxGo](https://github.com/ReactiveX/RxGo), an implementation of
  [reactive](http://reactivex.io) style (in short: observer pattern + iterator
pattern + functional programming)

Fast things:

* [GNU parallel](https://www.gnu.org/software/parallel/)
* [livegrep](https://github.com/livegrep/livegrep)
* [fzf](https://github.com/junegunn/fzf)

Basic (database) things:

* BW trees, [What is a BW-tree?](https://stackoverflow.com/questions/18859123/what-is-a-bw-tree), [Paper](https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/bw-tree-icde2013-final.pdf)
* log-structured merge trees
* and many more (see: [databass.dev/](https://databass.dev/))

----

[Join our meetup](https://www.meetup.com/Leipzig-Golang) to get notified of upcoming events!

