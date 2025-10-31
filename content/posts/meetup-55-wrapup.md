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
[Basislager Leipzig](https://basislager.co) and we had a great presentation by
[Klaas Freitag](https://www.linkedin.com/in/klaasf/) (CTO) and Principal Architect
[Dr. Jörn F. Dreyer](https://www.xing.com/profile/Joern_Dreyer) from
[OpenCloud](https://opencloud.eu). A recording can be found [here](https://youtu.be/DChn7mZuiIA):

![](/images/screenshot-2025-10-30-134510-leipzig-gophers-55-opencloud-youtube.png)

OpenCloud is a widely deployed cloud storage and collaboration platform built
on a variation of a microservices architecture. It scales from homelab
installations to large clusters with millions of users.

The presentation reflected on some architectural and deployment changes over
the years - densily packed with engineering wisdom that extends beyond code and
include aspects like deployment, backwards compatibilty and scalability.

<!--

[![](/meetup-55-opencloud/screenshot-2025-10-31-112307-opencloud-landscape.png)](https://docs.opencloud.eu/)

-->

Some highlights from the presentation:

* [reva](https://reva.link/) is a CERN storage interop layer and is also where opencloud story started

> Reva is an interoperability platform consisting of several daemons written in
> Go. It acts as bridge between high-level clients (mobile, web, desktop) and
> the underlying storage (CephFS,
> [EOS](https://github.com/cern-eos/eos), local
> filesytems). It exports well-known APIs, like WebDAV, to faciliate access
> from these devices. It also exports a high-performance gRPC API, codenamed
> CS3 APIs, to easily integrate with other systems. Reva is meant to be a high
> performant and customizable HTTP and gRPC server. --
> [github.com/cs3org/reva/](https://github.com/cs3org/reva/)

[EOS](https://eos-docs.web.cern.ch/diopside/) itself is an impressive storage system:

> EOS instances at CERN store more than seven billion files and provide 780
> petabytes of disk storage capacity using over 60k hard drives (as of June
> 2022), matching the exceptional performance of the LHC machine and
> experiments.

[CERNBox](https://cern.service-now.com/service-portal?id=service_element&name=CERNBox-Service)
acts as an file sync and service layer over EOS and is based on ownCloud (from
which opencloud [was
forked](https://www.linux-magazin.de/artikel/opencloud-forkt-owncloud-neue-wendung-bei-den-freien-speichercloud-versionen/)).

> CERNBox is a cloud storage and file synchronization service developed at
> CERN, built on the open-source software ownCloud and EOS. It enables users to
> securely store, access, and share files from any device. It offers 1TB of
> personal space (just login to cernbox.cern.ch) and 1-100TB for (justified)
> project space.

More background on CERNBox: [Turning CephFS into a collaborative space with
CERNBox](https://www.epj-conferences.org/articles/epjconf/pdf/2025/22/epjconf_chep2025_01041.pdf)
(2025)

* not uncontroversial: you can get rid of a database at the core of your application (which was, in parts, a bottleneck) and move to a file based setup (plus lots of caches)
* moving from individual shares to the concept of spaces opened up a more maintainably way to handle users (and users that left)
* moving from from individual microservices to a more monolithic microservice architecture has been beneficial; internally opencloud uses [nats](https://nats.io/) for messaging (cf. list of microservices in the docs: [section services](https://docs.opencloud.eu/docs/dev/server/))
* large scale deployments with predictable, but still spiky patterns inspired changes to the node communication setup
* while user report that opencloud feels fast, it is hard to attribute this to the move from PHP to Go, only
* the layer between a (distributed) filesystem or object store and the end user
  view has is developed by an active community, which in parts is organized
  under the [CS3](https://www.cs3community.org/) umbrella


Thanks again to [Klaas](https://www.linkedin.com/in/klaasf/) and
[Jörn](https://www.xing.com/profile/Joern_Dreyer) for the inspiring
presentation.

----

[Join our meetup](https://www.meetup.com/de-DE/leipzig-golang/) to get notified of upcoming events.
