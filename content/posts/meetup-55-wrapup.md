---
title: "Hybrid Meetup #55 wrap-up"
date: 2025-10-30T10:30:00+02:00
draft: true
tags:
- summary
- meetup
---

Hybrid Meetup #55 took place
[2025-10-28](https://www.meetup.com/de-de/leipzig-golang/events/305626267) 19:00 at
[Basislager Leipzig](https://basislager.co) and we had a great presentation from
[Klaas Freitag](https://www.linkedin.com/in/klaasf/) (CTO) and Principal Architect
[Dr. Jörn F. Dreyer](https://www.xing.com/profile/Joern_Dreyer) from
[OpenCloud](https://opencloud.eu). A recording can be found [here](https://youtu.be/DChn7mZuiIA):

![](screenshot-2025-10-30-134510-leipzig-gophers-55-opencloud-youtube.png)

OpenCloud is a widely deployed OpenCloud cloud storage and collaboration
platform built on a microservices architecture. It scales from homelab
installations to large clusters with millions of users.

Some highlights from the presentation:

* [reva](https://reva.link/) is a CERN storage interop layer and is also where opencloud story started

> Reva is an interoperability platform consisting of several daemons written in
> Go. It acts as bridge between high-level clients (mobile, web, desktop) and
> the underlying storage (CephFS,
> [EOS](https://eos-docs.web.cern.ch/diopside/introduction/index.html), local
> filesytems). It exports well-known APIs, like WebDAV, to faciliate access
> from these devices. It also exports a high-performance gRPC API, codenamed
> CS3 APIs, to easily integrate with other systems. Reva is meant to be a high
> performant and customizable HTTP and gRPC server. --
> [github.com/cs3org/reva/](https://github.com/cs3org/reva/)

EOS itself is an impressive storage system:

> EOS instances at CERN store more than seven billion files and provide 780
> petabytes of disk storage capacity using over 60k hard drives (as of June
> 2022), matching the exceptional performance of the LHC machine and
> experiments.

[CERNBox](https://cern.service-now.com/service-portal?id=service_element&name=CERNBox-Service) acts as an file sync and service layer over EOS and is based on
ownCloud (from which openCloud [was
forked](https://www.linux-magazin.de/artikel/opencloud-forkt-owncloud-neue-wendung-bei-den-freien-speichercloud-versionen/)).

> CERNBox is a cloud storage and file synchronization service developed at
> CERN, built on the open-source software ownCloud and EOS. It enables users to
> securely store, access, and share files from any device. It offers 1TB of
> personal space (just login to cernbox.cern.ch) and 1-100TB for (justified)
> project space.



----

[Join our meetup](https://www.meetup.com/de-DE/leipzig-golang/) to get notified of upcoming events.
