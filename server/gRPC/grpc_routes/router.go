//
// EPITECH PROJECT, 2024
// AREA
// File description:
// router
//

package grpc_routes

import (
	"area/gRPC/api"
	helloworld "area/protogen/gRPC/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func LaunchServices() {
	const addr = "0.0.0.0:50051"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	helloService := api.NewHelloService(nil)
	dtService := api.NewDateTimeService(nil)

	helloworld.RegisterHelloWorldServiceServer(s, &helloService)
	helloworld.RegisterDateTimeServiceServer(s, &dtService)

	log.Printf("gRPC server listening at %v", listener.Addr())
	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
