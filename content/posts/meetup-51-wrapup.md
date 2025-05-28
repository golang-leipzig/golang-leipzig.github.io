---
title: "Hybrid Meetup #51 wrap-up"
date: 2025-05-28T10:00:00+01:00
draft: false
tags:
- summary
- meetup
---

## Flaky LLMs and secure containers

Hybrid Meetup #51 took place
[2025-05-27](https://www.meetup.com/leipzig-golang/events/306803728) 19:00 at
renewable energy startup [Gridfuse](https://gridfuse.com) and we had two
shorter presentations:

* Beyond Benchmarks: [Human x LLM for Go and other code](https://github.com/miku/nightjet/blob/main/notes/2025-05-27-lgo-51-short-talk.md) (intermediate report)
* StackRox architecture and repository setup

### LLM for coding related tasks

LLM use for coding remains constrained to well scoped problems with some human domain knowledge and
oversight and review. It is still possible to spot generated code, it lacks a kind of handwriting.

Using an LLM can teleport you into a codebase and let you explore or ask
specific questions. For code generation, a disciplined approach would be to
review every line of output.

[![](/images/lgo-51-llm-palm-montage.png)](https://github.com/miku/nightjet/blob/main/notes/2025-05-27-lgo-51-short-talk.md)

This was an intermediate report and we hope, we can have a more comprehensive
assessment of the effects on (Go) software development in the future.

### Container security and StackRox architecture bits

We did some architecture and code review along [diagrams](https://github.com/stackrox/stackrox/blob/92e5d0badaf6a86f0691ad39b739fe233ed193bc/central/platform/reprocessor/singleton.go#L11-L25) and [code](https://github.com/stackrox/stackrox):

[![](/images/lgo-51-acs-architecture-scannerv4.png)](https://docs.redhat.com/en/documentation/red_hat_advanced_cluster_security_for_kubernetes/4.6/html/architecture/acs-architecture)

Some notes:

* [Open Policy Agent](https://www.openpolicyagent.org/) could have been a component, but for the StackRox use case it was deemed to slow at the time
* StackRox is an example of a vertically scaled project (albeit at some point in the past they had a microservices architecture)
* data structures are modeled in with [protobuf](https://protobuf.dev/), e.g. for [storage](https://github.com/stackrox/stackrox/tree/master/proto/storage), other code is then generated from these definitions
* we looked at a singleton service pattern, e.g. used here: [singleton.go](https://github.com/stackrox/stackrox/blob/92e5d0badaf6a86f0691ad39b739fe233ed193bc/central/platform/reprocessor/singleton.go#L11-L25)

Meetup [#48](/posts/meetup-48-wrapup/) also gave a high level overview of StackRox.


### Misc

* [What Is ChatGPT Doing ... and Why Does It Work?](https://openlibrary.org/books/OL47286904M/What_Is_ChatGPT_Doing_..._and_Why_Does_It_Work)
* [SWE-Bench](https://paperswithcode.com/search?q_meta=&q_type=&q=swe-bench)


### Thanks, Gridfuse!

Thanks a lot to [Gridfuse](https://gridfuse.com) for providing a great venue
and atmosphere for our meetup (and
[others](https://www.meetup.com/leipzig-devops/))!

----

[Join our meetup](https://www.meetup.com/de-DE/leipzig-golang/) to get notified of upcoming events.
