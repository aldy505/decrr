# Decrr

[![Go Reference](https://pkg.go.dev/badge/github.com/aldy505/decrr.svg)](https://pkg.go.dev/github.com/aldy505/decrr) [![Go Report Card](https://goreportcard.com/badge/github.com/aldy505/decrr)](https://goreportcard.com/report/github.com/aldy505/decrr) ![GitHub](https://img.shields.io/github/license/aldy505/decrr) [![CodeFactor](https://www.codefactor.io/repository/github/aldy505/decrr/badge)](https://www.codefactor.io/repository/github/aldy505/decrr) [![codecov](https://codecov.io/gh/aldy505/decrr/branch/master/graph/badge.svg?token=Noeexg5xEJ)](https://codecov.io/gh/aldy505/decrr) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/9b78970127c74c1a923533e05f65848d)](https://www.codacy.com/gh/aldy505/decrr/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=aldy505/decrr&amp;utm_campaign=Badge_Grade) [![Test](https://github.com/aldy505/decrr/actions/workflows/ci.yml/badge.svg)](https://github.com/aldy505/decrr/actions/workflows/ci.yml)

A modification (and a bit of simplification) of the [tracerr](https://github.com/ztrue/tracerr) package.

This essentially does pretty much the same, but instead of returning another struct like tracerr package does, this decorate the error and return another regular error type.

## Usage

Simply put it on the deepest function/method on your project.

```go
package main

import "github.com/aldy505/decrr"

func main() {
  err := ExecuteSomething()
  log.Fatal(err)
}

func ExecuteSomething() error {
  return SomethingElse()
}

func SomethingElse() error {
  err := errors.New("a goat just passes by!")
  return decrr.Wrap(err)
}
```

And yes, it only has one `.Wrap(error)` function. Nothing else.

## License

[MIT License](./LICENSE)