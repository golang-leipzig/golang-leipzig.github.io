---
title: "Second meetup wrap-up"
date: 2019-03-16T13:15:00+01:00
draft: false
tags:
- summary
- meetup
---

Our [second
meetup](https://www.meetup.com/Leipzig-Golang-and-Cloud/events/259045114/) took
place yesterday (2019-03-15) at [Basislager Leipzig](https://www.basislager.co), here's a
short wrap-up:

1. [miku](https://github.com/miku) gave a lightning talk about generating
   Go structs from JSON and XML documents, e.g. via
[JSONGen](https://github.com/bemasher/JSONGen) or
[zek](https://github.com/miku/zek). Code generation is not perfect, but can
save you time.

2. [Michael](https://www.meetup.com/Leipzig-Golang-and-Cloud/members/201296430)
   presented the paper [Understanding Real-World Concurrency Bugs in
Go](https://songlh.github.io/paper/go-study.pdf), which compares bug classes by
analyzing popular open source projects like
[BoltDB](https://github.com/etcd-io/bbolt),
[etcd](https://github.com/etcd-io/etcd), [Docker](https://github.com/docker)
and [Kubernetes](https://kubernetes.io/). Did you know that some concurrency
bugs stay undetected for months, [even
years](https://speakerdeck.com/embano1/concurrency-bugs-in-go-go-meetup-leipzig-03-15-2019-gasch?slide=10)?
While message passing does not make multithreaded programs less-error prone, it
can still offer a clean form of inter-thread communication.

3. Staying on the topic of concurrency, we reviewed a classic, concurrent
   producer-consumer problem, translated into Go by
[Michael](https://github.com/panzerdev). In less than [150 lines of
code](https://gist.github.com/panzerdev/cf8278b346770814088c006643f9bfd1) is it
possible to write an auto-scaling worker-pool with retries that is using Redis
as backing store. An alternative to spawning goroutines for every task is to
feed a fixed number of workers through a channel (where work is distributed
round-robin to workers). Both designs work, but we were unsure about some
performance implications - we'll briefly report on new findings at the next
meetup.

Apart from these topics, Go and cloud native topics offer plenty of room for
discussions: the field is moving fast, while a popular implementation language
for [CN projects](https://landscape.cncf.io/) - Go - [stays boring](https://golang.org/doc/go1.12):

> **Changes to the language**
>
> There are no changes to the language specification.

If you have an idea for a lightning talk or presentation, or if you already use
Go or cloud native technologies, we love to hear real-world insights, just ping us on
[meetup](https://www.meetup.com/Leipzig-Golang-and-Cloud/),
[twitter](https://twitter.com/golang_leipzig),
[github](https://github.com/golang-leipzig) or
[e-mail](mailto:martin.czygan@gmail.com).

Thanks for dropping by!

## References

### Code generation

* JSON, [JSONGen](https://github.com/bemasher/JSONGen), [online version, json-to-go](https://mholt.github.io/json-to-go/)
* XML, [zek](https://github.com/miku/zek), [online version, xmltogo](https://www.onlinetool.io/xmltogo/)

### Concurrency bugs

* [Lightning Talk Slides](https://speakerdeck.com/embano1/concurrency-bugs-in-go-go-meetup-leipzig-03-15-2019-gasch)
* [Understanding Real-World Concurrency Bugs in Go](https://songlh.github.io/paper/go-study.pdf) (2019)
* Book recommendation: [Concurrency in Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/) (2017)

### Worker pools

* [Topic based Worker with Redis example](https://gist.github.com/panzerdev/cf8278b346770814088c006643f9bfd1)

### Misc

* [Go package management](https://github.com/golang/go/wiki/PackageManagementTools) - we plan to have a Go modules deep dive in April 2019 meetup.
* Meeting forms: [lean coffee](http://leancoffee.org/).
* Video: [GOTO 2018 • Old Is the New New • Kevlin Henney](https://www.youtube.com/watch?v=AbgsfeGvg3E) (2018) - Fundamentals.
* Channel: [computerphile](https://www.youtube.com/channel/UC9-y-6csu5WGm29I7JiwpnA) - Videos all about computers and computer stuff. Sister channel of Numberphile.

Now defunkt company [Basho](https://en.wikipedia.org/wiki/Basho_Technologies),
creators of database technology (RIAK) had a format, similar to [papers we
love](https://paperswelove.org/), but their video recordings seem to be lost.
If you have happen to have a copy, please consider contacting us :)

* A classic tool: [sed](https://www.gnu.org/software/sed/manual/sed.html) -- a stream editor (1974).
* Book recommendation: [Designing Data-Intensive
  Applications](https://dataintensive.net/) (2017);
[avro](https://en.wikipedia.org/wiki/Apache_Avro); [Pat
Helland](https://dblp.uni-trier.de/pers/hd/h/Helland:Pat), author of [building
on quicksand](http://db.cs.berkeley.edu/cs286/papers/quicksand-cidr2009.pdf) (2009).

