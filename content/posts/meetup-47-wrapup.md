---
title: "Drinkup #47 wrap-up"
date: 2024-12-20T10:00:00+01:00
draft: false
tags:
- summary
- meetup
---

## Misc

Drinkup #47 took place [2024-12-19](https://www.meetup.com/leipzig-golang/events/298066373) 19:00 at
[Cantona Leipzig](https://www.google.com/search?q=cantona+leipzig).

* LLM coding assistants became so mainstream this year that people are starting to turn them off
  (even if they are given away for free); but new interaction modes like *canvas*
or deeper integration with your *code project* or *IDE* feel a bit like having swimming
fins 🧜
* do you know about software that is usable, used and finished? haha! [tell us about it!](https://github.com/golang-leipzig/software-that-is-finished) -- the [DoD would also listen!](https://media.defense.gov/2019/May/01/2002126690/-1/-1/0/SWAP%20EXECUTIVE%20SUMMARY.PDF) cf. [HN](https://news.ycombinator.com/item?id=34558707)
* well, [are we really engineers](https://www.hillelwayne.com/post/are-we-really-engineers/), then [why can't we make simple software](https://www.youtube.com/watch?v=czzAVuVz7u4)? Does a [time travelling debugger](https://www.youtube.com/watch?v=NiGzdv84iDE) help? If you think the old times were better, or less strange, please check out [The Old New Thing: Practical development throughout the evolution of Windows](https://devblogs.microsoft.com/oldnewthing/)
* if you are developing software, you are going through many different stages
  along different axes, with one axis being testing: you start with enjoying unit test
and you end with mostly focussing on integration tests
* [Greg Wilson](https://third-bit.com/) nicely summarizes a lots of empirical research on code and
  sofware development into a few slides, like this one:
[greatest-hits/#20](https://third-bit.com/talks/greatest-hits/#20): [SLOC](https://en.wikipedia.org/wiki/Source_lines_of_code) is still OK, more on that complex topic: [From Code Complexity Metrics to Program
Comprehension](https://dl.acm.org/doi/pdf/10.1145/3546576)
* speaking of SLOC, *did you know* that [wireguard](https://www.wireguard.com/), the VPN tunnel protocol,
  shines on this metric: [3904 vs
116730](https://www.wireguard.com/talks/fosdem2017-slides.pdf#4) lines of code, compared to OpenVPN (2017,
wireguard has been added Linux
[5.6](https://kernelnewbies.org/Linux_5.6#WireGuard.2C_a_faster.2C_simpler.2C_secure_VPN))
* speaking of wireguard, *did you know* that you can run a [tailscale control server](https://tailscale.com/opensource) yourself, e.g. with
[headscale](https://headscale.net/)? both tailscale and headscale are written
in Go, so we may explore this topic further in an upcoming event (the name tailscale is *flipped* from the 2013 Google paper [The tail at scale](https://www.barroso.org/publications/TheTailAtScale.pdf))

> The long tail of products are never going to be that big [ie. as a
> hyperscaler] &mdash; almost everybody building almost everything doesn't have
> any of those problems. -- [Avery Pennarun](https://apenwarr.ca/log/) at [45:56...](https://www.buzzsprout.com/1822302/episodes/9890092-tailscale-with-avery-pennarun-brad-fitzpatrick)

[...] Which *may* mean that we will see more people regaining ownership over their personal
data (or even more [cloud exits](https://www.bsi.bund.de/SharedDocs/Downloads/DE/BSI/KRITIS/UPK/upk-exit-strategie-cloud-dienstleistungen.pdf?__blob=publicationFile&v=8) (de)), because it gets constantly cheaper to store, host and run stuff on own
*infra* -- I may be overly optimistic, but then
[208K](https://github.com/awesome-selfhosted/awesome-selfhosted) ⭐ also may
mean something, after all. The aspect of privacy can also be put at the key of
products, like [ente.io](https://github.com/ente-io/ente), a *fully open
source, End to End Encrypted alternative to Google Photos and Apple Photos*.

Note:

> The people predicting the end of Moore's law is doubling every 18 months. --
> [[Jim Keller](https://en.wikipedia.org/wiki/Jim_Keller_(engineer)), maybe]

Anyone up for a challenge to implement an [array programming
language](https://en.wikipedia.org/wiki/Array_programming) in Go? Or
maybe [presenting](https://golangleipzig.space/proposals/) something
yourself?

In any way, subscribe to our [RSS
feed](https://golangleipzig.space/posts/index.xml) or [join our meetup
group](https://www.meetup.com/de-DE/leipzig-golang/) to get notified of
upcoming events. Happy holidays.

[![](/images/EIJOPZQVXHNVF4FIQBUGSLCZFUEC57JN.gif)](https://gifcities.org)

