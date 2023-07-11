---
title: "Hybrid Meetup #36 wrap-up"
date: 2023-06-30T10:00:00+02:00
draft: false
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

Beside motivation and overview, André [live-coded](https://youtu.be/yYEXDmy3zUA?t=1824) an [uptime monitoring
service](https://github.com/encoredev/examples/tree/main/uptime), ran it
locally and deployed it onto Google Cloud within a few minutes. A [recording of
the talk](https://youtu.be/yYEXDmy3zUA) is available on YouTube!

[![](/images/36-yt-thumb.png)](https://youtu.be/yYEXDmy3zUA)

A few more takeaways from the talk and discussion:

* most of Encore is open source: [github.com/encoredev](https://github.com/encoredev/)
* Encore can be cloud agnostic by focussing on the 80% cases, that most cloud applications need: see [primitives](https://encore.dev/docs/primitives) and [production infrastructure](https://encore.dev/docs/deploy/infra#production-infrastructure)
* Encore does code to code transformations via static analysis and leverages Go's static typing to make working with services save and efficient
* currently, the large cloud providers are supported but on-premise deployments [are possible](https://encore.dev/docs/how-to/import-kubernetes-cluster) <strike>will be possible soon (deployment target will be an on-premise kubernetes cluster) </strike>
* Go is a language made for tools, with a [regular grammar](https://go.dev/talks/2009/go_talk-20091030.pdf#page=14), which makes it easier to write code that manipulates Go code

Find out more:

* quick video showing [Encore in action](https://www.youtube.com/watch?v=IwplIbwJtD0)
* more [example projects](https://github.com/encoredev/examples)

Some more insights from the Q&A:

* static analysis of Go code is easier that static analysis of Python code
* a large number of projects need only core cloud primitives
* Encore app can be deployed side by side existing infrastructure, and can share access to resources, e.g. message brokers

The space of tools to reduce cloud complexity for developers is becoming larger, with
projects in this space being [ampt](https://getampt.com/), [Terraform,
Pulumi](https://encore.dev/docs/other/vs-terraform) and others. One open source kit in the cloud-agnostic realm is the [Go Cloud Development
Toolkit](https://gocloud.dev/), which also aimed at *providing commonly used,
vendor-neutral generic APIs that you can deploy across cloud providers*.

Thanks again to [André](https://www.linkedin.com/in/erikssonandre/) for taking the time to demo an amazing project.

----

Have you reduced the complexity of your deployments? Then you can [join our meetup](https://www.meetup.com/Leipzig-Golang/) in the time you saved.

