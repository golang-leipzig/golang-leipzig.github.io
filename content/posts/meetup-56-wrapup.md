---
title: "Hybrid Meetup #56 wrap-up"
date: 2025-11-26T00:30:00+01:00
draft: true
tags:
- summary
- meetup
---

### An older version of the matrix

Hybrid Meetup #56 took place
[2025-11-25](https://www.meetup.com/de-de/leipzig-golang/events/305626275/) 19:00 at
[Basislager Leipzig](https://basislager.co) and we looked into basic agents with Go.

Agents are possible because of the reasoning and tool support of
language models.

An early paper on tools was [*Toolformer: Language Models Can Teach Themselves to Use Tools*](https://arxiv.org/pdf/2302.04761) (2023-02-09)

> We introduce Toolformer, a model trained to decide which APIs to call, when
> to call them, what arguments to pass, and how to best incorporate the results
> into future token prediction. This is done in a self-supervised way,
> requiring nothing more than a handful of demonstrations for each API. We
> incorporate a range of tools, including a calculator, a Q&A system, a search
> engine, a translation system, and a calendar.

[...]

> Given just a handful of human-written examples of how an API can be used, we
> let a LM annotate a huge language modeling dataset with potential API calls.
> We then use a self-supervised loss to determine which of these API calls
> actually help the model in predicting future tokens. Finally, we finetune the
> LM itself on the API calls that it con- siders useful. As illustrated in
> Figure 1, through this simple approach, LMs can learn to control a va- riety
> of tools, and to choose for themselves which tool to use when and how.

### Misc

* [go4lage gemini CV](https://go4lage.com/geminicv), for *escaping vendor lock-in*


----

[Join our meetup](https://www.meetup.com/de-DE/leipzig-golang/) to get notified of upcoming events.
