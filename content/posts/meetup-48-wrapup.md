---
title: "Hybrid Meetup #48 wrap-up"
date: 2025-03-10T08:00:00+01:00
draft: false
tags:
- summary
- meetup
---

## Kubernetes Security and checkpoint/restore in userspace

Hybrid Meetup #48 took place
[2025-02-25](https://www.meetup.com/leipzig-golang/events/305626240/) 19:00 at
[Basislager](https://www.basislager.co/) and we had a great presentation about
Kubernetes security.

### StackRox/RHACS

StackRox (founded in
[2014](http://web.archive.org/web/20151002074324/https://www.stackrox.com/))
focussed on container/cluster security and was acquired by Red Hat in
[2021](https://www.redhat.com/en/blog/red-hat-releases-open-source-stackrox-to-the-community).

> Today, Red Hat is excited to announce that Red Hat Advanced Cluster Security
> for Kubernetes (RHACS) is now open sourced as StackRox. The Kubernetes and
> container security community can now use and contribute to the codebase of
> [StackRox on Github](https://github.com/stackrox/stackrox).

There are three security layers on the cluster:

* build time (CVE handling, image checks, ...); supported by [roxctl](https://docs.redhat.com/en/documentation/red_hat_advanced_cluster_security_for_kubernetes/4.0/html-single/roxctl_cli/index#check-policy-compliance_cli-getting-started)
* deploy time ([admission controller](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/))
* runtime (agent collecting telemetry from nodes)

Interestingly, the core application can run on a single node, using a single
Postgres instance (up to 300GB); vertically scaled to (in the order of) 32
cores and 64GB RAM.  While this can be a bottleneck, clusters up to 3000 nodes and
40,000 cores are currently well supported by the application.

Some compute intensive parts of the application include the database, spikes
from user queries, long running queries or analytics.

StackRox will use other tools, such as [falco](https://falco.org/)
([source](https://github.com/falcosecurity/falco)) for event [monitoring](https://falco.org/docs/#what-does-falco-check-for).

> At its core, Falco is a kernel monitoring and detection agent that observes
> events, such as syscalls, based on custom rules. Falco can enhance these
> events by integrating metadata from the container runtime and Kubernetes. The
> collected events can be analyzed off-host in SIEM or data lake systems.

For node communication, StackRox will use
[mTLS](https://en.wikipedia.org/wiki/Mutual_authentication#mTLS) and
[gRPC](https://grpc.io/).

### Checkpoint/Restore in userspace

Creating snapshots from containers can be useful for forensic analysis or container migration.
We briefly looked at [criu](https://criu.org/Main_Page):

> It can freeze a running container (or an individual application) and
> checkpoint its state to disk. The data saved can be used to restore the
> application and run it exactly as it was during the time of the freeze. Using
> this functionality, application or container live migration, snapshots,
> remote debugging, and many other things are now possible.

More on that topic:

* [Forensic Analysis of Container Checkpoints - DevConf.CZ 2023](https://www.youtube.com/watch?v=pySOkAqlGtY)
* [Forensic container checkpointing and analysis](https://www.youtube.com/watch?v=hpoWOc8QAzU) (ASG23)

Thanks again to
[Simon](https://www.linkedin.com/in/simon-b%C3%A4umer-a61042177/) for the great
high-level archtectural overview.
