---
title: "CLT23"
date: 2023-01-11T14:00:00+02:00
draft: false
---

<!-- Note to self: change date to 2023-01-11 to 2023-03-11 before talk -->

## Hello CLT23!

This is a hop between the go tool requesting **golangleipzig.space/clt23** and
the repository hosting the actual code.


    $ go install golangleipzig.space/clt23@latest
    $ clt23

If you want to bypass the default module proxy, you can use:

    $ GOPRIVATE="*" go install golangleipzig.space/clt23@latest

If you take a look a the [source of this
page](view-source:https://golangleipzig.space/clt23/), you'll discover the
*go-import* meta tag.

The go install flow, visualized:

![](/images/go-install-flow.png)
