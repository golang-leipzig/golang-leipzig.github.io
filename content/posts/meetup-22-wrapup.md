---
title: "Virtual Meetup #22 wrap-up"
date: 2021-10-31T14:00:00+01:00
draft: false
tags:
- summary
- meetup
---

## Confidential Computing with Go

We had a great presentation by [Moritz Eckert](https://twitter.com/m1ghtymo)
from [edgeless systems](https://edgeless.systems) about Confidential Computing
with Go, and especially [ego.dev](https://ego.dev), a modified Go compiler and
additional tooling which allow you to run code within a trusted execution
environment, e.g. using [Intel® SGX
enclaves](https://en.wikipedia.org/wiki/Software_Guard_Extensions) with zero
code changes.

Here's a recording of the talk, code walkthroughs and Q&A:

[![Leipzig Gophers #22: Confidential Computing with Go](https://img.youtube.com/vi/oycZLZdI8s8/0.jpg)](https://www.youtube.com/watch?v=oycZLZdI8s8)

A few highlights from the talk:

* a trusted execution environment can be implemented in various ways, Intel SGX
  being one that it usable today (with some others in development, e.g. [Arm
  CCA](https://www.arm.com/company/news/2021/06/arm-cca-will-put-confidential-compute-in-the-hands-of-every-developer), ...)
* there are various applications, from security and privacy enhancements (e.g.
  medical records; private contact discovery in
  [Signal](https://github.com/signalapp/ContactDiscoveryService; ...) to data
  sharing options between mistrusting parties (e.g. federated learning)
* the [SGX
  architecture](https://sgx101.gitbook.io/sgx101/sgx-bootstrap/overview) is
non-trivial &mdash; however, [ego](https://www.ego.dev/) reduces the amount of
work required for secure deployments of Go projects considerably
* [ego](https://ego.dev) builds on top of [OpenEnclave](https://openenclave.io/sdk/), an open source SDK that provides consistent API surface across enclave technologies

We looked at a few examples from the [ego.dev](https://github.com/edgelesssys/ego) project:

* [helloworld](https://github.com/edgelesssys/ego/tree/master/samples/helloworld)
* [remote attestation](https://github.com/edgelesssys/ego/tree/master/samples/remote_attestation), which shows how to connect to a secured server

Confidential computing is most likely becoming more deployed in the coming
years and it's really nice to be able to experiment and deploy solutions in Go
today.

----

[Join our meetup](https://www.meetup.com/Leipzig-Golang) to get notified of
upcoming events!

