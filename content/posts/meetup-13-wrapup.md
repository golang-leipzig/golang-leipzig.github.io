---
title: "Virtual Meetup #13 wrap-up"
date: 2020-10-16T22:00:00+02:00
draft: false
tags:
- summary
- meetup
---

Today we had two input presentations:

* Haystack and seaweedfs: [miku/haystack](https://github.com/miku/haystack)
* Functional Options Pattern: [embano1/funcy-ops](https://github.com/embano1/funcy-ops)

## Haystack

The haystack talk gave an overview about an early Facebook photo storage system
and a Go project inspired by it, called
[seaweedfs](https://github.com/chrislusf/seaweedfs). The seaweedfs project is
great, as it is developer friendly, scales up and down and is easy to use out of
the box. It offers an [S3 compatible
API](https://github.com/chrislusf/seaweedfs/wiki/Amazon-S3-API) and a [FUSE
filesystem](https://github.com/chrislusf/seaweedfs/wiki/FUSE-Mount) option.

![](/images/assembly_pic13.png)

## Functional Options Pattern

The *Functional Options Pattern* addresses API stability when it comes to
configuration options. A simple way to configure a type is by passing in
options directly (there can be too many), or by passing in a single config
struct (which becomes part of the public API, for better or for worse). The functional options
pattern works because in Go functions are first class values and can be passed
as arguments into or returned by functions.

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

Functional Options is a Go *pattern* and addresses some issues with too many
arguments, variants of *New* and API stability.

## Misc

Discussion evolved around testing (Go) code with controlled, restricted
environments, like memory limits or maybe different
[ulimits](https://man7.org/linux/man-pages/man3/ulimit.3.html). Running tests
in containers can address these tasks (or one can get
[systemd](https://wiki.archlinux.org/index.php/Cgroups#With_systemd_2) to
impose limits).

Furthermore, another way to embed files (there is a proposal from [July
2020](https://go.googlesource.com/proposal/+/master/design/draft-embed.md)):

* [https://github.com/klingtnet/embed](https://github.com/klingtnet/embed)

Book recommendation:

* [databass.dev](https://www.databass.dev/)

Go runtime interplay with containers and pods? Fasten your seatbelt:
[https://twitter.com/embano1/status/1149654812595646471](https://twitter.com/embano1/status/1149654812595646471)

And finally, a few ideas that came up as topics for another talk:

* (Go) Crypto
* Kubernetes operators (and what they can teach about (event oriented) systems design in general)
* Higher Order Go, or functional patterns in Go

Join us [next time](https://www.meetup.com/Leipzig-Golang/events/268785591/)!

----

Image credit: [Kobol](https://wiki.kobol.io/helios64/intro/)
