package main

import (
	"github.com/GlebSr/app/internal/app/server"
	"log"
)

func main() {
	print("starting")
	log.Fatal(server.Start())
}
