---
title: "Meetup #4 wrap-up"
date: 2019-05-17T22:00:00+02:00
draft: false
tags:
- summary
- meetup
---

## Lightning Network Micropayments

Go is popular in the crypocurrency space. The well-rounded standard library and
great support for networking applications might be one reason. In [Meetup
#4](https://www.meetup.com/Leipzig-Golang-and-Cloud/events/261418733/)
[Philipp](https://github.com/philippgille) presented an introduction to
Bitcoin, [Lighting Network](https://en.wikipedia.org/wiki/Lightning_Network)
and a Go project, that makes micropayments for API usage really simple:
[ln-paywall](https://github.com/philippgille/ln-paywall).

An http middleware takes care of creating
[invoices](https://github.com/lightningnetwork/lightning-rfc/blob/master/11-payment-encoding.md)
per request, a complete example can be implemented in [a few
lines](https://github.com/philippgille/ln-paywall/blob/78fd1dfbf10f549a22f4f30ac7f68c2a2735e989/examples/ping/handlerfunc/main.go).
Middleware
[implementations](https://github.com/philippgille/ln-paywall/tree/master/examples/ping)
are available for [echo](https://echo.labstack.com/),
[gin](https://github.com/gin-gonic/gin) and other frameworks.  You can
experiment on testnet and with a [test wallet](https://htlc.me/) or build you own
private chain, although this a bit more involved.

[Philipp](https://github.com/philippgille) also wrote
[gokv](https://github.com/philippgille/gokv), a key value store that abstracts
away over twenty local and cloud based implementations.

![](/images/go-bitcoin-books.jpg)

## The zero value in the wild

Cloud providers SDK API differ in style, this might play a role in the decision
for a cloud provider as well.

Go has the concept of a [zero value](https://golang.org/ref/spec#The_zero_value):

> Each element of such a variable or value is set to the zero value for its
> type: false for booleans, 0 for numeric types, "" for strings, and nil for
> pointers, functions, interfaces, slices, channels, and maps. This
> initialization is done recursively, so for instance each element of an array
> of structs will have its fields zeroed if no value is specified.

While in general this is a useful feature, you lose the distinction between an
uninitialized and an empty value. One example for this can be found in an older
issue in the [Go Amazon SDK](https://github.com/aws/aws-sdk-go-v2):
[#114](https://github.com/aws/aws-sdk-go/issues/114).

> One big issue migration to this SDK from other community SDKs is that most of
> the string types are represented as string pointers. This make porting
> existing code difficult. Also, simple initialization become cumbersome ...

## Open Source, Companies and Sustainability

Many companies love open source, fewer like to take an active role in
development and maintenance of free software. However, there are many good
reasons to do so:

* a chance to improve software quality
* allow engineers to build a portfolio
* excellent way to attract the developers you want
* contribute to the strength of software ecosystems

The Go world is no different than others in this respect. Some developers of
great open source tools (e.g.
[Fatih Arslan](https://arslan.io/2018/10/09/taking-an-indefinite-sabbatical-from-my-projects/)) take sabbaticals,
others projects may go unmaintained.

One initiative for Go is [gof.rs](https://gof.rs/):

> The Gofrs (pronounced Gophers) is a community-formed group working together
> to better the entire Go ecosystem. Some of these efforts include picking up
> the maintenance of projects that are widely used or have large impact, as
> well as to consider new solutions to problems that arise as the number of Go
> programmers continues to grow.

> We initially formed in the summer of 2018 to take one of the most popular
> UUID packages in the Go ecosystem,
> [github.com/satori/go.uuid](https://github.com/gofrs/uuid), and have started
> to look at contributing to more projects.

## References

* [ln-paywall](https://github.com/philippgille/ln-paywall)
* [gokv](https://github.com/philippgille/gokv)
* [HTTP 402](https://httpstatuses.com/402)
* [The zero value](https://golang.org/ref/spec#The_zero_value)
* [The Gofrs](https://gof.rs)

