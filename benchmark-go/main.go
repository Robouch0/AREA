//
// EPITECH PROJECT, 2024
// AREA
// File description:
// main
//

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// <MOD-NAME>/<PACKAGE-NAME>
	pb "area/greeting"

	"google.golang.org/grpc"
)

type server struct {
	pb.GreetingServiceServer
}

func (s *server) Greeting(ctx context.Context, req *pb.GreetingServiceRequest) (*pb.GreetingServiceReply, error) {
	fmt.Printf("I receive a message \"%s\"\n", req.Name)
	return &pb.GreetingServiceReply{
		Message: fmt.Sprintf("Hello, %s", req.Name),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetingServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// New service + DB
// Try to call one service from another
