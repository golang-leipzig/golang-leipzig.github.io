---
title: "Virtual Meetup #26 wrap-up"
date: 2022-04-27T01:00:00+02:00
draft: false
tags:
- summary
- meetup
---

## Type Signatures and Cloud Storage

[Meetup #26](https://www.meetup.com/Leipzig-Golang/events/282941906/) was
virtual again and we were glad to meet Gophers from Leipzig and beyond. We had a two short talks and code walkthroughs on type signatures and rclone, a
cloud storage tool.

### Type Signatures

How would you generate a unique fingerprint for a type? A fingerprint that would change, even
if the name of the type stays the same, but e.g. fields in a struct are renamed, added or removed?

The term [Type Signature](https://en.wikipedia.org/wiki/Type_signature) is
often used in the context of functions and methods and contains name and
parameters of a function.

In Java, we have [serialVersionUID](https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/io/Serializable.html):

> If a serializable class does not explicitly declare a `serialVersionUID`, then
> the serialization runtime will calculate a default `serialVersionUID` value for
> that class based on various aspects of the class, as described in the
> Java(TM) Object Serialization Specification. However, it is strongly
> recommended that all serializable classes explicitly declare `serialVersionUID`
> values, since the default `serialVersionUID` computation is highly sensitive to
> class details that may vary depending on compiler implementations, and can
> thus result in unexpected InvalidClassExceptions during deserialization.

In case of struct types, one way would be to serialize the zero value of
the type and compute some hash, e.g.
[FNV](https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function)
or [other](https://pkg.go.dev/hash#section-directories).

### Cloud storage tool: rclone

[Rclone](https://rclone.org/) is a

> command-line program to manage files on cloud storage. It is a feature-rich
> alternative to cloud vendors' web storage interfaces. Over 40 cloud storage
> products support rclone including S3 object stores, business & consumer file
> storage services, as well as standard transfer protocols.

Beside file transfers rclone allows to mount filesystems, analyze disk usage
(similar to [ncdu](https://dev.yorhel.nl/ncdu)) and expose cloud storage
providers with classic protocols like FTP, webdav or http.

Find out more: [The rsync for the cloud era: Rclone](https://github.com/miku/rclone-lightning-talk).


## Misc

* rclone does not seem to support [upspin](https://upspin.io/) (nor [ipfs](https://github.com/rclone/rclone/issues/128), for that matter)
* in 2022 go is still niche!
  [152](https://web.archive.org/web/20220427105818/https://www.stepstone.de/5/ergebnisliste.html?what=golang&searchOrigin=Homepage_top-search)
vs
[3420](https://web.archive.org/web/20220427105843/https://www.stepstone.de/5/ergebnisliste.html?what=java&searchOrigin=Homepage_top-search)
jobs for "golang" vs "java" (are you looking for a "golang" job? check out: [aboutyou.de/jobs](https://corporate.aboutyou.de/en/our-jobs))
* for more unusual uses of [reflect](https://pkg.go.dev/reflect) ([laws](https://go.dev/blog/laws-of-reflection)) take a look at: [https://github.com/kstenerud/go-subvert](https://github.com/kstenerud/go-subvert)
* tricky things: [Go interfaces, the tricky parts](https://www.timr.co/go-interfaces-the-tricky-parts/), [Go Gotchas](https://github.com/golang-leipzig/gotchas)
* new things: [Go 1.18 with some fuzzing and a focus on generics](https://www.klingt.net/articles/go-1-18-with-some-fuzzing-and-a-focus-on-generics.html)
* large changes in Go 1.18 like generics and fuzzying are seemingly not used by all projects immediately (example: [mockery](https://github.com/vektra/mockery))
* postgres go driver: [lib/pq](https://github.com/lib/pq), [jackc/pgx/v4](https://github.com/jackc/pgx/tree/v4); a new kind of ORM for Go: [bun](https://bun.uptrace.dev)
* for quick online tools: [CyberChef](https://gchq.github.io/CyberChef/)
* more interactive online things: [https://messwithdns.net](https://messwithdns.net)
* sad, but true as ever: [https://xkcd.com/2347](https://xkcd.com/2347) (discussing [jq](https://stedolan.github.io/jq/))
* tiny demo: we had a short demo of a Go program called *runpad*, collaborative code editing - related to [mob programming](https://mob.sh)

![](/images/runpad-screenie.png)

----

[Join our meetup](https://www.meetup.com/Leipzig-Golang) to get notified of
upcoming events!

