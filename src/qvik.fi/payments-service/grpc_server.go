package main

import (
	"fmt"
	"net"

	payments "qvik.fi/payments"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func mustCreateServer() *grpc.Server {
	opts := []grpc.ServerOption{}
	return grpc.NewServer(opts...)
}

// Initializes & runs the gRPC server in another go routine. Returns
// created grpc.Server.
func mustRunGrpcServer(port int) *grpc.Server {
	log.Debugf("Starting to listen to insecure HTTP/2 gRPC port %v..", port)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := mustCreateServer()
	payments.RegisterPaymentsServer(s, &paymentsServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		// Serve() will block until aborted
		log.Debugf("gRPC server listening..")
		err = s.Serve(listen)
		if err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	return s
}
