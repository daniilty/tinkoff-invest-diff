package main

import "flag"

type config struct {
	token string
	id    string
}

func getConfig() *config {
	cfg := config{}

	flag.StringVar(&cfg.token, "t", "", "tinkoff openapi token")
	flag.StringVar(&cfg.id, "id", "", "tinkoff account id")
	flag.Parse()

	return &cfg
}
