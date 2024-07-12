---
title: "Hybrid Meetup #43 wrap-up"
date: 2024-05-29T13:00:00+01:00
draft: false
tags:
- summary
- meetup
---

## Finite state machines

Hybrid Meetup #43 took place
[2024-05-28](https://www.meetup.com/leipzig-golang/events/298066360/) at
[Gridfuse](https://gridfuse.com), [Hauptpost](https://de.wikipedia.org/wiki/Hauptpost_(Leipzig)) Leipzig. We had a
great presentation about a lightweight (~200 LOC), open source state machine
library written in Go: [ekstatic](https://github.com/Metamogul/ekstatic) &mdash; presentation [slides (PDF)](https://golangleipzig.space/downloads/ekstatic.pdf)

![](/meetup-43-fsm/crop.png)

The use case in the context of [RAIDA](https://raida.de/) is the modelling of
workflows, composed of a number of steps (and the loose coupling of business
logic with state transistions). The library is flexible, but the limits of the
Go type system require to resort to
[any](https://go.dev/ref/spec#Predeclared_identifiers) (which will defer some
type checks to runtime).

> For convenience, the predeclared type `any` is an alias for the empty
> interface. [[Go 1.18](https://go.dev/ref/spec#Go_1.18)]

The library is open source at
[Metamogul/ekstatic](https://github.com/Metamogul/ekstatic), and contains some
[examples](https://github.com/Metamogul/ekstatic/tree/main/examples) as well. Thanks again, [Jan](https://www.linkedin.com/in/jan-z-255b28225/) for the insights,
and [Gridfuse](https://gridfuse.com/) for hosting another Leipzig Gophers event.


## Misc

* code generation may improve type safety of a generic FSM Go implementation, similar to [sqlc](https://sqlc.dev/) and other libraries
* [railway oriented control flow](https://vimeo.com/113707214), reminding one of [The happy path is left-aligned](https://maelvls.dev/go-happy-line-of-sight/)
* Rust [std::result](https://doc.rust-lang.org/std/result/) in combination with
  the [? operator](https://doc.rust-lang.org/std/result/#the-question-mark-operator-)
makes working with errors easier; there are libraries in Go, like
[alexthomas/types](https://pkg.go.dev/github.com/alecthomas/types/result), that
try emulate that; or even language [proposals](https://github.com/golang/go/issues/19991) (declined at the time)
* another, albeit much more extensive library for workflows (or durable execution) is [temporal](https://temporal.io/), which has a [go](https://learn.temporal.io/getting_started/go/) client as well
* contracts has been tried in various places in software systems, e.g. [design by contract](https://en.wikipedia.org/wiki/Design_by_contract), or [PACT](https://docs.pact.io/) for REST testing
* GraphQL is a nice alternative to representional state transfer (REST); libraries: [gqlgen](https://gqlgen.com/getting-started/) for servers, [gqlgenc](https://github.com/Yamashou/gqlgenc) for clients
* Rob Pike on state machines, 12 years ago: [Lexical Scanning in Go - Rob Pike](https://youtu.be/HxaD_trXwRE), [What is a state? [13:45]](https://youtu.be/HxaD_trXwRE?si=wIGRoDxp78G8eXn4&t=825)

## Monads

A finite state machine consists of states and state *transitions*;
implementationwise, a state may be *any* type, but at the same time we would
benefit from marking a type as a state, hence unifying different actions. A [variant type](https://www.cs.cornell.edu/courses/cs3110/2013sp/lectures/lec04-types/lec04.html) could be a solution, but [Go does not support
variant types](https://go.dev/doc/faq#variant_types) (albeit interfaces and type
switches allow for some unification).

In [essence](https://www.st.cs.uni-saarland.de/edu/seminare/2005/advanced-fp/docs/wadler-essence-fp.pdf),
we would like chainable computation, or workflow composition, which is reminding of monads.

* [the morning paper: Monads for functional programming](https://blog.acolyer.org/2014/11/10/monads-for-functional-programming/)
* [A monad is just a monoid in the category of endofunctors, what's the problem?](https://stackoverflow.com/q/3870088/89391)

Sidenote: For [Leibniz](https://en.wikipedia.org/wiki/Gottfried_Wilhelm_Leibniz), monads
were the essential substance
([1714](https://www.plato-philosophy.org/wp-content/uploads/2016/07/The-Monadology-1714-by-Gottfried-Wilhelm-LEIBNIZ-1646-1716.pdf)),
which *has no parts and is therefore indivisible*. They also were windowless, or immutable.

> The Monads have no windows, through which anything could come in or go out.

A few years earlier, in 1703, Leibniz recognized something else essential: the
binary number system.

[![](/images/Leibniz_binary_system_1703-s.png)](https://www.leibniz-translations.com/static/media/binary.2e9089b038d26b586697.pdf)

> But reckoning by twos, that is, by 0 and 1, as
compensation for its length, is the most fundamental way of reckoning for
science, and offers up new discoveries, which are then found to be useful, even
for the practice of numbers and especially for geometry. The reason for this is
that, as numbers are reduced to the simplest principles, like 0 and 1, a
wonderful order is apparent throughout.


