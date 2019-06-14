---
title: "Meetup #5 wrap-up"
date: 2019-06-15T01:00:00+02:00
draft: false
tags:
- summary
- meetup
---

{{% h2 %}}Bits from the IO package{{% /h2 %}}

The [IO package](https://golang.org/pkg/io/) is a central package in the
standard library, as it provides (among other things) two main interfaces:

* [io.Reader](https://golang.org/pkg/io/#Reader)
* [io.Writer](https://golang.org/pkg/io/#Writer)

In a lightning talk we digged a bit into the package:

* [Slides](https://github.com/miku/io15min/blob/master/Slides.md) ([PDF](https://github.com/miku/io15min/blob/master/Slides.pdf))
* [Repo](https://github.com/miku/io15min)

The IO model is simple and powerful (and reminds one of UNIX pipes).


{{% h2 %}}Automate the Web with chromedp{{% /h2 %}}

In a code walkthrough we explored [chromedp](https://github.com/chromedp/chromedp), a pure Go library
talking the Chrome [devtools protocol](https://github.com/ChromeDevTools/devtools-protocol). It allows to run a headless
browser. We looked at three examples:

* Generate good looking full page screenshots
* Generate PDF invoices from HTML
* Rendering SVG

The whole setup can be made robust by isolating Chrome in a container. The
invoice example was a nice one, as working with HTML allows anyone with
HTML/CSS skills to adjust the PDF output (which is usually harder to generate
programmatically). Running many (even headless) Chrome instances will eat your
RAM.


{{% h2 %}}How do developers discover Go?{{% /h2 %}}

Have you ever wondered about how programmers move from language to language?
The author of this entertaining and enlightning blog post (and [other
things](https://github.com/sshuttle/sshuttle)) did as well:

* [Programmer migration patterns](https://apenwarr.ca/log/20190318)



{{% h2 %}}More Linux and Cloud{{% /h2 %}}

We have a presentation about libpod (a tool to work with OCI images) in the
pipeline. Meanwhile, do not miss the next [Linux
Meetup](https://www.meetup.com/de-DE/Linux-Meetup-Leipzig/events/261912346/) on
Tue, 2019-06-18 about runc and CRI-O.




{{% h2 %}}References{{% /h2 %}}

* [IO lightning talk repo](https://github.com/miku/io15min/)
* [chromedp](https://github.com/chromedp/chromedp)
* [Linux Meetup #6](https://www.meetup.com/de-DE/Linux-Meetup-Leipzig/events/261912346/)
