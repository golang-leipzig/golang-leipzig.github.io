---
title: "Summer Drinkup #53 wrap-up"
date: 2025-07-16T12:30:00+02:00
draft: false
tags:
- summary
- meetup
---

## Misc

* Math: [Long(er)-Form Mathematics](https://longformmath.com/), [Terrence Tao, Analysis I](https://www.google.com/search?q=terrence+tao+analysis), [Mathematics: Its Contents, Methods and Meaning](https://archive.org/details/MathematicsItsContentsMethodsAndMeaningVol3)
* Music: [Guitar Practice](https://www.captrice.io/)
* [CLRS](https://www.betterworldbooks.com/search/results?q=%22Introduction%20to%20Algorithms%22%20cormen), [Introduction to Algorithms](https://en.wikipedia.org/wiki/Introduction_to_Algorithms)
* regular languages, [regular expressions](https://stackoverflow.com/q/1732348/89391), chomsky hierarchy, metamathematics

[![](/images/640px-Chomsky-hierarchy-s.svg.png)](https://en.wikipedia.org/wiki/Chomsky_hierarchy)

[![](/images/6691437-M.jpg)](https://en.wikipedia.org/wiki/Stephen_Cole_Kleene)

* [haystack](https://haystack.deepset.ai/), [langchain](https://www.langchain.com/)
* one does not simply reindexing a whole [vector database](https://en.wikipedia.org/wiki/Vector_database)
* [HNSW](https://en.wikipedia.org/wiki/Hierarchical_navigable_small_world), [approximate nearest neighbor search](https://en.wikipedia.org/wiki/Nearest_neighbor_search#Approximation_methods) (classic [kNN](https://en.wikipedia.org/wiki/K-nearest_neighbors_algorithm)
is simple, robust and an algorithm that does not require any training)
* [42 school](https://www.42network.org/) -- "42 is the largest free IT school network in the world, with a [global presence](https://www.42network.org/42-schools/) of 50+ campuses across 30+ countries."
* running CI locally, e.g. with [GL runner](https://docs.gitlab.com/runner/install/), or [act](https://github.com/nektos/act)
* challenges of system designs that include language models or related technologies: levels or randomness (LLM parameters, temperature, ...)

> However, the influence of temperature on creativity is far more nuanced and
> weak than suggested by the "creativity parameter" claim; overall results
> suggest that the LLM generates slightly more novel outputs as temperatures
> get higher. -- [Is Temperature the Creativity Parameter of Large Language Models?](https://arxiv.org/pdf/2405.00492)

Randomness in embeddings: Do you get different embedding vectors from the same
model on the same input? Yes and no. [Quick
test](https://github.com/golang-leipzig/golang-leipzig.github.io/tree/source/static/vembedtest)
using [ollama](https://ollama.com) running on a Intel
[N3450](https://www.intel.com/content/www/us/en/products/sku/95596/intel-celeron-processor-n3450-2m-cache-up-to-2-20-ghz/specifications.html)
2016 6W TDP CPU (and [1.6W idle](/images/zima-idle-1.6w.png) power
consumption) reveals deterministic embeddings, whereas there are [observations of the opposite](https://github.com/golang-leipzig/golang-leipzig.github.io/tree/source/static/vembedtest) as well:

![](/vembedtest/vembedtest.gif)

Estimated cost of running this test (90 embeddings, total): at most 9W for
about 6 minutes at [0.30 EUR/kWh](/images/l-strom-dyn-2025-07.png): EUR 0.00027
(or framed differently: you can calculate about 640000 embeddings for the price
of one scoop of ice cream - albeit it would take about a month to calculate
them with this device) 🍦

* [Fabrice Bellard](https://en.wikipedia.org/wiki/Fabrice_Bellard),
  [superproductivity](http://web.archive.org/web/20121006002711/http://blog.smartbear.com/software-quality/bid/167059/Fabrice-Bellard-Portrait-of-a-Superproductive-Programmer),
plus: [jslinux](https://bellard.org/jslinux/), [NNCP](https://bellard.org/nncp/), [libNC](https://bellard.org/libnc/), ...

> If there’s a secret to this superhero-level productivity, it appears to have
> less to do with comic-book mutation and radioactivity, and far more with
> discipline, confidence, rigor, and many years of practice.

* another helpful thing: accountability

> The positive effect of accountability was supported: those who sent weekly
> progress reports to their friend accomplished significantly more than those
> who had unwritten goals, wrote their goals, formulated action commitments or
> sent those action commitments to a friend. -- [The Impact of Commitment, Accountability, and Written Goals on Goal AchievementGoal Achievement](https://scholar.dominican.edu/cgi/viewcontent.cgi?article=1002&context=psychology-faculty-conference-presentations)

----

[Join our meetup](https://www.meetup.com/de-DE/leipzig-golang/) to get notified of upcoming events.
