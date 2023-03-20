package main

import (
	"log"
	"strings"
	"time"

	"github.com/bbruun/grpc-test-2/messaging"
	"github.com/bbruun/grpc-test-2/server"
)

var mc *messaging.Minions

func init() {
	messaging.MinionStateCollector = messaging.NewMinions()
	mc = messaging.MinionStateCollector
}

func main() {

	go func() {
		for {
			names := mc.GetMinions()
			if len(names) > 0 {
				log.Printf("registered minions: %s\n", strings.Join(names, " "))
			}
			time.Sleep(time.Second)
		}
	}()

	err := server.RunServer()
	if err != nil {
		log.Fatal(err)
	}
}
