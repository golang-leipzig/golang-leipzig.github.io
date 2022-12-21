---
title: "Drinkup #33 wrap-up"
date: 2022-12-21T20:00:31+02:00
draft: false
tags:
- summary
- meetup
---

## Happy Holidays!

![](/images/cantona_pixel.png)

We met 2022-12-20 19:00 at Café Cantona, with no specific Go agenda, for a
great exchange of ideas and experiences - too many to capture them in a blog
post, so here's just a fraction of the topics we passed (slightly categorized):

## Book recommendations

* [Grokking Functional Programming](https://www.manning.com/books/grokking-functional-programming) - a very readable introduction to the topic; more academic angles include the classic [Out of the tar pit](https://curtclifton.net/papers/MoseleyMarks06a.pdf) or [Purely Functional Data Structures](https://openlibrary.org/works/OL1863182W/Purely_functional_data_structures) (always nice to [find references to this book in code](https://github.com/golang/go/blob/78fc81070a853d08a71f70fa20b2093f5535e6c5/src/net/http/transport.go#L1256-L1271))
* in 04/2022, a JS version of [SICP](https://en.wikipedia.org/wiki/Structure_and_Interpretation_of_Computer_Programs) (1984) came out: [SICP JS edition](https://en.wikipedia.org/wiki/Structure_and_Interpretation_of_Computer_Programs,_JavaScript_Edition), there's an [open access version](https://sourceacademy.org/sicpjs/index), too
* [Podman in Action](https://www.manning.com/books/podman-in-action), container tools with a nice Unix backdrop
* the classic, albeit just five year old [DDIA](https://www.oreilly.com/library/view/designing-data-intensive-applications/9781491903063/)
* [100 mistakes in Go and how to avoid them](https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them) - actionable advice for Go code improvements

## Software General

* [ISO 28500:2017](http://bibnum.bnf.fr/WARC/), aka [WARC](https://web.archive.org/web/20120619151338/http://www.iwaw.net/05/kunze.pdf) is how you can capture the web - and there are many libraries and tools around this format, just to name a few: [Heritrix](https://en.wikipedia.org/wiki/Heritrix), [brozzler](https://github.com/internetarchive/brozzler), [webrecorder](https://github.com/webrecorder), [warcprox](https://github.com/internetarchive/warcprox), [replayweb.page](https://replayweb.page/), [and more](https://github.com/iipc/awesome-web-archiving)

If you have [wget](https://www.gnu.org/software/wget/) installed, you can already create WARC files yourself:

```shell
$ wget -rkc --warc-file golangleipzig.space --warc-cdx https://golangleipzig.space
```

<!--

```
Opening WARC file ‘golangleipzig.space.warc.gz’.

--2022-12-21 16:26:47--  https://golangleipzig.space/
Resolving golangleipzig.space (golangleipzig.space)... 185.199.110.153, 185.199.109.153
Connecting to golangleipzig.space (golangleipzig.space)|185.199.110.153|:443... connected.
HTTP request sent, awaiting response...

...

$ ls -hgG golangleipzig.space.*
-rw-rw-r-- 1 88K Dec 21 16:27 golangleipzig.space.cdx
-rw-rw-r-- 1 18M Dec 21 16:27 golangleipzig.space.warc.gz

$ wc -l golangleipzig.space.cdx # urls captured
352 golangleipzig.space.cdx

$ shuf -n 3 golangleipzig.space.cdx | awk '{print $1}'
https://golangleipzig.space/images/meetup_30_goodies_tile.png
https://golangleipzig.space/images/christmasxpalm01.gif
https://golangleipzig.space/tags/summary/
```

-->

Since web archiving aims for a complete preservation, the whole HTTP exchange
is recorded - which can be analyzed later. The [HTTP Archive](https://httparchive.org/) creates various reports, e.g. a
[yearly state of the web report](https://almanac.httparchive.org/en/2022/).

* Parquet is a great columnar storage format, with a [few libraries](https://pkg.go.dev/search?q=parquet) in Go, too, e.g. from the [Apache Arrow](https://github.com/apache/arrow/tree/master/go) project

There's a new wave of big data and analytics tools written in [rust](https://www.rust-lang.org/), e.g. the [distributed SQL query engine ballista](https://github.com/apache/arrow-ballista) [NY statprog meetup recording](https://www.youtube.com/watch?v=ZZHQaOap9pQ&t=397s), and others

* Did you know? Google BigQuery does not have `MEDIAN` - [workaround](https://www.pascallandau.com/bigquery-snippets/calculate-median/)

## Go

* JSON matcher: [quamina](https://github.com/timbray/quamina)

> Quamina implements a data type that has APIs to create an instance and add
> multiple Patterns to it, and then query data objects called Events to
> discover which of the Patterns match the fields in the Event.

* Go generics utils: [gogu](https://github.com/esimov/gogu)

> Gogu is a versatile, comprehensive, reusable and efficient concurrent-safe
> utility functions and data structures library taking advantage of the Go
> generics. It was inspired by other well established and consecrated
> frameworks like lodash or Apache Commons and some concepts being more closer
> to the functional programming paradigms.

* Functional options [pattern](https://github.com/tmrts/go-patterns/blob/master/idiom/functional-options.md) (we talked about that in
  [#13](https://golangleipzig.space/posts/meetup-13-wrapup/)) has some
advantages over the still widely used technique of passing config structs to
functions.

### A Go interface puzzle

> Declare an interface where it is used, not where it is implemented. Unless
> the interface is well discovered.

Via: [1605116543553019905](https://twitter.com/inancgumus/status/1605116543553019905) -- join the [conversation](https://twitter.com/embano1/status/1605173329836404738) ...

Our take: It's a - slightly subtle - consequence of structural
typing: An interface can be *added later*, in an ad-hoc style, e.g. to facilitate testing
or establishing a protocol of limited scope.

Example: A function only uses `client.Do` for HTTP requests. We can create a
protocol, so we are able to use the default HTTP client from the standard
library or some third-pary library as well (as long at it has a suitable `Do` method). We can define a Doer:

```go
type Doer interface {
    Do(req *http.Request) (*http.Response, error)
}
```

and let our function work with `Doer` interface instead of a concrete type.
This interface should be defined where it is used (and not in the net/http
package, for example). On the flipside, sometimes you need central interfaces,
e.g. [io.Reader](https://pkg.go.dev/io#Reader), or a database abstraction
layer, or a [filesystem abstraction
layer](https://github.com/rclone/rclone/blob/5ac8cfee56b58c242c60b8fd319b8a2dd4420c9b/fs/types.go), and so on. These are *well discovered*.

> In a structural setting, a type expression is a closed entity: it carries
> with it all the information that is needed to understand its meaning. In a nominal system, we
> are always working with respect to some global collection of type names and
> associated deﬁnitions. This tends to make both deﬁnitions and proofs more
> verbose.

More on nominal and structural typing can be found in [Chapter
19.3](https://www.cis.upenn.edu/~bcpierce/tapl/contents.pdf#page=5) of [Types
and programming
languages](https://en.wikipedia.org/wiki/Types_and_Programming_Languages)
[TAPL](https://www.cis.upenn.edu/~bcpierce/tapl/).

## Python

* [pip-tools](https://github.com/jazzband/pip-tools) is a small tool for pinning dependencies and predictable builds
* Python in 2022 with [gradual typing](https://en.wikipedia.org/wiki/Gradual_typing) feels like a static language now - a long way since [Static Typing Where Possible, Dynamic Typing When Needed:
The End of the Cold War Between Programming Languages](http://web.archive.org/web/20060111181527/http://pico.vub.ac.be/~wdmeuter/RDL04/papers/Meijer.pdf) (discussed at the time by [LtU](http://lambda-the-ultimate.org/node/834))

## API Design

* [Expanding objects](https://stripe.com/docs/api/expanding_objects) - another example in the wild: [fatcat API](https://api.fatcat.wiki/v0/release/qaa7ysrn5rfbnkjec7rtrkcao4?expand=files) (*List of sub-entities to expand in response. For releases, 'files', 'filesets, 'webcaptures', 'container', and 'creators' are valid.*)
* [Idempotency tool for AWS lambda](https://awslabs.github.io/aws-lambda-powertools-python/2.4.0/utilities/idempotency/)

Related, from [Cloud Native Go](https://www.oreilly.com/library/view/cloud-native-go/9781492076322/) (a book we gave away in [#30](https://golangleipzig.space/posts/meetup-30-wrapup/), courtesy of [O'Reilly Media](https://www.oreilly.com/pub/cpc/323592)), page 168:

> [Holly Cummins](https://hollycummins.com/), the worldwide development community practice lead for the IBM
Garage, famously said that *if cloud native has to be a synonym for anything, it would be idempotent*.

* [Joshua Bloch: Bumper-Sticker API Design](https://www.infoq.com/articles/API-Design-Joshua-Bloch/) ([Slides](https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/32713.pdf))


## Cloud

* [Running kubernetes with rootless podman on WSL2](https://www.salilmishra.ml/posts/k8s-podman-wsl2/)

> Running Kubernetes with rootless podman is documented on both kind and
> minikube but to get it up and running on [WSL2](https://learn.microsoft.com/en-us/windows/wsl/install) requires some additional tweaks
> as by default WSL2 uses init daemon, instead of systemd and you have to
> enable cgroupv2 explicitly.

* The perennial question: how can you beat *cloud* infrastructure from a
  developer experience (and economical) standpoint? For example, how can we create a robust, scalabale
queue with less effort that writing e.g. [15 lines in
Go](https://github.com/awsdocs/aws-doc-sdk-examples/blob/5458e2b9fd71abb916bca4ed53d8c1a894e4fe87/go/example_code/sqs/sqs_createqueues.go#L42-L59) (including error handling);
(btw, [SQS](https://aws.amazon.com/blogs/aws/aws-blog-the-first-five-years/)
started in 2004); at some places, [infra works differently](https://archive.org/details/jonah-edwards-presentation) [❤️](https://archive.org/donate/)


## Culture

* [Visual Studio Code is designed to fracture](https://ghuntley.com/fracture/)

## Writing and blogs

* How to improve as an software engineer? We got some ideas, like keeping a
  (public) technical journal, but are curious about what you think - so [join
our meetup](https://www.meetup.com/Leipzig-Golang) to connect and to get
notified of upcoming events!

## Misc

* Airtags and Parcel: [DHL "Lost" My AirTag Parcel (but I knew where it was)](https://www.youtube.com/watch?v=W8SER24F0U8)

