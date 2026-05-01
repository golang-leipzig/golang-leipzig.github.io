# yoloAI impressions

> 2026-05-01, [yoloai](https://github.com/kstenerud/yoloai), [yoloai.dev](https://yoloai.dev), [LGO](https://golangleipzig.space/) [#59](https://golangleipzig.space/posts/meetup-59-wrapup/)

## Impressions

Some impressions after setting up yoloAI and running it.

```
$ yoloai version
yoloai version dev (commit: 7ee813c, built: 2026-04-30T21:03:10Z)
```

* the value proposition is great: clear separation and faster agent progress; clear handover via diff; multiple loops supported;
* after setup, yoloai looks very robust, clear lifecycle, new session, kick off agent, check progress, decide on diff, destroy; repeat
* always easy to just destroy and start over
* yoloai subcommands are comprehesive and clear, e.g. `yoloai diff`, `yoloai ls`, ...
* ...
* the **system setup** step looks a bit heavyweight, as it assembles the docker image (lots of network access to fetch components, etc.):

```
$ docker images 2> /dev/null | grep yolo
yoloai-base:latest                                     f5efb33ad86f       7.06GB         1.84GB
```

Transparent operations are great, it may be ok to *hide the image assembly* behind a few
log messages, that could explain what is going on; "assembling base image",
"installing agents", ...

* on startup being dropped into a new **tmux** session may be surprising; when you
  already run tmux, the user needs to know about CTRL-b CTRL-b, etc.
* inner pane can be "Pane is dead" quickly, by accident (but easy enough to destroy and start over)
* yoloai log -f cannot be stopped unless the agent exists?

## Development

* adding a new agent is not too hard, but not requires changes in a couple of places

## Summary

Great sandboxing tool that requires a bit of setup, but once that's done offers
a familiar lifecycle for secure agentic coding. Agent agnostic, but requires
code changes to add new agents of additional tools when building the sandbox
image.
