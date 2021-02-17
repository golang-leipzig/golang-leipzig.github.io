---
title: "Virtual Meetup #16 wrap-up"
date: 2021-02-17T13:00:00+01:00
draft: false
tags:
- summary
- meetup
---

## Go 1.16

Meetup #16 was all about Go 1.16 (which [was released](https://blog.golang.org/go1.16) after the meetup).

We had an informative breakdown of the updates coming in this release, the full slides can be found here:

* [A quick tour of Go 1.16](https://github.com/golang-leipzig/go-1.16-and-beyond)

A few highlights are: the go tool is now [module
aware](https://golang.org/doc/go1.16#go-command) by default,
[ioutil](https://golang.org/pkg/io/ioutil/) gets deprecated, support for file
embedding (we talked about that in
[#14](https://golangleipzig.space/posts/meetup-14-wrapup/)) and there is a new
io subpackage [io/fs](https://golang.org/pkg/io/fs/), defining among other
things interfaces for filesystems.

We also saw examples

Check out the [slides](https://github.com/golang-leipzig/go-1.16-and-beyond),
[blog post](https://blog.golang.org/go1.16) and [release
notes](https://golang.org/doc/go1.16) for more details.

Thanks [@klingtnet](https://github.com/klingtnet/) for the great summary!

## Generics Pro/Con

Go is simple, and generics will add complexity, while helping to reduce
repetition (and lines of code). In the best case, this results in less, yet
still clear code.

As programming languages tend to *add* features, the Go release note
section on [Changes to the language](https://golang.org/doc/go1.16#language)
happily reported *There are no changes to the language* - over the years.

Less can be more, especially in complex technological landscapes and dynamic
environments. Being a *reduced language* was a bit special, but many
experienced it also as an advantage, or at least not as a problem for the types
of problems people use Go for.

Anyway, generics have [been accepted](https://golang.org/doc/faq#generics) - and
may be available in Go 1.18 (2022).

## What do you miss from Go?

Go is simple, but what are features people miss? We found the following:

* [ternary if](https://en.wikipedia.org/wiki/%3F:) (e.g. C, Java, ...)
* pattern matching (Rust, Erlang, ...)
* enum (C, Rust, Python, ...)
* union types (C, PHP8, ...)
* a built-in set datatype (Python, ...)
* functional utilities, like list comprehensions (Python, ...)
* lambda expressions (Java, Python, ...)

## Misc

* how about [boring technology](http://boringtechnology.club/)
* shout-out to [google/ko](https://github.com/google/ko)
* [distroless](https://github.com/GoogleContainerTools/distroless) are stripped
  down container images (but more than
[scratch](https://github.com/containers/buildah/blob/master/docs/tutorials/01-intro.md#building-a-container-from-scratch)
[images](https://hub.docker.com/_/scratch))

We quizzed ourselves on what kind of interview question might be suitable for
intermediate Go programmers? Some idea we came up with:

* explain [slices](https://blog.golang.org/slices-intro)
* discuss [concurrency](https://tour.golang.org/concurrency/1) and [context](https://blog.golang.org/context)
* a walkthrough project, assessing familiarity with tooling and practices around building projects
* a joint merge request review

----

[Join our meetup](https://www.meetup.com/Leipzig-Golang) to get notified of upcoming events!

