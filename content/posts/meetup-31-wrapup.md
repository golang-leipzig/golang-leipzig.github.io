---
title: "Meetup 31 Wrapup"
date: 2022-10-22T13:00:31+02:00
draft: false
tags:
- summary
- meetup
---

## Personal Cloud

On [2022-10-18 19:00 CEST](meetup) we had our #31 meetup, this time [Basislager Coworking][basislager] and with a hybrid setup (that still needs some improvement, but we're getting there).

[Max Eusterbrock][max] and his colleague [Aavash Shrestha](aavash) told us about [_The personal cloud and the monster that makes it_][slides].  In this presentation they gave us a preview on one of their new platforms that provides one-click hosting of applications, which they demonstrated with the deployment of a [note taking service][knotro].  One of the great things about this new platform was that anyone is able to install services, without requiring preliminary knowledge on cloud services, and that the client also owns the data that is produced by the installed service.  This is in contrast to common cloud services where you just create an account and all your data is owned and stored by whoever runs the service.  Aavash gave us a deep dive on the architecture of the new platform, explaining how they isolate workloads, and also how they designed it to easily scale with increased load.

## Misc

One of our meetup attendees asked about creating PDF files in Go, which is a potential topic for a future meetup, but because of the lack of experience with this domain we could just recommend an article that was featured in a recent [golangweekly][goweekly] issue called [How To Create a PDF in Go: A Step-By-Step Tutorial][gopdf] ([medium](https://medium.com/the-godev-corner/how-to-create-a-pdf-in-go-157355429a94)).

Also, someone stumbled upon a very common [Go gotcha][gotchas] which is about [value references in range loops][gofaq].  There was a [Go language proposal by Russ Cox][proposal], just a couple of days ago, to implicitly redefine loop variables.  Assume we have the following code

```go
xs := []*string{}
for _, v := range []string{"a", "b", "c"} {
    xs = append(xs, &v)
}
```

what do you think will be the value of xs after the loop?  It's not what most people would expect, that is `["a", "b", "c"]`.  To get the expected output you need to redefine the loop variable

```go
xs := []*string{}
for _, v := range []string{"a", "b", "c"} {
    v := v
    xs = append(xs, &v)
}
```

which is not obvious ([playground](https://go.dev/play/p/OpE2m09Q-OP)).

----

[Join our meetup](https://www.meetup.com/Leipzig-Golang) to get notified of
upcoming events!



[slides]: /downloads/deta-leipzig-gophers-31-slides.pdf
[meetup]: https://www.meetup.com/leipzig-golang/events/282941951/
[basislager]: https://www.basislager.co/
[deta]: https://www.deta.sh/
[max]: https://de.linkedin.com/in/xeust
[aavash]: https://de.linkedin.com/in/aavshr
[knotro]: https://knotro.com/getting-started/deploy-your-app
[goweekly]: https://golangweekly.com
[gopdf]: https://scribe.rip/the-godev-corner/how-to-create-a-pdf-in-go-157355429a94
[gotchas]: https://github.com/golang-leipzig/gotchas
[proposal]: https://github.com/golang/go/discussions/56010
[gofaq]: https://go.dev/doc/faq#closures_and_goroutines
