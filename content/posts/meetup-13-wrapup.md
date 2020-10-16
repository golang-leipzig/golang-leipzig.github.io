---
title: "Virtual Meetup #13 wrap-up"
date: 2020-10-16T22:00:00+02:00
draft: false
tags:
- summary
- meetup
---

Today we had two input presentations:

* Haystack and seaweedfs: [https://github.com/miku/haystack](https://github.com/miku/haystack)
* Functional Options Pattern: [https://github.com/embano1/funcy-ops](https://github.com/embano1/funcy-ops)

## Haystack

The haystack talk gave an overview about an early Facebook photo storage system
and a Go project inspired by it, called
[seaweedfs](https://github.com/chrislusf/seaweedfs). The seaweedfs project is
great, as it is developer friend, scales up and down and is easy to use out of
the box. It offers an [S3 compatible
API](https://github.com/chrislusf/seaweedfs/wiki/Amazon-S3-API) and a [FUSE
filesystem](https://github.com/chrislusf/seaweedfs/wiki/FUSE-Mount) option.

![](/images/assembly_pic13.png)

## Functional Option Pattern

The function option pattern addresses API stability when it comes to
configuration options. A simple way to configure a type is by passing in
options directly (there can be too many), or by passing in a single config
struct type (which becomes a part of the public API). The functional option
pattern works because in Go function as just like any other value and can be
passed as arguments or returned by functions.

```go
// WithInsecure skips certificate verification
func WithInsecure(in bool) ServerOption {
    return func(s *Server) error {
        if in {
            return errors.New("security alert")
        }
        s.insecure = in
        return nil
    }
}
```

The complete example can be found here:
[https://github.com/embano1/funcy-ops/blob/main/server.go](https://github.com/embano1/funcy-ops/blob/main/server.go).

The functional options is a Go pattern and addresses some issues with too many
arguments, variants of New and API stability.

## Misc

Discussion evolved around testing (Go) code with controlled, restricted
environments, like memory limits or maybe different
[ulimits](https://man7.org/linux/man-pages/man3/ulimit.3.html). Running tests
in containers can address these tasks (or one can get
[systemd](https://wiki.archlinux.org/index.php/Cgroups#With_systemd_2) to
impose limits).

Furthermore, another way to embed files:

* [https://github.com/klingtnet/embed](https://github.com/klingtnet/embed)

Book recommendatation:

* [databass.dev](https://www.databass.dev/)

Go runtime interplay with containers and pods? Fasten your seatbelt:
[https://twitter.com/embano1/status/1149654812595646471](https://twitter.com/embano1/status/1149654812595646471)


And finally, a few ideas that came up as topics for another talk:

* crypto?
* kubernetes operators (and what they can teach about (event oriented) systems design in general)
* Higher Order Go, or functional patterns in Go

Join us [next time](https://www.meetup.com/Leipzig-Golang/events/268785591/)!

----

Image credit: [Kobol](https://wiki.kobol.io/helios64/intro/)
