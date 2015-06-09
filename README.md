# Report

## Description

Generate report for daily in Golang

## Features

* [ ] [Toggl](https://www.toggl.com/)

## Requirements

* [request](https://github.com/mozillazg/request) - Go HTTP Requests for Humans™
* [cli](https://github.com/codegangsta/cli) - A small package for building command line apps in Go
* [cli-init](https://github.com/tcnksm/cli-init) - The easy way to start building Golang command-line application.

## Usage

```bash
$ report toggl
```

## Configuration

please create `config.tml`:

```
[toggl]
api_token = "<your api token>"
```

## Installation

To install, use `go get`:

```bash
$ go get -d github.com/Sota Yamashita/report
```

## Contribution

1. Fork ([https://github.com/Sota Yamashita/report/fork](https://github.com/Sota Yamashita/report/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## License

MIT © [Sota Yamashita](https://github.com/Sota Yamashita)
