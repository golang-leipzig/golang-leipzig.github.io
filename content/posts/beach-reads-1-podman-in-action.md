---
title: "Beach Reads #1"
date: 2023-01-29T16:00:00+02:00
draft: true
tags:
- review
- book
- beachreads
---

## Beach Reads #1: Podman in Action

> Walsh, Daniel. [Podman in Action](https://www.manning.com/books/podman-in-action). [Manning Publications](https://www.manning.com), 2023.

[Podman](https://www.redhat.com/en/topics/containers/what-is-podman) is an
alternative to Docker to manage the container life cycle. As of 02/2023 podman
contains 167K SLOC of Go code spread across 1149. The bulk of the functionality
comes from a set of libraries found under the [Container
Tools](https://github.com/containers) umbrella, e.g. like storage, image
handling or container monitoring. Combining various modular tools into a single
project and binary allow podman to be a drop-in replacement for docker.

Podman works without a daemon.

> The core Podman engineering team come from an operating system background
> more grounded in the Unix Philosophy.

This means that a container process started with podman will not descend of a
central daemon process. However, it still needs some system resources made
available ([details](https://unix.stackexchange.com/a/534934/376)), so it will
hand over process management to systemd, which itself uses
[conmon](https://github.com/containers/conmon/) to implement communication
between the OCI runtime and the container manager.

```shell
$ podman run alpine sleep 600

...

$ pstree -s -p 45992
systemd(1)───systemd(3324)───conmon(45988)───sleep(45992)

$ pstree -s -p 45943
systemd(1)───systemd(3324)───tmux: server(5608)───bash(7191)───podman(45943)─┬─slirp4netns(45985)
                                                                             ├─{podman}(45944)
                                                                             ├─...
                                                                             └─{podman}(45983)
```

There is nothing constraining to run a separate server component on a system, in fact
it may interfere with existing tools when it comes to process management.

> When Podman was designed, the developers wanted to make sure it fully integrated with systemd.


