---
title: "Hybrid Meetup #38 wrap-up"
date: 2023-11-22T22:00:00+01:00
draft: false
tags:
- summary
- meetup
---

## Cache Rules Everything Around Me

Hybrid meetup #38 took place [2023-11-20
19:00](https://www.meetup.com/leipzig-golang/events/290666185/) at
[CHECK24](https://check24.de) Leipzig Office.

We had a great presentation from
[Fabian](https://www.linkedin.com/in/fabian-g%C3%A4rtner-913584141/) about
building a mission critical query caching database (MemoryDB) in Go and the
challenges involved.

A challenge is the sheer variety of query parameters and near overlapping
values. A key for performance is to express various parameters values (e.g.
dates) with the help of [BitMaps](https://en.wikipedia.org/wiki/Bit_array),
especially [Roaring Bitmaps](https://r-libre.teluq.ca/1402/1/1709.07821.pdf).

> Bitmap indexes are commonly used in databases and search engines. By
> exploiting bit-level parallelism, they can significantly accelerate queries. -- [RoaringBitmap.pdf](https://r-libre.teluq.ca/602/1/RoaringBitmap.pdf)

Among other things, the project also uses [tableflip](https://github.com/cloudflare/tableflip) - a library that allows you to

> update the running code and / or configuration of a network service, without disrupting existing connections.

A very practical concern has been the struct design for GC-friendlyness. Pop
quiz: Is the following struct GC-friendly? When would it become a problem? What
could be improved?

```go
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

One way to go about it would be to reduce any heap allocated value with stack allocated value, if possible, e.g.

```go
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

## Testdriving OLLAMA

Another, lightning talk was concerned with [Testdriving
OLLAMA](https://github.com/miku/localmodels). A human selection of three Go
programming related haiku generated on a 15W TDP [2017
CPU](https://www.intel.com/content/www/us/en/products/sku/122589/intel-core-i78550u-processor-8m-cache-up-to-4-00-ghz/specifications.html).
Note: A typical haiku follows the 5-7-5 scheme.

![](/images/three-genai-haiku.png)

## Thanks!

Thanks [CHECK24](https://check24.de) for hosting Leipzig Gophers November 2023
Meetup, [Fabian](https://www.linkedin.com/in/florianbr%C3%A4utigam/) for the
great talk and [Florian](https://www.linkedin.com/in/florianbr%C3%A4utigam/)
for the excellent event organisation.

----

Are you using an interesting data structure like bitmaps to improve
performance? Then [join our meetup](https://www.meetup.com/Leipzig-Golang/) and
tell us about it!

