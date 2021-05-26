---
title: "Virtual Meetup #18 wrap-up"
date: 2021-05-26T12:00:00+02:00
draft: false
tags:
- summary
- meetup
---

## Seeking Data

Meetup #18 took place May 25, 2021 19:00 CEST, and was virtual again. We
continued the data and storage supertopic for this year had a lightning talk
called:

* [Seeking Data](https://gist.github.com/miku/6dcffb5c104bc44709c330ec90682189)

![](/images/floppy_icon.gif)

We looked at a few patterns when working with data, e.g. I/O considerations,
what kernel hackers [think](http://varnish-cache.org/docs/trunk/phk/notes.html)
about the memory hierarchy, working with large scale data at rest, or how to
speed up container startup times (via [lazy
pulling](https://github.com/opencontainers/image-spec/issues/815)).


## Misc

* We learned about [Concise Encoding](https://concise-encoding.org/) and hope to
do a deep dive on this topic in a future meetup.
* Technologies, which are widespread, despite some flaws, are hard to replace,
  e.g. JSON might have issues, yet it is widely used and has lots of tooling
around it (might be an instance of [Worse is
Better](https://en.wikipedia.org/wiki/Worse_is_better)).
* Fuzzying is an interesting topic, and many tools are written in Go, e.g.
  [ffuf](https://github.com/ffuf/ffuf), a fast web fuzzer written in Go.
* For autoincrement URL snooping, you can use [urlbisect](https://github.com/miku/urlbisect).
* Go is not a classic functional programming language, but it supports function
  types ([Walkthrough: First-Class
Functions](https://golang.org/doc/codewalk/functions/), [Function
types](https://golang.org/ref/spec#Function_types)), which will allow for some
*higher order programming*.
* How to structure Go projects? It will depend. There is not one *standard* layout. Start small and then extend as
  needed. Some people enjoy putting old-school
[Makefiles](https://www.gnu.org/software/make/manual/make.html#Overview) in their project root as well.
* Testing strategies in Go: classic [Table-Driven
  Tests](https://github.com/golang/go/wiki/TableDrivenTests), [Test
Containers](https://github.com/testcontainers/testcontainers-go) (example:
testing [esbulk](https://github.com/miku/esbulk) against [a few major
versions](https://git.io/JGUUS) of elasticsearch).
* [ORM](https://en.wikipedia.org/wiki/Object%E2%80%93relational_mapping) seems
  to have fallen out of favor a bit; switching databases is not that common
after all, so why not start with some lighter helpers like
[sqlx](https://github.com/jmoiron/sqlx), or specific database libraries.

## Upcoming Go 1-day workshop

There will be an upcoming, online, 1-day workshop on Go at [Spartakiade
2021](https://spartakiade.org/), a developer-centric community conference (not
free, but with nice variety of topics plus a bag of conference swag). The [Go
workshop](https://github.com/miku/goforprogrammers) will be a condensed
overview of the language - feel free to join!

## Thanks

Thanks everyone for dropping by - great to see people join from across Europe and the globe!

----

[Join our meetup](https://www.meetup.com/Leipzig-Golang) to get notified of upcoming events!

