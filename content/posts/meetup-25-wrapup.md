---
title: "Virtual Meetup #25 wrap-up"
date: 2022-03-06T01:00:00+01:00
draft: false
tags:
- summary
- meetup
---

## Go 1.18 and generics

[Meetup #25](https://www.meetup.com/Leipzig-Golang/events/282941887/) was
virtual and we were glad to have an international audience again.

We had a great presentation about the upcoming [Go 1.18](https://tip.golang.org/doc/go1.18) release, including generics and fuzzing:

* [Go 1.18 with some fuzzing and a focus on generics](https://www.klingt.net/articles/go-1-18-with-some-fuzzing-and-a-focus-on-generics.html)

As always, Go 1.18 will include lots of smaller tweaks and improvements, like a
new [debug/buildinfo](https://pkg.go.dev/debug/buildinfo) package, which

> provides access to information embedded in a Go binary about how it was built,

or [MaxBytesHandler](https://pkg.go.dev/net/http@master#MaxBytesHandler), a
middleware to protect e.g. against denial-of-service type attacks.

You can [read the blog
post](https://www.klingt.net/articles/go-1-18-with-some-fuzzing-and-a-focus-on-generics.html)
to take a deeper dive.

## Misc

* the influencial [How to become a Hacker](http://www.catb.org/~esr/faqs/hacker-howto.html) rev 1.52 notes (2020-01-03):

> Go makes a place as a plausible learning language, displacing Java. [...]

And further:

> A better alternative to Java is to learn Go. This relatively new language is
> pretty easy to move to from Python, and learning it give you a serious leg up
> on the possible next step, which is learning C. Additionally, one of the
> unknowns about the next few years is to what extent Go might actually
> displace C as a systems-programming language. There is a possible future in
> which that happens over much of C's traditional range.

Other programmers reflecting and discussing this topic:
[HN23377186](https://news.ycombinator.com/item?id=23377186),
[LO](https://lobste.rs/search?q=%22how+to+become+a+hacker%22&what=stories&order=newest), ...

* Go 1.18 adds a new environment variable, `GOAMD64`, which splits support
  for 64-bit x86 processors into [4 different levels](https://github.com/golang/go/wiki/MinimumRequirements#amd64), currently - unlocking
various optimizations. Pop quiz: Do you know, which instruction has also been
named [The NSA instruction](https://vaibhavsagar.com/blog/2019/09/08/popcount/)?
Hint: it belongs to "v2" and has to do with the [Hamming weight](https://en.wikipedia.org/wiki/Hamming_weight) of a string
* [sync.Pool](https://pkg.go.dev/sync#Pool) helps to reduce GC pressure, by allowing reuse of allocated memory, akin to a [free list](https://en.wikipedia.org/wiki/Free_list)
* we gave away a [Jetbrains](https://www.jetbrains.com/) Personal Subscription and a copy of [100 mistakes in
  Go and how to avoid
them](https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them) --
thank you, sponsors!

----

🇺🇦 &mdash; and one more thing: нет войны! Stop the
[War](https://en.wikipedia.org/wiki/2022_Russian_invasion_of_Ukraine)! -- to help locally: [Leipzig helps Ukraine](https://leipzig-helps-ukraine.de).


----

[Join our meetup](https://www.meetup.com/Leipzig-Golang) to get notified of
upcoming events!

