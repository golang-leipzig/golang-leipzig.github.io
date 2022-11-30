---
title: "Meetup #32 wrap-up"
date: 2022-11-30T13:00:31+02:00
draft: false
tags:
- summary
- meetup
---

## Gridfuse x Leipzig Gophers


Leipzig Gophers Meetup #32 took place on [2022-11-29 19:00 CET][meetup] and we had a
great meetup at [Gridfuse'](https://gridfuse.com/) new office in the heart of
Leipzig. Gridfuse was founded in 2022 and builds software to bundle renewable
energy sources (wind, solar, ...) to react to market prices - in Germany and
Europe.

Predicting future loads is one key component, and renewables pose their own set
of challenges since they depend more on external factors (like the
weather). Bundling (many smaller) assets is important, too -- e.g. 1-20MW is a relevant target range.

Affected by the [European energy crisis
2022](https://www.consilium.europa.eu/en/infographics/eu-measures-to-cut-down-energy-bills/)
the previously used gas-powered plants, which used to stabilize the grid, are
faded out - and the need to attain the same stabilization goals with renewable energy sources
becomes much more urgent.

Gridfuse choose Go, since a few components like influxdb and NATS use Go, but
also because of an overall good performance profile of the language.

[![](/images/gridfuse-gophers-s.png)](https://www.linkedin.com/posts/lisaekern_gridfuse-leipziggophers-leipzigstartup-activity-7003703311040143360-PnjL?utm_source=share&utm_medium=member_desktop)

## Data Collectors and Flows

We heard from [Jörg](https://www.linkedin.com/in/j%C3%B6rg-werner-b49798105/)
about parts of their backend, data collection architecture.

Data flows from asset connectors, into [NATS](https://nats.io/), then
via [telegraf](https://docs.influxdata.com/telegraf/v1.24/) into
[influxdb](https://www.influxdata.com/).

The telegraf plugin has no capability to parse subjects from NATS messages,
only the payload, so Gridfuse started to implemented this feature and
contributed it to upstream:
[influxdata#12274](https://github.com/influxdata/telegraf/pull/12274). Subject parsing is similar to [MQTT](https://mqtt.org/).


Running influxdb is an interim solution, as historical data can also be put
into other (more static) storage like S3 or relational databases for the use case.

## Energy efficiency of programming languages

We had a lightning talk about a paper: [Ranking Programming Languages by Energy
Efficiency](https://haslab.github.io/SAFER/scp21.pdf)
([fatcat](https://fatcat.wiki/release/dccqbrxp55fozmzmqytgmgysaq)): [Go energy efficiency](https://github.com/miku/goenergy), and a few areas [where Go shines](https://github.com/miku/goenergy#where-go-shines).

## Thanks!

We'd like to thank [Gridfuse](https://gridfuse.com) for a wonderful tech+social event
and the insights into a fascinating domain - and last but not least the incredible Gopher cookies :)

[![](/images/gridfuse-gopher-cookies-s.png)](https://www.linkedin.com/posts/lisaekern_gridfuse-leipziggophers-leipzigstartup-activity-7003703311040143360-PnjL?utm_source=share&utm_medium=member_desktop)

----

[Join our meetup](https://www.meetup.com/Leipzig-Golang) to get notified of
upcoming events!



[meetup]: https://www.meetup.com/leipzig-golang/events/282941959/
