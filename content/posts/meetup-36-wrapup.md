---
title: "Hybrid Meetup #36 wrap-up"
date: 2023-06-22T01:00:00+02:00
draft: false
unlisted: true
tags:
- summary
- meetup
---

## Cut Cloud Complexity with Encore

Hybrid meetup #36 took place [2023-06-20
19:00](https://www.meetup.com/leipzig-golang/events/290666177/) at [Basislager
Leipzig](https://www.basislager.co/).

We had a great presentation from [André
Eriksson](https://www.linkedin.com/in/erikssonandre/) about
[Encore](https://encore.dev/) - a platform written in Go and for Go for
radically simplifying cloud-native backend development.

Beside motivation and overview, André live-coded an [uptime monitoring
service](https://github.com/encoredev/examples/tree/main/uptime), ran it
locally and deployed it onto Google Cloud within a few minutes. A [recording of
the talk](https://youtu.be/yYEXDmy3zUA) is available on YouTube!

<iframe width="560" height="315" src="https://www.youtube.com/embed/yYEXDmy3zUA" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

A few more takeaways from the talk and discussion:

* most of Encore is open source: [github.com/encoredev](https://github.com/encoredev/)
* Encore can be cloud agnostic by focussing on the 80% cases, that most cloud applications need: see [primitives](https://encore.dev/docs/primitives) and [production infrastructure](https://encore.dev/docs/deploy/infra#production-infrastructure)
* Encore does code to code transformations via static analysis and leverages Go's static typing to make working with services save and efficient
* currently, the large cloud providers are supported but on-premise deployments will be possible soon (deployment target will be an on-premise kubernetes cluster)

Some more insights from the Q&A:

* static analysis of Go code is easier that static analysis of Python code
* a large number of projects need only core cloud primitives
* Encore app can be deployed side by side existing infrastructure, and can share access to resources, e.g. message brokers

The space of tools to reduce cloud complexity for developers is large, with
other projects in this spaces being [ampt](https://getampt.com/), [Terraform,
Pulumi](https://encore.dev/docs/other/vs-terraform) and others.

Thanks again to [André](https://www.linkedin.com/in/erikssonandre/) for taking the time to demo an amazing project.

----

Have you reduced the complexity of your deployments? Then you can [join our meetup](https://www.meetup.com/Leipzig-Golang/) in the time you saved.

