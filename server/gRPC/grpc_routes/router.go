//
// EPITECH PROJECT, 2024
// AREA
// File description:
// router
//

package grpc_routes

import (
	"area/gRPC/api"
	"area/gRPC/api/dateTime"
	"area/gRPC/api/hello"
	services "area/protogen/gRPC/proto"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func LaunchServices() {
	const addr = "0.0.0.0:50051"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	helloService := hello.NewHelloService(nil)
	dtService := dateTime.NewDateTimeService(nil)
	reactService := api.NewReactionService(nil)

	services.RegisterHelloWorldServiceServer(s, &helloService)
	services.RegisterDateTimeServiceServer(s, &dtService)
	services.RegisterReactionServiceServer(s, &reactService)

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		defer wg.Done()

		log.Printf("gRPC server listening at %v", listener.Addr())
		if err = s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return
	}

	reactService.InitServiceClients(conn)
	dtService.InitReactClient(conn)
	wg.Wait()
}
