---
title: "Virtual Meetup #28 wrap-up"
date: 2022-06-22T14:00:00+02:00
draft: false
tags:
- summary
- meetup
---

## Errors and Glamour

[Meetup #28](https://www.meetup.com/leipzig-golang/events/282941935/) took
place on 2022-06-21 19:00 CEST virtually and was one of our most visited meetup to date - thank you all for dropping by.

We had a short presentation *Never forget to handle errors, rediscovered the
1000th time* by [Leandro](https://www.linkedin.com/in/leandrosansilva/) from
[Ligthmeter](https://lightmeter.io/) ([gitlab](https://gitlab.com/lightmeter/),
[jobs](https://www.ycombinator.com/companies/lightmeter/jobs/PB5qMli-principal-network-engineer))
about a few ways to handle errors in Go.  The compiler does not complain about
unchecked errors, but some [linters do](https://github.com/kisielk/errcheck).

The presentation can be found here:
[presentation.md](https://gitlab.com/leandrosansilva/talks/-/blob/master/errors-golang/presentation.md)

Thanks [Leandro](https://www.linkedin.com/in/leandrosansilva/) for your great input!

A discussion about errors reminded us that as a first measure, errors need to be
handled, actually. There's even some research on that - e.g. compiled here into
Software Engineering's Greatest Hits:
[https://youtu.be/HrVtA-ue-x0?t=595](https://youtu.be/HrVtA-ue-x0?t=595).

We had another input talk about [Text User
Interfaces](https://github.com/miku/glamline), using the
[charm](https://github.com/charmbracelet/) libraries as example. Text
interfaces are nice (and retro) but their widespread adoption is debatable.
Nonetheless, these interface can be efficient and pleasant to use.

The presentation showed some example, e.g. [glow](https://github.com/charmbracelet/glow) for reading markdown:

```shell
$ glow tinyurl.com/m-28-wp
```

or a stock [ticker app](https://github.com/achannarasappa/ticker):

![](/images/ticker.png)

----

[Join our meetup](https://www.meetup.com/Leipzig-Golang) to get notified of
upcoming events!

