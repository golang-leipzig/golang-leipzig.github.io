---
title: "Fast XML processing with Go"
date: 2024-03-22T22:00:01+01:00
draft: true
tags:
- go
- xml
- performance
---

Go supports XML processing in the standard library package
[xml](https://pkg.go.dev/encoding/xml). Most of the time, XML handling falls
into one of the following buckets:

* You want to read or write a **single** XML document? Then,
  [Marshal](https://pkg.go.dev/encoding/xml#Marshal) and
[Unmarshal](https://pkg.go.dev/encoding/xml#Unmarshal) are convenient. An
example for this may be an XML configuration file.
* You want to write a large number of XML documents in a stream? Then you most likely want an [Encoder](https://pkg.go.dev/encoding/xml#Encoder.Encode)
* You want to read a large number of elements into a list of values? Then you
