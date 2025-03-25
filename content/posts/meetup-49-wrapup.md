---
title: "Hybrid Meetup #49 wrap-up"
date: 2025-03-26T08:00:00+01:00
draft: true
tags:
- summary
- meetup
---

## Memory management in Go or *To GC or not to GC*

Hybrid Meetup #49 took place
[2025-03-25](https://www.meetup.com/leipzig-golang/events/305626247) 19:00 at
[CHECK24](https://www.check24.de/) Leipzig Office and we had a great
introduction to memory management with Go.

Slides and quiz can be found at [mentimeter.com](https://www.mentimeter.com/app/presentation/alog8a9xsqj6hwbyi6t32m5qzh295rue/view?question=9bym8yfwp7c4).

[![](/images/meetup-49-menti-screenie-2025-03-25-231344.png)](https://www.mentimeter.com/app/presentation/alog8a9xsqj6hwbyi6t32m5qzh295rue/view?question=9bym8yfwp7c4)

Some rules to program by:

* Stack allocation are better then heap allocation
* Number of heap allocations are more important than size of heap allocations
* Reduce the number of pointers in heap to reduce mark phase costs
* Be aware of the memory layout of you structs
* Make sure that every goroutine returns at some point (to avoid leaks)
* Use pprof to determine problems in you program

There is a great GC guide at: [doc/gc-guide](https://go.dev/doc/gc-guide), and
a few SO questions [tagged with Go and
GC](https://stackoverflow.com/questions/tagged/go%2bgarbage-collection?tab=Votes)

A tool for visualizing is structlayout.

```
$ go install honnef.co/go/tools/cmd/structlayout@latest
$ go install honnef.co/go/tools/cmd/structlayout-pretty@latest
$ go get github.com/ajstarks/svgo/structlayout-svg
```

You can then generate struct layout diagrams.

```
$ structlayout -json runtime slice | structlayout-svg -t "runtime.slice" > /tmp/struct.svg
$ structlayout -json bufio Reader | structlayout-svg -t "bufio.Reader" > /tmp/struct.svg
$ structlayout -json net/http Response | structlayout-svg -t "net/http.Response" > /tmp/struct.svg
```

![](/images/meetup-49-structlayout-combined.png)


### Misc

