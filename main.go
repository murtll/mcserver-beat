package main

import (
	"log"

	"github.com/murtll/mcserver-beat/internal/worker"
	"github.com/murtll/mcserver-beat/rest"
)

func main() {
	go worker.Start()
	log.Fatal(rest.Start())
}