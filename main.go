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
	flag.BoolVar(&health, "health", false, "make healthcheck and exit.")
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

	go worker.Start()
	log.Fatal(rest.Start())
}