---
title: "Hybrid Meetup #46 wrap-up"
date: 2024-11-27T13:00:00+01:00
draft: false
tags:
- summary
- meetup
---

## Log, Transform, Monitor: Journals, CI and serverless databases

Hybrid Meetup #46 took place
[2024-11-26](https://www.meetup.com/leipzig-golang/events/298481354/) 19:00 at
[Basislager](https://www.basislager.co/) and we had a nice mix of people, both
online and offline, with different backgrounds in addition to two great
presentations.

### Go journald exporter

[Leandro](https://www.linkedin.com/in/leandrosansilva/) motivated the need for
a tool that joins a legacy C++ application with modern event monitoring tool
like [Prometheus](https://github.com/prometheus/prometheus), through systemd,
specifically
[journald](https://man7.org/linux/man-pages/man8/systemd-journald.service.8.html):
[go-journald-exporter](https://gitlab.com/leandrosansilva/go-journald-exporter).

journald has a [native protocol](https://systemd.io/JOURNAL_NATIVE_PROTOCOL/)
that allows processes to write directly to the log;
[go-journald-exporter](https://gitlab.com/leandrosansilva/go-journald-exporter)
will use [go-systemd](https://github.com/coreos/go-systemd) bindings to seek to
the end of the journal, listen for specific messages and transform them to
prometheus metrics, which the tool exposes through an HTTP endpoint (because Prometheus is
[pull-based](https://prometheus.io/docs/introduction/faq/#why-do-you-pull-rather-than-push)).

This way, a lightweight [tool of a few hundred lines of code](https://gitlab.com/leandrosansilva/go-journald-exporter/), can offer better
visibility into a fleet of legacy applications running low resource
environments, where go is viable and a performant option, too.

> More details [in the slides](/downloads/leipzig-gophers-46-2024-11-26-go-journald-exporter.html).

### From GitHub workflow JSON to postgres to grafana

[Fedor](https://www.linkedin.com/in/fedor-dikarev/) turned GitHub CI run
results (taken from the API, JSON format) into much more comprehensible visual
aggregations with a small tool:
[gh-workflow-stats-action](https://github.com/neondatabase/gh-workflow-stats-action/).

[![](/images/meetup-46-screenie.png)](https://github.com/neondatabase)

It reads data off the GitHub API, populates a postgres database and uses
[grafana
postgres](https://grafana.com/docs/grafana/latest/datasources/postgres/) data
source to aggregate and visualize metrics from CI runs.

For larger projects, GitHub API rate limiting needs to be accounted for, but
it's possible to turn CI output from a larger project into database tables over
a weekend.

As a nice recursive twist, the tool used [neon.tech](https://neon.tech/) to
create a postgres database for storing the workflow run results in the cloud
with one click: you get a [DSN](https://en.wikipedia.org/wiki/Data_source_name) and you are ready to go.
[Neon](https://github.com/neondatabase/neon) splits storage and compute
components of postgres and can unlock lots of useful features this way, like
scaling and branching. With branching, the
[WAL](https://www.postgresql.org/docs/current/wal-intro.html) gets written
forward differently for different branches, which is useful for testing and
staging environments.

On the side: neon.tech is [hiring](https://neon.tech/careers)!

### Thanks for the tools!

[SLOC](https://github.com/XAMPPRocky/tokei) of the tools presented: 484 (go-journald-exporter) and 622 SLOC
(gh-workflow-stats-action), both help to turn streams of diagnostics into
suitable formats for modern monitoring tools and are adaptable in own projects.

Thanks again to [Leandro](https://www.linkedin.com/in/leandrosansilva/) and
[Fedor](https://www.linkedin.com/in/fedor-dikarev/) for the inspiring talks and
demos.


### Misc

* kubernetes test grid: [https://testgrid.k8s.io/sig-release](https://testgrid.k8s.io/sig-release)
* another CI tool (written in Go): [concourse](https://concourse-ci.org/)
* more CI tools: [https://argoproj.github.io/cd/](https://argoproj.github.io/cd/), [https://fluxcd.io/](https://fluxcd.io/)
* [CI/CD Observability using OpenTelemetry](https://www.cloudraft.io/blog/cicd-observability-using-opentelemetry)
* go build infrastructure is open source, too: [build](https://github.com/golang/build), using a custom framework to model the different workflow steps leading to release artifacts
* if you are into the classic web, check out the
  [tildeverse](https://tildeverse.org/), e.g.
[https://tilde.club](https://tilde.club),
[https://tilde.town/](https://tilde.town/), ...

