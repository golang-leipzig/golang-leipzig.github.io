---
title: "Hybrid Meetup #50 wrap-up"
date: 2025-05-01T08:00:00+01:00
draft: true
tags:
- summary
- meetup
---

## Live and Let Die: The Mail is Not Enough

Hybrid Meetup #50 took place
[2025-04-29](https://www.meetup.com/leipzig-golang/events/306803728) 19:00 at
[Basislager Leipzig](https://basislager.co) and we had an insightful
presentation on *How a Go project dies; lessons learned, challenges to revive
it and the scars I got on the way*.

[Leandro](https://www.linkedin.com/in/leandrosansilva/) tooks us on a journey
into the [postfix](https://www.postfix.org/) mail transport agent (MTA) monitoring project
[controlcenter](https://gitlab.com/leandrosansilva/controlcenter), written in
Go, how it started and factors that contributed to its stalling. The company
around it, [Lightmeter](https://lightmeter.io/), was part of the [YC-W22
batch](https://www.ycombinator.com/companies/lightmeter), and covered by
[heise.de](https://www.heise.de/news/Lightmeter-Neues-Monitoring-Werkzeug-fuer-den-E-Mail-Server-4647424.html),
[Forbes](https://www.forbes.com/sites/davidjeans/2021/03/01/elastic-war-on-amazon-web-services/)
and others.

The slides can be found here:

> XXX

Some more takeaways:

* Monoliths are still ok!
* A monolith can look like a set of microservices too; with one database
  per component and intra-process communication through channels
* Reusability as beneficial [emergent property](https://www.dcs.gla.ac.uk/~johnson/papers/RESS/Complexity_Emergence_Editorial.pdf) -- and not a goal in and of itself
* Parsing postfix logs can be a challenge (cf. "[A User-Extensible and
  Adaptable Parser Architectur](https://scispace.com/pdf/a-user-extensible-and-adaptable-parser-architecture-4nocz4o09k.pdf)")
* Controlcenter took inspiration from Apple's [Grand Central
  Dispatch](https://en.wikipedia.org/wiki/Grand_Central_Dispatch) for handling
concurrency and taking advantage of multicore machines in the presence of
limitations (e.g. *usually, SQLite allows at most one writer to proceed
concurrently*)
([docs](https://www.sqlite.org/src/doc/begin-concurrent/doc/begin_concurrent.md))
* Sqlite3 [does not implement](https://sqlite.org/forum/info/78a60bdeec7c1ee9)
  stored procedures, but it has [application defined
functions](https://sqlite.org/appfunc.html); with the Go sqlite3 wrapper, you
can use
[RegisterFunc](https://pkg.go.dev/github.com/mattn/go-sqlite3#SQLiteConn.RegisterFunc)
to blend Go and SQL
* The universe of possible postfix setups is huge, and it is impossible to
  anticipate the layout of all log lines the application would ever encounter;
use a crash-first approach to iterate; implemented with a custom assertion
package (since Go does not have
[assertions](https://en.wikipedia.org/wiki/Assertion_(software_development)))
* Early returns, or guard clauses, or the [left-hand rule](https://scribe.rip/@matryer/line-of-sight-in-code-186dd7cdea88) can lead to *else-less code* -- also also avoid the [arrow anti-pattern](http://wiki.c2.com/?ArrowAntiPattern)
* Controlcenter decided against ORM: tighter coupling, but less complexity and more transparency
* Unfortunately, great software with users does not automatically translate to a product with customers

The open source project is looking for contributors. If you are into mailops, then please take
a look at [controlcenter repository](https://gitlab.com/leandrosansilva/controlcenter)


## Misc

* Google [tried
  out](https://opensource.googleblog.com/2023/03/introducing-service-weaver-framework-for-writing-distributed-applications.html)
to a monolithic microservice approach with the now discontinued [Service
Weaver](https://serviceweaver.dev/):

> Under the covers, Service Weaver will dissect your binary along component
> boundaries, allowing different components to run on different machines. -- [docs](https://serviceweaver.dev/docs.html)

* Go has an a bit unusual set of ORM libraries: classical ones like
  [GORM](https://gorm.io/index.html), more code-generation oriented ones like
[sqlc](https://sqlc.dev/) or [ent](https://github.com/ent/ent); fluent query
builder, like [bob](https://github.com/stephenafamo/bob), or thin wrappers,
like [sqlx](https://github.com/jmoiron/sqlx); generic
[goe](https://github.com/go-goe/goe) and more
* Object-relational mapping works best until it does not
* Hosting an email server yourself is not too complicated (cf. [mail-in-a-box](https://github.com/mail-in-a-box/mailinabox)), but there is a risk
  being blacklisted and then miss out of important messages
* If you run your own email server, you may receive less spam
* [Is email dead](https://www.emailisnotdead.com/)? Hint: while you read this
  blog post, the world moved 115,403,640 (legitimate) emails forward.

Thanks again to [Leandro](https://www.linkedin.com/in/leandrosansilva/) for the
talk and insights!



[Join our meetup](https://www.meetup.com/de-DE/leipzig-golang/) to get notified of upcoming events.
