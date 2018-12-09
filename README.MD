[![LICENSE](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)
[![Build Status](https://travis-ci.com/helios-ag/sberbank-acquiring-go.svg?branch=master)](https://travis-ci.com/helios-ag/sberbank-acquiring-go)
[![Test Coverage](https://api.codeclimate.com/v1/badges/2d43f247eca97931f341/test_coverage)](https://codeclimate.com/github/helios-ag/sberbank-acquiring-go/test_coverage)

[![Godocs](https://img.shields.io/badge/golang-documentation-blue.svg)](https://godoc.org/github.com/helios-ag/sberbank-acquiring-go)

#Sberbank Acquiring API Wrapper

##Example

```golang
package main

import (
	"context"
	"fmt"
	"github.com/helios-ag/sberbank-acquiring-go/acquiring"
	"github.com/helios-ag/sberbank-acquiring-go/currency"
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
