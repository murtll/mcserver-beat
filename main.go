package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/murtll/mcserver-beat/internal/worker"
	"github.com/murtll/mcserver-beat/rest"
)

func main() {
	var health bool
	var beatOnly bool
	var apiOnly bool
	flag.BoolVar(&health, "health", false, "make healthcheck and exit.")
	flag.BoolVar(&beatOnly, "beat-only", false, "start only beat worker.")
	flag.BoolVar(&apiOnly, "api-only", false, "start only REST API.")
	flag.Parse()

	if health {
		status, err := rest.Health()
		if err != nil {
			fmt.Println(status, "\n", err)
			os.Exit(1)
		}
		fmt.Println(status)
		os.Exit(0)
	}

	if beatOnly {
		worker.Start()
		os.Exit(0)
	}

	if apiOnly {
		log.Fatal(rest.Start())
	}

	go worker.Start()
	log.Fatal(rest.Start())
}