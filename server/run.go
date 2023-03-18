package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func RunServer() error {
	grpcListener, err := net.Listen("tcp", ":4505")
	if err != nil {
		log.Fatalf("error setting up net.Listen(): %s", err)
	}

	grpcServer := grpc.NewServer()
	registerSubscriberServics(grpcServer)

	err = grpcServer.Serve(grpcListener)
	if err != nil {
		return fmt.Errorf("failed to start gRPC server on port 4505: %s", err)
	}

	return nil
}
