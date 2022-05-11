---
title: "Virtual Meetup #27 wrap-up"
date: 2022-05-11T10:00:00+02:00
draft: false
tags:
- summary
- meetup
---

## Algorithm Challenge

[Meetup #27](https://www.meetup.com/Leipzig-Golang/events/285798083/) took
place on 2022-05-10 19:00 CEST and was virtual again and we had an interesting
problem-solving live-coding challenge.

Slides and code: [gitlab.com/telo_tade/prefix_suffix_arrays](https://gitlab.com/telo_tade/prefix_suffix_arrays)

### Problem

> You are given a list of numbers. Find its longest slice that sums to zero.


First, clarify:

* Is it a consecutive list of numbers?
* What kind of numbers, and how many?

We collaboratively went from a naive *O(n^3)* solution to a slightly improved
*O(n^2)* using an auxiliary data structure: a [prefix sum
array](https://en.wikipedia.org/wiki/Prefix_sum).

![](/images/Prefix_sum_16.svg)

Takeaways:

* Any auxiliary data structure that takes *O(n)* to build can be considered
  *free*, as most of the time a problem requires to iterate over all the data
  at least once anyway.
* Building a prefix sum array requires *O(n)* and is simple to build, e.g. with
  a recursive formulation.

The key insight may become obvious in an example:

```
input:      [1, 2, 3,  4, -3, -1,  9]
prefix sum: [1, 3, 6, 10,  7,  6, 15]
```

Two elements in the prefix array that have the same value (e.g. 6) allow us to
determine a subsequence that sums to zero.

A final pass through the prefix sum array allows to keep track of repeated
numbers and find the longest sequence. If you iterate throught the prefix sum
array from front and back simultaneously, you can stop at the first occurence
(same time complexity, still).

We implemented a version that defined a type set for numbers; similar to the
ones found in
[x/exp/constraints](https://pkg.go.dev/golang.org/x/exp/constraints).

In this case, a non-generic version may use reflection and be slower and more
inconvenient to write.

Find out more at:
[gitlab.com/telo_tade/prefix_suffix_arrays](https://gitlab.com/telo_tade/prefix_suffix_arrays);
we enjoyed the interactive format. If that's something for you too, check out
[Hamburg Whiteboard Coders](https://www.meetup.com/hamburg-whiteboard-coders/)
meetup.

----

[Join our meetup](https://www.meetup.com/Leipzig-Golang) to get notified of
upcoming events!

