package main

import (
	"log"

	"github.com/bbruun/grpc-test-2/server"
)

func main() {
	err := server.RunServer()
	if err != nil {
		log.Fatal(err)
	}
}
