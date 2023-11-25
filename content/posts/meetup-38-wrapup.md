---
title: "Hybrid Meetup #38 wrap-up"
date: 2023-11-22T22:00:00+01:00
draft: false
tags:
- summary
- meetup
---

## A cache is worth a thousand searches

Hybrid meetup #38 took place [2023-11-20
19:00](https://www.meetup.com/leipzig-golang/events/290666185/) at
[CHECK24](https://check24.de) Leipzig Office.

We had a great presentation from
[Fabian](https://www.linkedin.com/in/fabian-g%C3%A4rtner-913584141/) about
building a mission critical query caching database (MemoryDB) in Go and the
challenges involved.

One challenge is the variety of query parameters and nearly overlapping
values and ranges. A key for performance is to express various parameter values (e.g.
date boundaries) with the help of [BitMaps](https://en.wikipedia.org/wiki/Bit_array),
e.g. [Roaring Bitmaps](https://r-libre.teluq.ca/1402/1/1709.07821.pdf).

> Bitmap indexes are commonly used in databases and search engines. By
> exploiting bit-level parallelism, they can significantly accelerate queries. -- [RoaringBitmap.pdf](https://r-libre.teluq.ca/602/1/RoaringBitmap.pdf)

Among other things, the project also uses [tableflip](https://github.com/cloudflare/tableflip) - a library that allows you to

> update the running code and / or configuration of a network service, without disrupting existing connections.

A very practical concern has been the struct design for GC-friendlyness. Pop
quiz:

* Is the following struct GC-friendly?
* When would it become a problem?
* What could be improved?

```go
// Offer, abridged.
type Offer struct {
    HomeAirport        *Airport
    DestinationAirport *Airport
    Hotel              *Hotel
    RoomType           string
    MealType           string
    Airline            string
    DepartureTime      time.Time
    ReturnTime         time.Time
}
```

If possible, one can try to use stack allocated values (note: Go
[ref/spec](https://go.dev/ref/spec) never mentions to stack or heap, as these
concepts are abstracted by the language):

```go
// Offer, reduced, abridged.
type Offer struct {
    HomeAirportID        int
    DestinationAirportID int
    HotelID              int
    RoomType             RoomTypeEnum
    MealType             MealTypeEnum
    AirlineID            int
    DepartureTime        int64
    ReturnTime           int64
}
```

This is now a much more compact, GC-friendly struct that will require
additional object lookups for respective identifiers but would reduce GC load
significantly, when dealing with millions of objects. Simple, effective.

![](/images/meetup-38-pic-sketch.jpg)

## Testdriving [OLLAMA](https://ollama.ai)

A lightning talk was concerned with [Testdriving
OLLAMA](https://github.com/miku/localmodels) - a packaging tool for large
language model files. Ollama is inspired by docker and allows to wrap LLM
customizations (parameters, context) into a easy to distribute format.

Thanks to projects like [LLAMA](https://ai.meta.com/llama/) and
[llama.cpp](https://github.com/ggerganov/llama.cpp) it is possible to
experiment with LLMs on everyday hardware, e.g. a 15W TDP [2017
CPU](https://www.intel.com/content/www/us/en/products/sku/122589/intel-core-i78550u-processor-8m-cache-up-to-4-00-ghz/specifications.html).

[![](/images/three-genai-haiku.png)](https://golangleipzig.space/meetup-38-llm-haiku/meetup-38-llm-haiku.pdf)

## Thanks!

Thanks [CHECK24](https://check24.de) for hosting Leipzig Gophers November 2023
Meetup, [Fabian](https://www.linkedin.com/in/fabian-g%C3%A4rtner-913584141/) for the
great talk and [Florian](https://www.linkedin.com/in/florianbr%C3%A4utigam/)
for the excellent event organisation.

----

Are you using an interesting data structure like bitmaps to improve
performance? Then [join our meetup](https://www.meetup.com/Leipzig-Golang/) and
tell us about it!

