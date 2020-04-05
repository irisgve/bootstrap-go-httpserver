package main

import (
	"github.com/irisgve/bootstrap-go-httpserver/internal"
)

func main() {
	cfg, err := internal.GetConfig()
	if err != nil {
		panic(err)
	}

	server, err := internal.NewServer(cfg)
	if err != nil {
		panic(err)
	}

	server.Start()
}
