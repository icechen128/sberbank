[![LICENSE](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)
[![Go](https://github.com/icechen128/sberbank/actions/workflows/go.yml/badge.svg)](https://github.com/icechen128/sberbank/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/icechen128/sberbank/branch/master/graph/badge.svg)](https://codecov.io/gh/icechen128/sberbank)
[![Go Report Card](https://goreportcard.com/badge/github.com/icechen128/sberbank)](https://goreportcard.com/report/github.com/icechen128/sberbank)
[![Godocs](https://img.shields.io/badge/golang-documentation-blue.svg)](https://godoc.org/github.com/icechen128/sberbank)
# Sberbank Acquiring API Wrapper

[Sberbank Acquiring API](https://securepayments.sberbank.ru/wiki/doku.php/integration:api:start) written in Go

## Installation

Make sure your project is using Go Modules (it will have a `go.mod` file in its
root if it already is):

``` sh
go mod init
```

Then, reference stripe-go in a Go program with `import`:

``` go
import (
	"github.com/icechen128/sberbank/acquiring"
	"github.com/icechen128/sberbank/currency"
)
```

Run any of the normal `go` commands (`build`/`install`/`test`). The Go
toolchain will resolve and fetch the module automatically.

Alternatively, you can also explicitly `go get` the package into a project:

```bash
go get -u github.com/icechen128/sberbank
```

## Getting started

### Step 1
Get password and username 

### Step 2
Set mode sandbox or production, currency, language

### Step 3

Configure client as in an example below

Example below:

```golang
package main

import (
	"context"
	"fmt"
	"github.com/icechen128/sberbank/acquiring"
	"github.com/icechen128/sberbank/currency"
)

func main() {
	cfg := acquiring.ClientConfig{
		UserName:           "test-api", // Replace with your own
		Currency:           currency.RUB,
		Password:           "test", // Replace with your own
		Language:           "ru",
		SessionTimeoutSecs: 1200,
		SandboxMode:        true,
	}

	client, err := acquiring.NewClient(&cfg)
	if err != nil {
		panic(err)
	}
	order := acquiring.Order{
		OrderNumber: "test",
		Amount:      100,
		Description: "My Order for Client",
	}
	result, _, err := client.RegisterOrder(context.Background(), order)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.ErrorCode)
	fmt.Println(result.ErrorMessage)
	fmt.Println(result.FormUrl)
	fmt.Println(result.OrderId)
}

```

### Step 4

Run example `go build example.go`




