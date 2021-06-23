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
again. We had an input presentation on [etcd](https://etcd.io/) and did a deep dive into its
role in the container-orchestrator [kubernetes](https://kubernetes.io/).

We learned about the API of etcd, its data model, key value stores layours, append-only processing and compaction. A recording of the talk is available on youtube:

[![Leipzig Gophers #19: A Journey in the Kubernetes ListerWatcher Rabbit Hole](http://img.youtube.com/vi/Z9fwIzy0C_8/0.jpg)](https://www.youtube.com/watch?v=Z9fwIzy0C_8)

Some more highlights:

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

----

[Join our meetup](https://www.meetup.com/Leipzig-Golang) to get notified of upcoming events!

