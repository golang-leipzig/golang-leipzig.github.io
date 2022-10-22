# Go and Cloud-Native Leipzig

![github pages](https://github.com/golang-leipzig/golang-leipzig.github.io/workflows/github%20pages/badge.svg)

Find us on [Meetup.com](https://www.meetup.com/Leipzig-Golang/).

## Structure

We aim to write an announcement and a wrapup article for each of our meetups.  To create a new article I recommend to use something like

```sh
$ hugo new content/posts/meetup-123-invitation.md
```

but it's also fine to copy-paste from an existing article.

## Development

We use [hugo][1] to render the static web pages.  GitHub Actions will render the webpage automatically on each commit to the source branch.  The website, [golangleipzig.space][2] is hosted on GitHub pages.

To get a local preview of the website run

```sh
$ hugo serve
```

and open <http://localhost:1313> in your browser.  The website will be rendered again if any file has changed.  If you see an error like

```
failed to extract shortcode: template for shortcode "h2" not found
```

then make sure that you cloned or updated the subrepositories, e.g. by running

```sh
$ git submodule update --init --recursive
```

I prefer to install `hugo` with `go install github.com/gohugoio/hugo@latest`, but you can also use `brew install hugo` or whatever package manager is available on your system.


[1]: https://gohugo.io
[2]: https://golangleipzig.space
