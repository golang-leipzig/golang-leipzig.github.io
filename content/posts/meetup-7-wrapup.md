---
title: "Meetup #7 wrap-up"
date: 2019-09-28T14:00:00+02:00
draft: false
tags:
- summary
- meetup
---

{{% h2 %}}Concurrency in Go{{% /h2 %}}

![](/images/cignotes-chapter-1-topics.png)

Concurrency is hard, but the primitive building blocks of CSP in Go -
goroutines and channels - can make things simpler.

We glanced at the six chapters of [Concurrency in
Go](http://shop.oreilly.com/product/0636920046189.do), and learned a bit about
the problem setting, goroutine sizes, context switch times and patterns.

The notes can be found here:

* [Notes on Concurrency in Go](https://github.com/miku/cignotes)

In the book, the [perf](https://perf.wiki.kernel.org) tool is used to benchmark
context switch times of the kernel - and to compare it with context switch
times in the Go scheduler.

Sidenote: you can see the number of context switches with various tools, e.g. [dstat](https://linux.die.net/man/1/dstat).

![](/images/dstat.png)

Another performance debugging tool are flame graphs, which can be extended in
time, e.g. with flamescope:

* [flamescope](https://github.com/Netflix/flamescope)

Apart from technical topics, we discussed adoption stories of Go in companies.
It's impossible to replace a huge system at once, but people might be surprised
by how resource efficient Go can be - and get curious.


