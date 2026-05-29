# README

gridanim is a small Go program that generates nice gif animations. The canvas
is subdivided into a grid. Each grid tile can use a different color. Then,
following some algorithms, we can change the color of individual tiles. We
could for example have random colors and then have a single dot with a fixed
color wandering around the grid in a random walk like way.

There are a couple of basic "animations" we could think of:

* animate a horizontal bar, from bottom to to top (or down)
* animate a vertical bar
* animate a fixed set of points, e.g. a shape in a random walk, or some "linear" movement

There are many other primitives we could think of, in the best case, some
animations can be combined and overlayed in an interesting way.

For the first iterations, we want to have a slightly advanced animation. We
want to render a text in on hour grid. The user can specify what text in
stricly ASCII letters and digits.

The grid animation could then start from a random fluctuation of colors to
successively drawing out the text of the user. And once fully visible then keep
it, freeze the frames, so the viewer has enough time to read.

I once wrote a program for slight animations in Go in
@/home/tir/code/miku/nightjet/x/nb/ if that is of help.

## Task

Write the basic components necessary into a single file. Use either Go or
Python, which ever is more conventient. If you choose Python, assume "uv" is
installed and that a script can be made self contained like this:

```
#!/usr/bin/env -S uv run --script
#
# /// script
# requires-python = ">=3.12"
# dependencies = ["somedep", "other"]
# ///

...
```

The user should be able to adjust the palette (does not need super fine
control, just some basic variants), specify a text, specify the individual tile
size, overall canvas size. The letters can be rendered roughly, but should be
readable on the grid. There is probably a mapping from letter to grid tiles. If
possible, the size of the text could be adjusted by the user as well.

The program can be called gridanim, gridgif, giftiles, tileanim, tiledanim,
tileseq, or similar. I like tileseq, because there are tiles and we are working
towards a small animation.

Do not care about gif optimization too much. The user should have some control
over the size by adjusting the overall size. And then there is also tools like
gifsicle, etc.
