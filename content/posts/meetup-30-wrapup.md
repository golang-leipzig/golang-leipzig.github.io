---
title: "Hybrid meetup #30 wrap-up"
date: 2022-09-23T14:00:00+02:00
draft: false
tags:
- summary
- meetup
---

## Secure data with immudb

Hybrid Meetup #30 took place [2022-09-20 19:00
CEST](https://www.meetup.com/leipzig-golang/events/286871365/) at [Lancaster
University Leipzig](https://www.lancasterleipzig.de/) Campus.

We had a great presentation by [Dennis
Zimmer](https://www.linkedin.com/in/denniszimmer) and [Bartłomiej
Święcki](https://www.linkedin.com/in/bswiecki/)
([GitHub](https://github.com/byo)) from [CodeNotary](https://codenotary.com/)
about their cutting edge open source database project
[immudb](https://immudb.io/).

Presentation slides are available here:

[![](/downloads/codenotary-immudb_Sep-22_dz_1.pptx_cover.png)](https://golangleipzig.space/downloads/codenotary-immudb_Sep-22_dz_1.pptx.pdf)

Originally, the CodeNotary team looked at blockchains to implement the idea of
tamperproof data store - but found that they could build something more performant.
This work resulted in [immudb](https://immudb.io), a lightweight, high-speed
immutable database which is written in Go and uses
[Merkle-Trees](https://en.wikipedia.org/wiki/Merkle_tree) at its core.

Starting with a key-value database, immudb by now supports a dialect and subset
of SQL as well and can speak the Postgres [wire
protocol](https://www.postgresql.org/docs/current/protocol.html). Some
challanges include the immutable representation of statements like `ALTER
TABLE` and the like.

The project takes advantage of Go's excellent deployment story and is set up in
a few minutes. Currently, [SDKs](https://github.com/codenotary/immudb#how-to-integrate-immudb-in-your-application) are available for
Java, Go, .NET, Python and nodejs - which makes integration in applications
much simpler (see also
[examples](https://github.com/codenotary/immudb-client-examples)).

To ensure the validity of the data, clients and server work together: Clients
can keep track of the state of the database (which takes only a few bytes) and
can use this to validate against the server.

To learn more about the project and the internals of immudb, please also check
out the FOSDEM 2022 talk by Bartłomiej Święcki: [Don't trust us, trust the math
behind immudb: How immudb protects safety critical data (with math and
cryptography)](https://archive.fosdem.org/2022/schedule/event/safety_dont_trust_us_trust_the_math_behind_immudb/)
- and the paper [immudb: A Lightweight, Performant Immutable
Database](http://codenotary.s3.amazonaws.com/Research-Paper-immudb-CodeNotary_v3.0.pdf).

Apart from the tech-talk and discussion, we gave away goodies: A
[Zimaboard](https://www.zimaboard.com/), courtesy of
[CodeNotary](https://codenotary.com/) and a copy of [Cloud Native
Go](https://learning.oreilly.com/library/view/cloud-native-go/9781492076322/),
courtesy of [O'Reilly Media](https://www.oreilly.com/pub/cpc/323592).

We may feature a project built with immudb and the winner of the Zimaboard in
the future and hope *Cloud Native Go* will be a page turner!

![](/images/meetup-30-tile.png)


## Misc

* Thanks to [Lancaster University Leipzig](https://www.lancasterleipzig.de/) for providing a meetup space on a short notice (change of plans
caused by a [bomb
find](http://web.archive.org/web/20220923115054/https://www.mdr.de/nachrichten/sachsen/leipzig/bombe-entschaerft-evakuierung-sperrkreis-polizei-100.html)
near [Basislager](https://www.basislager.co/) - our original meetup location).
* Security of systems can increase, when the number of witnesses go up. A
  curious example: The New York Times has been hiding a blockchain since 1995,
apparently: [The World’s Oldest Blockchain Has Been Hiding in the New York
Times Since
1995](https://www.vice.com/en/article/j5nzx4/what-was-the-first-blockchain)

----

[Join our meetup](https://www.meetup.com/Leipzig-Golang) to get notified of
upcoming events!


