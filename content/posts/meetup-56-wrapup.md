---
title: "Hybrid Meetup #56 wrap-up"
date: 2025-11-26T00:30:00+01:00
draft: false
tags:
- summary
- meetup
---

### An older version of the matrix · 不気味の谷

Hybrid Meetup #56 took place
[2025-11-25](https://www.meetup.com/de-de/leipzig-golang/events/305626275/)
19:00 at [Basislager Leipzig](https://basislager.co) and we looked into basic
agents with Go, notes can be found here:
[miku/unplugged](https://github.com/miku/unplugged).

Agents are possible because of the *reasoning* and tool support of language
models (and they are [simple to
write](https://fly.io/blog/everyone-write-an-agent/).

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
> actually help the model in predicting future tokens.

An agentic setup then is mostly a loop that manages a context over time with
the help of tools.

![](/images/lgo-56-looptool.gif)

[Google ADK for Go](https://github.com/google/adk-go) at this time only
supports gemini out of the box
([#233](https://github.com/google/adk-go/pull/233),
[#242](https://github.com/google/adk-go/pull/242), ...), so we wrote a simple
agent from scratch (against any openai compatible endpoint) and ended up with
an program that had a short list of tools (some of them just stubs):

```
get_weather
add_numbers
get_time
search_library_catalog
ping
list_files
read_file
grep
write_file
append_file
run_command
```

We used both an [RTX 4000 SFF
ADA](https://www.nvidia.com/content/dam/en-zz/Solutions/rtx-4000-sff/proviz-rtx-4000-sff-ada-datasheet-2616456-web.pdf)
and an [AMD AI MAX+ 395 with an
8060S](https://www.amd.com/en/products/processors/laptop/ryzen/ai-300-series/amd-ryzen-ai-max-plus-395.html) (via [FWD](https://frame.work/de/en/desktop)).

| Spec | AMD Radeon 8060S | NVIDIA RTX 4000 SFF Ada |
|------|------------------|-------------------------|
| FP16 (theoretical) | 59.4 TFLOPS | ~19.2 TFLOPS |
| Memory Bandwidth | ~212 GB/s (DDR5-8000) | 280 GB/s (GDDR6) |

However, prefill is a bit faster on the nvidia card:

```
$ time OLLAMA_MODEL=qwen3:14b OLLAMA_HOST=http://ada:11434 ./one -m "how warm is it in leipzig?"
2025/11/26 09:57:17 user: how warm is it in leipzig?               ...
2025/11/26 09:57:17 context length: 4974                           ...
2025/11/26 09:57:23 assistant wants to call 1 tool(s)              ...
2025/11/26 09:57:23 calling tool: get_weather                      ...
2025/11/26 09:57:23 args: {"city":"Leipzig"}                       ...
2025/11/26 09:57:23     Result: {"city": "Leipzig", "temperature": ...
2025/11/26 09:57:23 context length: 5227                           ...
2025/11/26 09:57:30 assistant: The current temperature in Leipzig i...

real    0m13.423s
user    0m0.001s
sys     0m0.013s

$ time OLLAMA_MODEL=qwen3:14b OLLAMA_HOST=http://strix:11434 ./one -m "how warm is it in leipzig?"
2025/11/26 09:57:34 user: how warm is it in leipzig?             ...
2025/11/26 09:57:34 context length: 4974
2025/11/26 09:57:41 assistant wants to call 1 tool(s)
2025/11/26 09:57:41 calling tool: get_weather
2025/11/26 09:57:41 args: {"city":"Leipzig"}
2025/11/26 09:57:41     Result: {"city": "Leipzig", "temperature"...
2025/11/26 09:57:41 context length: 5227
2025/11/26 09:57:50 assistant: The current temperature in Leipzig...

real    0m15.826s
user    0m0.007s
sys     0m0.007s
```

Still, the interplay is interesting to observe. Requests like *"save the
temperature in leipzig to temp.txt"* or *"fetch
[https://golangleipzig.space/leipzig-gopher.png](https://golangleipzig.space/leipzig-gopher.png)
and convert it to jpg"* work with a 9.3GB 14B tool supporting LLM like
[Qwen3-14B](https://huggingface.co/Qwen/Qwen3-14B).

> Expertise in agent capabilities, enabling precise integration with external
> tools in both thinking and unthinking modes and achieving leading performance
> among open-source models in complex agent-based tasks. -- [model card](https://huggingface.co/Qwen/Qwen3-14B)

Example: (1) fetch image, (2) convert to jpg, (3) calculate sha1 and (4) write the result to a file (speedup 1.5x):

![](/images/lgo-56-agent-3.gif)

Future ideas for tools:

* [ ] browser use to facility web search
* [ ] manage context by offloading to text files, similar to deep research agents
* [ ] code snippets for subtasks and sandboxed execution

Go's concurrency facilities seems to be helpful when implementing agents.

### Misc

* There is no shortage of agent frameworks, e.g.
  [tRPC-Agent-Go](https://github.com/trpc-group/trpc-agent-go),
[eino](https://github.com/cloudwego/eino),
[genkit](https://genkit.dev/docs/get-started/?lang=go),
[swarmgo](https://github.com/prathyushnallamothu/swarmgo),
[go-agent](https://github.com/vitalii-honchar/go-agent),
[suricata](https://github.com/ostafen/suricata),
[agent-sdk-go](https://github.com/Ingenimax/agent-sdk-go), and many more in Go
or [any language](https://www.shakudo.io/blog/top-9-ai-agent-frameworks)
* [KI BARCAMP HALLE (SAALE) 2025-11-29](https://www.klaustor-coworking.de/events-1/ki-barcamp-2025)
* [go4lage gemini CV](https://go4lage.com/geminicv), for *escaping vendor lock-in*
* [Can Google's ADK Replace LangChain and MCP? (with Christina Lin)](https://www.youtube.com/watch?v=nMnQ63YkftE)
* AI will be similar to the internet, e.g. in terms of omnipresence
* AI [uncanny valley](https://en.wikipedia.org/wiki/Uncanny_valley), 不気味の谷

Did you implement a cool agent in Go (or something else)? Then why not [join
us](https://www.meetup.com/de-DE/leipzig-golang/) and let it introduce itself.

----

[Join our meetup](https://www.meetup.com/de-DE/leipzig-golang/) to get notified of upcoming events.
