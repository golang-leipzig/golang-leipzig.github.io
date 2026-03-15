---
title: "Hybrid Meetup #57 wrap-up"
date: 2026-01-30T08:30:00+01:00
draft: false
tags:
- summary
- meetup
---

### Downhill from here

Hybrid Meetup #57 took place
[2026-01-27](https://www.meetup.com/de-de/leipzig-golang/events/312537678)
19:00 at [Basislager Leipzig](https://basislager.co) and we had a great
presentation by [Karl](https://www.linkedin.com/in/karl-breuer-4b71a8177/),
mixing mathematical optimization, Go and webassembly in his open source [OnlyOffice NLP
Solver](https://github.com/Karl1b/only-office-nlp-solver) plugin:

[![](/downloads/lgo-57-only-office-nlp-solver-slides-01-frame.png)](/downloads/lgo-57-only-office-nlp-solver-slides.pdf)

The project can be found on [GitHub](https://github.com/Karl1b/only-office-nlp-solver).

[![](/images/screenshot-2026-03-15-135110-karl1b-only-office-nlp-solver.png)](https://github.com/Karl1b/only-office-nlp-solver)

Some bits from the talk:

* nonlinear optimization can be made approachable to a larger user base by embedding it into a common (data) environment, like a spreadsheet application
* the [plugin](https://github.com/Karl1b/only-office-nlp-solver) handles constrained parameters and side conditions
* parsing arbitrary arithmetic expressions requires nesting and recursion, handled well by [context-free grammars](https://en.wikipedia.org/wiki/Context-free_grammar)
* wiring WASM to the onlyoffice spreadsheet api is another step to consider
* for linear optimization, you can use the simpler [simplex](/images/simplex-9783642615788.png) (one of the [top ten algorithms](https://pi.math.cornell.edu/~ajt/presentations/TopTenAlgorithms.pdf) of the 20th century)


An example nonlinear optimization problem could be optimal warehouse placement
(FACILITY LOCATION PROBLEM): Given a number of customers with fixed locations
and different demands, where should a warehouse be optimally placed to minimize
the total transport distances (if you are a delivery service and you want to
open a [dark store](https://en.wikipedia.org/wiki/Dark_store), you may want to
use this).

In the visualization, blue dots represent customers (their size is proportional
to their demand) and the red dot is the optimal location, as discovered by the
Nelder-Mead algorithm; more specific algorithms
[exist](https://en.wikipedia.org/wiki/Geometric_median#Computation) for this
problem).

![](/code/lgo-57-opt/lgo-57-wrapup-optimization.gif)

## Misc

* Will open source survive AI? When it becomes cheaper to generate code than to
  search for an open source solution, that may or may not cover all the
  requirements? A recent personal example is [gh-repos](https://github.com/miku/gh-repos).
* Go has [Web Assembly](https://go.dev/wiki/WebAssembly) support since [1.11](https://go.dev/doc/go1.11#wasm), released 2018-08-24


Thanks again [Karl](https://karlbreuer.com/) for the nice presentation and project!

----

[Join our meetup](https://www.meetup.com/de-DE/leipzig-golang/) to get notified of upcoming events.
