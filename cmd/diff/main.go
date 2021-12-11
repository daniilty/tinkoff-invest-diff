package main

import (
	"context"
	"fmt"
	"log"

	invest "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/daniilty/tinkoff-invest-diff/records"
)

func run() error {
	cfg := getConfig()

	client := invest.NewRestClient(cfg.token)

	p, err := client.Portfolio(context.Background(), cfg.id)
	if err != nil {
		return err
	}

	diffString := records.GetDiffString(p.Positions)

	fmt.Print(diffString)

	return nil
}

func main() {
	err := run()

	if err != nil {
		log.Fatal(err)
	}
}
