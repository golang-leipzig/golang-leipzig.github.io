---
title: "Virtual Meetup #12 wrap-up"
date: 2020-07-18T10:00:00+02:00
draft: false
tags:
- summary
- meetup
---

We had a short talk about HTTP proxies in Go. As expected, Go supports the
usual environment variables, like `HTTP_PROXY`, `HTTPS_PROXY` and `NO_PROXY` -
and you can easily customize the proxy selection, as we saw in an example code
walkthrough of an ip rotation proxy.

The [x/net](https://github.com/golang/net/) package contains support in
[x/net/http/httpproxy](https://github.com/golang/net/tree/master/http/httpproxy).

A couple of fun facts from the talk:

* The first book on HTTP proxies was published in 1998: "Web Proxy Servers" by
  Ari Luotonen (CERN), who also wrote a [web proxy tunneling
draft](https://tools.ietf.org/html/draft-luotonen-web-proxy-tunneling-01).
* Many popular HTTP projects (including Go) were affected by a [critical
  CVE](https://httpoxy.org/), related to [CGI](https://www.w3.org/CGI/) up until
  2016. Note that the discussion and development of "server scripts" [started
        in
1993](http://1997.webhistory.org/www.lists/www-talk.1993q4/0485.html).
* Core curl developer [explains](https://stackoverflow.com/a/62722840/89391),
  why curl does not follow RFC in each and every aspect, e.g. when it comes to
the `Proxy-Connection` header - because "The world wild web is a crazy place."

See you next time!


