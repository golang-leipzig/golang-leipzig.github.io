---
title: "Hybrid Meetup #45 wrap-up"
date: 2024-11-12T10:00:00+01:00
draft: true
tags:
- summary
- meetup
---

## Testing in Go

Hybrid Meetup #45 took place
[2024-10-29](https://www.meetup.com/leipzig-golang/events/298481354/) 19:00 at
[CHECK24](https://check24.de) Leipzig Office at Altes Landratsamt and we were
thrilled to see people joining from Leipzig
[and](https://en.wikipedia.org/wiki/Saxony-Anhalt)
[beyond](https://en.wikipedia.org/wiki/Thuringia), both on site and online. We
had two great input presentations about open source projects in the testing
domain.

### Coverage metrics

[Fabian](https://www.linkedin.com/in/fabian-g%C3%A4rtner-913584141/) analyzed
the Go coverage tool [cmd/cover](https://pkg.go.dev/cmd/cover), listed some of
its shortcomings and demonstrated a tool to address them and to improve some
metrics: [gocoverplus](https://github.com/Fabianexe/gocoverageplus).

XXX: example of three OSS projects and the coverage difference

```shell
$ go install github.com/Fabianexe/gocoverageplus@latest
```



### Go scheduler overlay

