---
title: "Hybrid Meetup #41 wrap-up"
date: 2024-02-28T10:00:00+01:00
draft: true
tags:
- summary
- meetup
---

## Go 1.22 and more

Hybrid Meetup #41 took place
[2024-02-27](https://www.meetup.com/leipzig-golang/events/298066352) at
[Basislager Leipzig](https://www.basislager.co/). We looked at some updates in
[Go 1.22](https://golang.org/doc/go1.22), especially the sharing bug fix in
for [loops](https://go.dev/wiki/LoopvarExperiment) and the enhanced routing pattern.

## Go tools for LLM

There is no shortage of tooling and applications for LLM and generative AI in
general, and increasingly, these tools abstract the model access away to you
can use remote or local models with the same code.

We also briefly looked at [lingoose](https://lingoose.io/), a lightweight LLM
framework, which offers some wrappers around common tasks for generative AI
models, like chat or document adapters for retrieval-augmented generation.

Notes: [What RAG?](https://github.com/miku/whatrag)

## The Cutoff

For the record, 2023 may be the cutoff year. The amount of syntheticly
generated material will dwarf everything that existed up to that point.

## Misc

* performance, arm, x86, cf. dwstalk
* what does an LLM replace in terms of programming? how much can and will be automated?
* ...

## Joke

COBOL is still used, and still may be in the future. Maybe someone can write a
Go-to-COBOL tool, we already would have a name: gobol (this name has also been
suggested by
[gemma](https://huggingface.co/docs/transformers/en/model_doc/gemma) 2B (91bff873f359), when
prompted *how would you name a software project that translates Go (golang) to
COBOL? offer multiple alternatives*)
