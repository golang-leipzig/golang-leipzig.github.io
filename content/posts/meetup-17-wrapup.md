---
title: "Virtual Meetup #17 wrap-up"
date: 2021-02-20T22:00:00+02:00
draft: false
tags:
- summary
- meetup
---

## A data engineering short story: (Fuzzy) Matching with command line tools and Go

Meetup #17 took place Apr 20, 2021 19:00 CEST, and was virtual again (crossing
one year of virtual meetups).  We had a (longer) lightning talk on a data
engineering topic:

* [A data engineering short story: (Fuzzy) Matching with command line tools and Go](https://gist.github.com/miku/fb429faad8b856caf6bba5305af024df)

How to build a graph dataset with about 1B nodes from semi structured data?
With Taco Bell style programming, you can reuse (UNIX) command line tools and
sprinkle a few additional Go tools.

The graph concern citations, so we looked at a part of the citation of:

> Hoare, Charles Antony Richard. "Communicating sequential processes." Communications of the ACM 21.8 (1978): 666-677.

![](https://i.imgur.com/6dSaW2q.png)

The custom tool exploits *sorted keys* and works in a merge sort style way to
run compution on groups of items with the same key. One might consider key
extraction a mapping and grouping operations a reduce step.

## Graph stores and algorithms

Are there interesting graph libraries and project written in Go?

* [https://github.com/cayleygraph/cayley](https://github.com/cayleygraph/cayley)
* [https://github.com/dgraph-io/dgraph](https://github.com/dgraph-io/dgraph)

A generic data science umbrella project is: [Gonum](https://www.gonum.org) -
Consistent, composable, and comprehensible scientific code. It contains a
package for [graph
processing](https://github.com/gonum/gonum/tree/master/graph) as well.

Some project in other languages include:

* [https://dbs.uni-leipzig.de/research/projects/gradoop](https://dbs.uni-leipzig.de/research/projects/gradoop)
* [https://networkx.org/](https://networkx.org/)

Sometimes people write custom code for specific algorithms, e.g. for [pagerank](https://github.com/miku/pgrk).

## Misc

* The [GOLAB conference](https://golab.io) hosts free (and paid) webinars over
  the coming months: [https://golab.io/agenda/](https://golab.io/agenda/), e.g.
[Go and
Tensorflow](https://golab.io/agenda/addestrare-e-distribuire-modelli-tensorflow-in-go/)
* Go garbage notes: [https://blog.golang.org/ismmkeynote](https://blog.golang.org/ismmkeynote)

Data stores and analytics engines (outside Go):

* [sqlite](https://www.sqlite.org/index.html)
* [Apache Drill](https://drill.apache.org/)
* [Presto](https://prestodb.io/)

Tiny, useful tools:

* [jq](https://stedolan.github.io/jq/) (json pretty printer and processing
  tool), [ijq](https://sr.ht/~gpanders/ijq/) - interactive version

A couple of reading recommendations:

* [Surprisingly Slow](https://gregoryszorc.com/blog/2021/04/06/surprisingly-slow/)

Some research questions:

* [ ] good caching libraries (e.g. for HTTP and other data), beside [LRU](https://github.com/hashicorp/golang-lru)
* [ ] how to write parsers (e.g. for DSL or markup languages like [simpleml](http://simpleml.com/) - example library: [https://github.com/alecthomas/participle](https://github.com/alecthomas/participle)

Misc in Go and other languages

* [Learning Parser Combinators With Rust](https://bodil.lol/parser-combinators/)
* [The Most Beautiful Program Ever Written](https://www.youtube.com/watch?v=OyfBQmvr2Hc) (Scheme)
* [Boundaries](https://www.destroyallsoftware.com/talks/boundaries), Gary Bernhardt

----

[Join our meetup](https://www.meetup.com/Leipzig-Golang) to get notified of upcoming events!

