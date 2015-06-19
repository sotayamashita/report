# Report

[![GoDoc](https://godoc.org/github.com/sotayamashita/report?status.svg)](https://godoc.org/github.com/sotayamashita/report)

## Description

Generate report for daily in Golang. I'm not all knowing so if you find anything you think I got wrong, needs more details or have any suggestion or questions to make this guide better, please create any [issues](https://github.com/sotayamashita/report/issues)

## Features

* [x] [Toggl](https://www.toggl.com/)
 * [ ] Sort by project
 * [ ] Sort by client
* [ ] Git

## Usage

### Toggl

Run command:

```bash
$ report toggl
```

Result:

```
Create CLI tool with Golang OSS [golang, oss] 42m6s 8:43 PM - 9:16 AM
Translate docs OSS [electron, oss] 42m6s 8:43 PM - 9:16 AM
...
```

### Git

Run command:

```bash
$ report git
```

Result:

```bash
git@github.com:sotayamashita/report.git
Sat Jun 13 21:40:30 2015 +0900 b954dfa Rename git\'s feature name
Sat Jun 13 21:36:47 2015 +0900 326e239 Add new features
Sat Jun 13 21:34:50 2015 +0900 c58ef8d Finish toggl feature
Sat Jun 13 21:33:42 2015 +0900 c080ae4 Rename const togglTimeForm
```

## Configuration

please create `config.yml`:

### Toggl

In config file, `api_token` and `workspace_id` is taken from [toggl.com](https://www.toggl.com/) and [Toggl API Documentation](https://github.com/toggl/toggl_api_docs#api-token) is more detail.

```yaml
toggl:
  api_token: "<your api token>"
  workspace_id: "<your workspace_id>"
```

### Git

In config file, each item should be a git repository (or any of the parent directories). It means that each directory has `.git/`.

```yaml
git:
repositories:
 - "~$GOPATH/src/github.com/sotayamashita/report",
 - "~$GOPATH/src/github.com/sotayamashita/coffee"
```

## Installation

To install, use `go get`:

```bash
$ go get -d github.com/sotayamashita/report
$ godep go install
$ go build
$ go install
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
