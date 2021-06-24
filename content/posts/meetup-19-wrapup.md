---
title: "Virtual Meetup #19 wrap-up"
date: 2021-06-23T12:00:00+02:00
draft: false
tags:
- summary
- meetup
---

## Journey to the Center of Kubernetes

Meetup #19 took place [Jun 22, 2021 19:00
CEST](https://www.meetup.com/Leipzig-Golang/events/275871279/), and was virtual
again. [Michael Gasch](https://www.mgasch.com) gave an input presentation on [etcd](https://etcd.io/) and did a deep dive into its
role in the container-orchestrator [kubernetes](https://kubernetes.io/).

The talk is based on a blog post: [Onwards to the Core: etcd](https://www.mgasch.com/2021/01/listwatch-part-1/). A shell transcript can be found here:

* [https://git.io/JnQ1n](https://gist.github.com/embano1/aedd423abe07c8012897658db15e139a)

We learned about the API of etcd, its data model, key value stores layouts,
append-only storage and compaction. A recording of the talk is [available on
YouTube](https://www.youtube.com/watch?v=Z9fwIzy0C_8):

[![Leipzig Gophers #19: A Journey in the Kubernetes ListerWatcher Rabbit Hole](http://img.youtube.com/vi/Z9fwIzy0C_8/0.jpg)](https://www.youtube.com/watch?v=Z9fwIzy0C_8)

A few highlights from the talk:

* etcd is a distributed key value store, using the [RAFT](https://raft.github.io/) consensus algorithm.

> Consensus algorithms allow a collection of machinesto work as a coherent
> group that can survive the fail-ures of some of its members. Because of this,
> they play akey role in building reliable large-scale software systems. -- [In Search of an Understandable Consensus Algorithm](https://raft.github.io/raft.pdf)

* etcd internally uses [boltdb](https://github.com/etcd-io/bbolt), especially
  their own fork, as the original project has been
[archived](https://github.com/boltdb/bolt) -- boltdb itself is a Go
implementation of the
[LMDB](https://en.wikipedia.org/wiki/Lightning_Memory-Mapped_Database) design,
which among many other things supports
[MVCC](https://en.wikipedia.org/wiki/Multiversion_concurrency_control)
* etcd had hierarchical, but moved to a flat key value design; albeit you can
  mimick hierarchy by using, e.g. path like notation; you can use the
`--prefix` flag to mimick recursive lists
* kubernetes `ResourceVersion` relates directly to the revision in etcd; which is a monotonically increasing counter
* kubernetes does not need to use etcd (others use [sqlite3](https://www.sqlite.org/index.html) or [Cosmos DB](https://docs.microsoft.com/en-us/azure/cosmos-db/introduction)), but etcd is the most common choice
* etcd comes with an HTTP and a gRPC API
* compaction in etcd is key to limit overall database size

Kubernetes uses etcd as a registry with components starting to watch specific keys for events.

* the difference between edge-driven (a sequence of events that leads to a
  state) and level-driven (transmit the whole state at once) design and its
reflection in the API (e.g. add, update, delete vs reconcile).

Some tools mentioned or used in the talk:

* [jq](https://stedolan.github.io/jq/), and its [base64](https://stedolan.github.io/jq/manual/#Formatstringsandescaping) capabilities
* [delta](https://github.com/dandavison/delta) for diffs
* [auger](https://github.com/jpbetz/auger) - encodes and decodes Kubernetes objects from the binary storage encoding used to store data to etcd
* [jaeger](https://www.jaegertracing.io/) or [zipkin](https://zipkin.io/) for tracing
* [zap](https://github.com/uber-go/zap) for structured logging

## Misc

* Does your favorite encoding schema support tree structures natively? Would you be interested in an encoding that supports one? Please share your ideas at [concise-encoding/issues/33](https://github.com/kstenerud/concise-encoding/issues/33)
* [Cuelang](https://cuelang.org/) - a language for configuration, database schemas, validation and more - e.g. to adress subtle issues like [different int max sizes](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/MAX_SAFE_INTEGER) across languages
* Two cloud native books: [Design Patterns for Cloud Native Applications](https://www.oreilly.com/library/view/design-patterns-for/9781492090700/) and [Cloud Native Patterns](https://www.oreilly.com/library/view/cloud-native-patterns/9781617294297/)
* [Peter Alvaro](https://scholar.google.com/citations?user=TKSjVTUAAAAJ&hl=en) at [Strange Loop](https://www.youtube.com/watch?v=R2Aa4PivG0g) (2015, Distributed Systems, Data Management Systems)

----

[Join our meetup](https://www.meetup.com/Leipzig-Golang) to get notified of upcoming events!

