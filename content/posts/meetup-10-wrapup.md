---
title: "Virtual Meetup #10 wrap-up"
date: 2020-04-21T10:00:00+02:00
draft: false
tags:
- summary
- meetup
---

{{% h2 %}}The sync package{{% /h2 %}}

Virtual [Meetup #10](https://www.meetup.com/Leipzig-Golang/events/268785531/)
took place on Friday, April 17, 2020, 19:00 CEST via [Zoom](https://zoom.us/)
(thanks to [saschagrunert](https://github.com/saschagrunert) and
[CNCF](https://www.cncf.io/)).

[Michael](https://twitter.com/embano1) prepared a [great overview](https://github.com/embano1/go-meetup-lej-04-2020) of the
[sync](https://golang.org/pkg/sync/) and
[x/sync](https://pkg.go.dev/golang.org/x/sync) packages, which implement
concurrency related facilities (concurrency is hard with either classical
approaches or CSP, as we [learned from a
presentation](https://speakerdeck.com/embano1/concurrency-bugs-in-go-go-meetup-leipzig-03-15-2019-gasch)
last year at [meetup
#2](https://golangleipzig.space/posts/second-meetup-wrapup/)).

The overview included among other things a bug hunt in a
[counter](https://github.com/embano1/go-meetup-lej-04-2020/blob/master/sync/sync-rwmutex/counter.go)
example and a benchmark of the builtin [map](https://golang.org/ref/spec#Map_types)
and [sync.Map](https://golang.org/pkg/sync/#Map).

![](/images/sync-map-bench.png)

The [x/sync](https://pkg.go.dev/mod/golang.org/x/sync) contains the very useful
[errgroup](https://github.com/embano1/go-meetup-lej-04-2020/tree/master/x-sync/errgroup)
and
[singleflight](https://github.com/embano1/go-meetup-lej-04-2020/tree/master/x-sync/singleflight).

One thing that the [errgroup](https://godoc.org/golang.org/x/sync/errgroup)
will not provide for out of the box is the pickup of multiple error, should
more than one occur.

> Go calls the given function in a new goroutine. The **first call** to return a non-nil error cancels the group; its error will be returned by Wait.

And:

> Wait blocks until all function calls from the Go method have returned, then
> returns the **first non-nil error** (if any) from them.

An implementation of this feature can be found in [k8s apimachinery's](https://github.com/kubernetes/apimachinery) [error
utilities](https://github.com/kubernetes/apimachinery/blob/06deae5c9c2c030d771a467e086b6c791e8800dc/pkg/util/errors/errors.go#L231-L246).

The singleflight pattern can also be found in the perkeep utils [go4](https://github.com/go4org/go4) repository, in the [singleflight](https://github.com/go4org/go4/blob/f5505b9728ddf920bb673137648788c5ac99de1b/syncutil/singleflight/singleflight.go#L17-L19) package.

> Package singleflight provides a duplicate function call suppression mechanism.

The way to [test
singleflight](https://github.com/go4org/go4/blob/f5505b9728ddf920bb673137648788c5ac99de1b/syncutil/singleflight/singleflight_test.go#L55-L85)
can be interesting, too. One pools up a number of goroutines (by starting and
waiting a bit). Then, a single value sent on a channel serves as a starting
shot. At the same time, the atomic counter on calls ensure there has only been
a single function call.


<!-- {{% h2 %}}Misc{{% /h2 %}} -->

