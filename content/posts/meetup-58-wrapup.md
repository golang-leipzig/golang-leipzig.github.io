---
title: "Hybrid Meetup #58 wrap-up"
date: 2026-04-01T08:30:00+01:00
draft: false
tags:
- summary
- meetup
---

## Agentic Coding with Beads

Hybrid meetup #58 took place on Tuesday [Mar 31, 2026 19:00
CET](https://www.meetup.com/leipzig-golang/events/312537698/) and we had a
great presentation by [Akim Zmerli](https://www.linkedin.com/in/akimzmerli/) on
[beads](https://github.com/steveyegge/beads) and agentic coding in general.

[![](/downloads/lgo-58/spot-s.jpg)](/downloads/lgo-58/spot.jpg)

Coding with AI or agentic coding has both upsides and downsides:

While it can **fill in blanks** in documentation, **help debug** or can quickly
get you up to speed with **new libraries**, the models also happily fill in the
blanks in the requirements, orienting themselves on patterns learned from
publicly available, average, code (but [synthetic
data](https://arxiv.org/pdf/2601.06953) (02/2026) is becoming more popular as well). A language model
can be a learning tool, or [erode skills](https://arxiv.org/pdf/2601.20245) (02/2026):

> The erosion of conceptual understanding, code reading, and debugging skills
> that we measured among participants using AI assistance suggests that workers
> acquiring new skills should be mindful of their reliance on AI during the
> learning process.

In general, model may misunderstand and misunderstand in [strange
ways](https://arxiv.org/pdf/2506.21521) (06/2025):

> success on benchmarks only demonstrates potemkin understanding: the illusion
> of understanding driven by answers irreconcilable with how any human would
> interpret a concept [...] potemkins are ubiquitous across models, tasks, and
> domains

### Taming models

Building a harness around an LLM can be seen as a consequence of its
probabilistic character. Beads is an agent memory system, aiming to steer an
agent through complex tasks by building a local, graph-based issue tracker.

Beads allows to formulate different types of tasks:

* bug
* feature
* task (default)
* epic
* chore
* decision

You can manually describe those, or even let the agent decompose a more complex
problem into subtasks. Tasks can have detailed descriptions, estimates,
acceptance criteria and other common features. Beads is quite extensible and
can also be used programmatically to build additional layers.

[![](/images/lgo-58-beads-ops.png)](https://github.com/gastownhall/beads/tree/main/examples/library-usage)

### Dashboard for beads

Akim also demod [kitty-beads](https://github.com/AkimZmerli/kitty-beads), a web
interface for beads, grouping issue, supporting terminal access in the browser
and a kanban board.

### Highlighted Go projects

The talk highlighed a few Go projects, too:

* [dolt](https://github.com/dolthub/dolt), git for MySQL
* [charm](https://github.com/charmbracelet) universe (we talked about it in [#28](/posts/meetup-28-wrapup), [miku/glamline](https://github.com/miku/glamline))
* [micasa](https://github.com/micasa-dev/micasa), home maintenance in a TUI
* [gastown](https://github.com/gastownhall/gastown), multi-agent workspace manager, by the author of beads

Note that Steve Yegge layed out in depth a vision for a pure agentic software
world in [Welcome to Gas
Town](https://steve-yegge.medium.com/welcome-to-gas-town-4f25ee16dd04) (34 min read). The
setup reminds one of the [set of markdown
files](https://github.com/garrytan/gstack) that were framed as virtual
engineering team.

### AI and Go

Rob Pike is [displeased](https://bsky.app/profile/robpike.io/post/3matwg6w3ic2s); and the
issue of AI-generated changelists came up on
[golang-dev](https://groups.google.com/g/golang-dev/c/4Li4Ovd_ehE), recently.
Russ Cox took the time to elaborate on the issue in a [longer
response](https://groups.google.com/g/golang-dev/c/4Li4Ovd_ehE/m/8L9s_jq4BAAJ):

> AI tools have seduced many people into a false belief that these fundamentals
> no longer apply. People brag about codebases of hundreds of thousands of
> lines that have never been viewed by people, churned out in record time. On
> closer inspection, these codebases inevitably turn out to be more like
> dancing elephants than useful engineering artifacts.

[![](/downloads/lgo-58/dancing-elephant/dancingelephant-xs.gif)](/downloads/lgo-58/dancing-elephant/dancingelephant.gif)

Just CTRL-C to stop the music.

## Misc

* if you use CC, also give [pi](https://github.com/badlogic/pi-mono) a try
* set `"includeCoAuthoredBy": false` in `~/.claude/settings.json` to disable agent co-author commits
* set env `"CLAUDE_CODE_ATTRIBUTION_HEADER": "0"` [may improve](https://github.com/musistudio/claude-code-router/issues/1161#issuecomment-3788859149) token usage
* [Context7](https://context7.com/), which is a service agents can use for up-to-date API docs
* [A sufficiently detailed spec is the code](https://haskellforall.com/2026/03/a-sufficiently-detailed-spec-is-code), remember, the [Silver Bullet](https://worrydream.com/refs/Brooks_1986_-_No_Silver_Bullet.pdf)?
* the [Kilo](https://kilo.ai/) orchestrator uses [memory banks](https://kilo.ai/features/memory-bank)
* AI slop is gluey
* experiential: keep your context usage below 40% for better results; do not
  clutter context with too many MCP definitions; minimize or optimizue context
(maybe try [cxpak](https://github.com/Barnett-Studios/cxpak),
[oo](https://github.com/randomm/oo),
[repoctx](https://github.com/miku/repoctx), ...)
* mandate to use AI tools varies across companies; depending on your context,
  AI tools speed you up or even slow you down
* if generating code is cheap and fast, bottlenecks appear elsewhere, e.g. in review
* switch quickly between model providers for claude code, with [claude-switch](https://github.com/miku/claude-switch)

## Leaks and recursion

LLM [GLM](https://z.ai/blog/glm-5) taking a few minutes to analyze the source
code of an agentic coding harness (that
[leaked](http://web.archive.org/web/20260331105520/https://x.com/Fried_rice/status/2038894956459290963)); summary report (more deep dives in [rosaboyle/awesome-cc-oss](https://github.com/rosaboyle/awesome-cc-oss)):
[2026-04-01-CC-BY-GLM.md](https://github.com/golang-leipzig/golang-leipzig.github.io/blob/source/static/downloads/lgo-58/2026-04-01-CC-BY-GLM.md); with [glow](https://github.com/charmbracelet/glow) md reader:

> glow -p https://tinyurl.com/2026-04-01-CC-BY-GLM

[![](/images/2026-04-01-an-agent-analyzing-an-agentic-coding-harness.gif)](/images/2026-04-01-an-agent-analyzing-an-agentic-coding-harness.gif)




Thanks again [Akim](https://www.linkedin.com/in/akimzmerli/) for the well structured and engaging presentation.

----

[Join our meetup](https://www.meetup.com/de-DE/leipzig-golang/) to get notified of upcoming events.
