---
title: "CLT23"
date: 2023-01-11T14:00:00+02:00
draft: false
---

<!-- Note to self: change date to 2023-01-11 to 2023-03-11 before talk -->

## Hello [CLT23](https://chemnitzer.linux-tage.de/2023)!

This is a hop between the go tool requesting the latest version of
**golangleipzig.space/clt23** and the
[repository](https://github.com/golang-leipzig/clt23/) hosting the actual code (demo for [CLT2023/234](https://chemnitzer.linux-tage.de/2023/de/programm/beitrag/234)).


    $ go install golangleipzig.space/clt23@latest
    $ clt23

If you want to bypass the default module proxy (e.g. to circumvent caching in the proxy), you can use:

    $ GOPRIVATE="*" go install golangleipzig.space/clt23@latest

If you take a look a the [source of this
page](view-source:https://golangleipzig.space/clt23/), you'll discover the
*go-import* meta tag. The go install flow is visualized below.

[![](/images/go-install-flow.png)](https://github.com/golang-leipzig/clt23/)

