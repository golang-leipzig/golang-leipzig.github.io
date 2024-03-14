---
title: "Hybrid Meetup #41 wrap-up"
date: 2024-03-01T10:00:00+01:00
draft: true
tags:
- summary
- meetup
---

## Go 1.22 and more

Hybrid Meetup #41 took place
[2024-02-27](https://www.meetup.com/leipzig-golang/events/298066352) at
[Basislager Leipzig](https://www.basislager.co/). We reviewed some updates in
[Go 1.22](https://golang.org/doc/go1.22), especially the sharing bug fix in
for [loops](https://go.dev/wiki/LoopvarExperiment) and the enhanced routing pattern.

## Go tools for [genai](https://en.wikipedia.org/wiki/Generative_artificial_intelligence) models

There is no shortage of tooling and applications for LLM and generative AI in
general, and increasingly, these tools abstract the model access away to you
can use remote or local models with the same code (and easily switch between
different providers).

We also briefly looked at [lingoose](https://lingoose.io/), a lightweight LLM
framework, which offers some wrappers around common tasks for generative AI
models, like chat or document adapters for retrieval-augmented generation.

Rough notes: [What RAG?](https://github.com/miku/whatrag)

## The Cutoff

For the record, 2023 may be the cutoff year. The amount of synthetically
generated material will dwarf everything that existed up to that point.

## Misc

* we wondered about some performance implications regarding ARM servers (which
  are becoming available on various cloud providers(cf. EUR 0.0070/h CAX11 instance at
[Hetzner](https://www.hetzner.com/cloud/))
* there are also tools to [reclaim your
  cloud](https://www.zimaboard.com/blade/)! Specifically, the zimablade runs a
low power 4-core [CPU](https://ark.intel.com/content/www/us/en/ark/products/95594/intel-celeron-processor-j3455-2m-cache-up-to-2-3-ghz.html) and can be equipped with 16GB RAM, plus two SATA drives
(back-of-the-envelope costs breakdown would be $128 (SBC), $15 for one 128GB
SATA SSD, a total of $158 or EUR 145 or 20714h of the cheapest (and weaker)
cloud instance; excluding positions like maintenance, bandwidth and energy, the
amorization period of on-prem hardware is somewhere between six month and three
years).

![](/images/41-comp-jungle.png)


## Joke

COBOL is still used, and still may be in the future. Maybe someone can write a
Go-to-COBOL tool, we already would have a name: gobol (this name has also been
suggested by
[gemma](https://huggingface.co/docs/transformers/en/model_doc/gemma) 2B (91bff873f359), when
prompted *how would you name a software project that translates Go (golang) to
COBOL? offer multiple alternatives*)
