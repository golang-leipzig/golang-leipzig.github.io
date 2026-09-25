---
title: "Wrapup: Everything is Open Source"
date: 2026-06-16T22:00:00+02:00
draft: false
tags:
- summary
- meetup
---

## Everything is Open Source

Hello, 世界!

Joint Rust and Go Meetup took place
[2026-06-16](https://www.meetup.com/rust-modern-systems-programming-in-leipzig/events/313813937/)
at Basislager Leipzig and we explored binary analysis in an interactive
workshop.

Current frontier proprietary and open weights models can iterate autonomously
via tool calls through an analysis:

* format analysis, binary structure
* linked libraries, signatures
* build metadata and even
* source layout and more.

As an example, on macos, a single, lightweight first analysis used:

* Binary/Mach-O inspection: [file](https://www.darwinsys.com/file/), `otool`, `nm`, `strings`, `lipo`, `size`, `xxd`
* Code signing / security: `codesign`, `spctl`
* Comparison & hashing: `cmp`, `shasum`
* System introspection: `log`
* Scripting / shell: [python3](https://www.python.org/), [bash](https://www.gnu.org/software/bash/) (builtins incl. `printf`, `echo`)
* Text processing in pipelines: `grep`, [awk](https://www.gnu.org/software/gawk/manual/gawk.html#Getting-Started), `sed`, `head`, `tail`, `wc`, `tr`, `cat`, `ls`

> In summer 2026, all software starts to look more and more like open source.

[![](/images/everything-is-open-source.png)](/downloads/everything-is-open-source.pdf)


