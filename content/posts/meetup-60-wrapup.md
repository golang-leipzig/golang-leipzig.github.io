---
title: "Hybrid Meetup #60 wrap-up"
date: 2026-05-28T12:00:00+02:00
draft: false
tags:
- summary
- meetup
---

## Nix, Go and learning with LLMs

Hybrid meetup #60 took place on Tuesday [May 26, 2026 19:00
CET](https://www.meetup.com/leipzig-golang/events/312537727) and we had an
great presentation by [Maxime](https://www.linkedin.com/in/plumps/) from [mxmlabs](https://mxmlabs.de/) on Nix and
LLM assisted active learning, the slides can be found [here](/downloads/lgo-60-nix-llm-slides.pdf) and we have [a recording, too](https://youtu.be/vMZeDe8wJu8):

[![](/images/lgo-60-nix-llm-slides-box.png)](/downloads/lgo-60-nix-llm-slides.pdf)

Some confusion around Nix is that it can mean a couple of different things: [a
language](https://nix.dev/tutorials/nix-language.html), [a package
manager](https://en.wikipedia.org/wiki/Nix_(package_manager)) and an [operating
system](https://en.wikipedia.org/wiki/NixOS).

The main concept of nix is the [derivation](https://nix.dev/manual/nix/2.25/language/derivations):

> a specification for running an executable on precisely defined input files to
> repeatably produce output files at uniquely determined file system paths

A great property of Nix is that it allows for gradual adoption. A downside of a
core abstraction in Nix, the [store](https://nix.dev/manual/nix/2.25/store/),
is that when run locally, it can consume more space than traditional
approaches, due to the strict isolation and immutability.

The learning with an LLM angle was illustrated on Nix examples, but the ideas are general:

* [Translate](http://localhost:1313/downloads/lgo-60-nix-llm-slides.pdf#page=13): style transfer something new to something known, e.g. nix language to Pseudocode or similar
* Evaluation (or execution) [Tracing](http://localhost:1313/downloads/lgo-60-nix-llm-slides.pdf#page=14) (for simple cases)
* [Socratic debugging](http://localhost:1313/downloads/lgo-60-nix-llm-slides.pdf#page=15), using the [socratic method](https://en.wikipedia.org/wiki/Socratic_method) (below an early computer based socratic tutoring system, called "WHY", [from 1977](https://files.eric.ed.gov/fulltext/ED138297.pdf))

[![](/meetup-60/screenshot-2026-05-28-222132-dialogue-with-the-system.png)](https://files.eric.ed.gov/fulltext/ED138297.pdf#page=9)

* [Three Variations](http://localhost:1313/downloads/lgo-60-nix-llm-slides.pdf#page=16); learn by comparing solutions across competency levels
  (below screencast shows an
[DFS](https://en.wikipedia.org/wiki/Depth-first_search) implementation with
mostly the same prompt, just varying style (**beginner**, **expert**, **compact**), [gemma4](https://ai.google.dev/gemma/docs/core/model_card_4) via [ollama](https://ollama.com/library/gemma4:latest) (8B, Q4KM) LLM
running an [FWD](https://frame.work/de/en/desktop), AMD Strix Halo 128GB):


[![](/meetup-60/synopsis/synopsis-recording-20260529-123126-github-light-O3.gif)](/meetup-60/synopsis/synopsis-recording-20260529-123126-github-light.gif)

## Misc

* The nix package repository is one of the most comprehesive ones, see: [Repository statistics](https://repology.org/repositories/statistics/newest), with 142992 packages, [AUR](https://repology.org/repository/aur) has 112992 (2026-05-28).
* Nix based tools that help setting update development environments: [flox](https://flox.dev/), [devenv](https://devenv.sh/), ...
* The past decade has seen mostly container based attempts at reproducibility, with [some caveats](https://arxiv.org/pdf/2601.12811):

[![](/images/screenshot-2026-05-28-144639-lgo-60-reproducibility-docker.png)](https://arxiv.org/pdf/2601.12811#page=19)

> Our results show that Docker does not guarantee reproducibility under any
> tested definition, nor is there a "silver bullet" set of rules for writing
> Dockerfiles yielding reproducible images. -- [Docker Does Not Guarantee Reproducibility](https://arxiv.org/pdf/2601.12811#page=25) (01/2026)

* [ostree](https://en.wikipedia.org/wiki/OSTree) spawned a set of distributions
  that build on the git-like model for bootable filesystem trees, and also tools like [flatpak](https://docs.flatpak.org/en/latest/under-the-hood.html)
* [clan.lol](https://clan.lol/), a declarative framework for managing fleets of machines without a central controller
* some ([personal](https://github.com/miku)) examples for "style-transferred" or translated programs: [goliza](https://github.com/miku/goliza), an ELIZA bot rewritten in Go, [animcan](https://github.com/miku/animcan), terminal animations inspired by [firew0rks](https://github.com/addyosmani/firew0rks) (js), ...

## Thanks!

Thanks again [Maxime](https://www.linkedin.com/in/plumps/) giving us a gimpse into a better way to manage the software stack.

----

[Join our meetup](https://www.meetup.com/de-DE/leipzig-golang/) to get notified of upcoming events.
