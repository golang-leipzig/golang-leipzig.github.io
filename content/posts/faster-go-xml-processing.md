---
title: "Faster XML processing in Go"
date: 2024-03-26T12:00:01+01:00
draft: false
tags:
- go
- xml
- performance
---

## Prelude

Go supports XML processing in the standard library package
[xml](https://pkg.go.dev/encoding/xml). The [Marshal](https://pkg.go.dev/encoding/xml#Marshal) and [Unmarshal](https://pkg.go.dev/encoding/xml#Unmarshal)
function will use a [Encoder](https://pkg.go.dev/encoding/xml#Encoder) and [Decoder](https://pkg.go.dev/encoding/xml#Decoder) under the hood, as these are more general.

With a [Decoder](https://pkg.go.dev/encoding/xml#Decoder) it is possible to iterate over a large number of XML
elements and to deserialize data in a way that will limit memory consumption.
Here is a playground example for using a decoder to repeatedly parse elements
from a [Reader](https://pkg.go.dev/io#Reader): [efY60PYLFm8](https://go.dev/play/p/efY60PYLFm8).

One limitation of this approach is that you can only parse top-level elements.
This restriction is circumvented by utilities like
[xmlstream](https://github.com/miku/xmlstream), which allow to parse a number
of different elements from any level in the XML document in a streaming
fashion.

## Performance

XML decoding is slow, less because Go is slow (it usually is not), but because
parsing it can be slow (you can also try to use [libxml with
CGO](https://eli.thegreenplace.net/2019/faster-xml-stream-processing-in-go/) to
make it faster). After all, XML is a markup language, it can do things JSON
cannot do (like [TEI](https://tei-c.org/)). Surprisingly, many read world uses
of XML can be covered by JSON just as well. As a result, w can observe
a decline in XML usage and a de-facto standard choice of JSON for lots of data
exchange tasks and implementations (according to [google trends](https://trends.google.com/trends/explore?date=all&q=xml,json&hl=en-GB), the term *JSON*
surpassed *XML* in January 2016).

As data work often involves an ad-hoc data scouting step (with tools like
[jq](https://stedolan.github.io/jq/), or
[duckdb](https://www.pgrs.net/2024/03/21/duckdb-as-the-new-jq/)), some variants
of JSON gained popularity, like [jsonlines](https://jsonlines.org/) (also
called [JSON streaming](https://en.wikipedia.org/wiki/JSON_streaming) or
newline-delimited JSON). You can continue to use many UNIX text utils, while enjoying all
the features of [RFC 8259](https://datatracker.ietf.org/doc/html/rfc8259).

A format like jsonlines then makes it easy to parallelize JSON transformation
tasks with a [fan-out, fan-in](https://go.dev/blog/pipelines#fan-out-fan-in)
pattern: Read N lines, pass batch to goroutine, collect results and write them
out. A tool like [miku/parallel](https://github.com/miku/parallel) (["fast filter"](https://gist.github.com/miku/738f361c8156264626c74f9b717927ff#fast-parallel-filters-in-go)) allows to abstract away some of the parallel
processing boilerplate (example of [extracting a value from
jsonlines](https://github.com/miku/parallel/blob/27272f36538b21baa3256ec2e9487cca73d20628/examples/extract/extract.go#L1-L59),
twice as fast as jq).

## No lines for XML

XML does not have a common, line oriented representation. How can we
still process XML faster than iterating through it sequentially? We have to
parallelize it, but instead of relying on a newline for delimiting records, we
have to isolate the elements we are interested in, and batch hand
them over to processing threads. There is already a suitable type in the
standard library to split a stream into tokens.

## A Scanner quickly

The [bufio.Scanner](https://pkg.go.dev/bufio#Scanner) shows the utility of first class functions for customizing type behaviour (another
example is [Proxy in http.Transport](https://github.com/miku/httpgetaway/blob/master/ProxyIntro.md#customizing-httptransport-proxy)).
We can implement a custom [SplitFunc](https://pkg.go.dev/bufio#SplitFunc) that would split a stream on XML tags.

Since we only want the element boundaries, parsing the input stream is much
faster, as all we need to do is to find the
[Index](https://pkg.go.dev/bytes#Index) of the start and end tags in the
stream. Following the optimization by batching, we can collect N elements or
put a (soft) limit on the number of bytes in a batch and then pass a chunk of
valid XML to a processing function, which then can run in parallel and do the
heavy lifting of proper XML parsing.

## Splitting on tags

We implemented
[TagSplitter](https://github.com/miku/parallel/blob/27272f36538b21baa3256ec2e9487cca73d20628/record/split.go#L29-L56)
which will split a stream on XML elements and will batch them into
approximately 16MB sized chunks by default (it currently has the limitation
that it will not handle nested XML elements of the same name). You can then use
standard [bufio.Scanner](https://pkg.go.dev/bufio#example-Scanner-Custom)
facilities to get smaller batches of valid XML to parse with e.g.
[xmlstream](https://github.com/miku/xmlstream) (an example for parsing complex
PubMed XML document can be [found
here](https://github.com/miku/parallel/blob/27272f36538b21baa3256ec2e9487cca73d20628/examples/xmlstream/main.go#L36-L67),
including a cpu [pprof
viz](https://raw.githubusercontent.com/miku/parallel/master/examples/xmlstream/cpu.png)
showing further, potential performance improvements).

## Anecdata, millions of XML documents

Here is a rough summary of a test run of this approach (using a
contemporary [CPU](https://www.intel.com/content/www/us/en/products/sku/230498/intel-core-i913900t-processor-36m-cache-up-to-5-30-ghz/specifications.html))
on a dataset consisting of 327GB XML in about 36M documents (36557510) &mdash; that
is the set of publicly available metadata from
[PubMed](https://pubmed.ncbi.nlm.nih.gov/) (the target [struct](https://github.com/miku/parallel/blob/27272f36538b21baa3256ec2e9487cca73d20628/examples/xmlstream/main.go#L75-L318) requires a few hundred lines). The test ran in 03/2024. The
sequential approach takes **177 minutes**, the parallel approach brings this down to
about **20 minutes**, a 9x improvement ⚡ in throughput.

To put this into perspective, you can take a [metadata
dump](https://academia.stackexchange.com/questions/38969/getting-a-dump-of-arxiv-metadata)
of the [popular](https://info.arxiv.org/help/stats/2021_by_area/index.html)
[Arxiv](https://arxiv.org/) preprint server site (hosting about 2.4 million
scholarly articles) and parse all of its XML, which amounts to more than 5GB,
in about 8s. This makes XML processing more convenient - and fun, again.

