# Report

## Description

Generate report for daily in Golang. I'm not all knowing so if you find anything you think I got wrong, needs more details or have any suggestion or questions to make this guide better, please create any [issues](https://github.com/sotayamashita/report/issues)

## Features

* [ ] [Toggl](https://www.toggl.com/)
* [ ] git commit log

## Requirements

* [request](https://github.com/mozillazg/request) - Go HTTP Requests for Humans™
* [cli](https://github.com/codegangsta/cli) - A small package for building command line apps in Go
* [cli-init](https://github.com/tcnksm/cli-init) - The easy way to start building Golang command-line application.
* [go-simplejson](https://github.com/bitly/go-simplejson) - a Go package to interact with arbitrary JSON

## Usage

```bash
$ report toggl
```

## Configuration

please create `config.tml`:

```
[toggl]
api_token = "<your api token>"
workspace_id = "<your workspace_id>"
```

## Installation

To install, use `go get`:

```bash
$ go get -d github.com/Sota Yamashita/report
$ go build report.go
```

## Contribution

1. Fork ([https://github.com/Sota Yamashita/report/fork](https://github.com/Sota Yamashita/report/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create a new Pull Request

## License

MIT © [Sota Yamashita](https://github.com/Sota Yamashita)
