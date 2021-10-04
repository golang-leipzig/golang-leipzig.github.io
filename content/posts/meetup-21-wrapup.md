---
title: "Virtual Meetup #21 wrap-up"
date: 2021-10-03T14:00:00+02:00
draft: false
tags:
- summary
- meetup
---

## Concise Encoding

We had a great presentation about [Concise
Encoding](https://concise-encoding.org/), a proposed encoding standard that aims to
be human and machine friendly.

* [Presentation: go-concise-encoding](https://github.com/kstenerud/go-concise-encoding/blob/presentation/presentation-go-concise-encoding.md)

Here's a recording:

[![Leipzig Gophers #21: Concise Encoding](https://img.youtube.com/vi/_dIHq4GJE14/0.jpg)](https://www.youtube.com/watch?v=_dIHq4GJE14)

Concise encoding is element-oriented (similar to e.g. JSON) and tries to solve
some issues in existing standards, such as lack of types, performance and
security (one similar format is [Amazon Ion](https://amzn.github.io/ion-docs/)).

In prerelease stage, [feedback is highly
appreciated](https://github.com/kstenerud/concise-encoding), a reference
implementation written in Go exists at:

* [https://github.com/kstenerud/go-concise-encoding](https://github.com/kstenerud/go-concise-encoding)

We looked at the
[architecture](https://github.com/kstenerud/go-concise-encoding/blob/presentation/presentation-go-concise-encoding.md#architecture)
and example iterators and builders and
[takeaways](https://github.com/kstenerud/go-concise-encoding/blob/presentation/presentation-go-concise-encoding.md#takeaways).

Usage examples:

*  [https://play.golang.org/p/6_pD6CQVLuN](https://play.golang.org/p/6_pD6CQVLuN)

## Misc

Go has great tooling, especially for performance analysis: The
[pprof](https://pkg.go.dev/runtime/pprof) suite allows to understand runtime
and memory issues of your program.

Additional helpers exists, such as
[benchcmp](https://pkg.go.dev/golang.org/x/tools/cmd/benchcmp) or the
possibility to inspect changes over time, with the `-diff` flag:

```sh
$ go tool pprof -web -diff_base cpu-prev.prof cpu-now.prof
```

Allocations can kill performance and sometimes they are not obvious, e.g.
converting a byte slice to a string requires an allocation. However, you can
manipulate a string header directly - but this is
[unsafe](https://pkg.go.dev/unsafe):

* [https://play.golang.org/p/t9KGXPmoMkA](https://play.golang.org/p/t9KGXPmoMkA)

A few more interesting projects:

* [https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/fieldalignment](https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/fieldalignment)
* [https://github.com/dominikh/go-tools/tree/master/cmd/structlayout](https://github.com/dominikh/go-tools/tree/master/cmd/structlayout)
* [https://github.com/segmentio/encoding/tree/master/json](https://github.com/segmentio/encoding/tree/master/json)

And a blog on contemporary computing architectures and performance:
[https://lemire.me/blog/](https://lemire.me/blog/).

Also, if you ever searched for

> a gentle, opinionated, hands-on introduction to NixOS

please check out:

* [https://github.com/kstenerud/nixos-beginners-handbook](https://github.com/kstenerud/nixos-beginners-handbook)

Finally, a short free course to improve your technical writing:

* [https://developers.google.com/tech-writing](https://developers.google.com/tech-writing)


Thanks [Karl](https://github.com/kstenerud) for the presentation and everyone for dropping by.

----

[Join our meetup](https://www.meetup.com/Leipzig-Golang) to get notified of
upcoming events!

