//
// EPITECH PROJECT, 2024
// AREA
// File description:
// helloService
//

package hello

import (
	"area/db"
	gRPCService "area/protogen/gRPC/proto"
	"context"
	"encoding/json"
	"log"

	"google.golang.org/grpc"
)

type HelloServiceClient struct {
	gRPCService.HelloWorldServiceClient
}

func NewHelloServiceClient(conn *grpc.ClientConn) *HelloServiceClient {
	return &HelloServiceClient{gRPCService.NewHelloWorldServiceClient(conn)}
}

func (hello *HelloServiceClient) SendAction(body []byte) (string, error) {
	msg := new(gRPCService.HelloWorldRequest)
	err := json.Unmarshal([]byte(body), msg)

	if err != nil {
		log.Println("Could not parse the body")
		return "", err
	}

	r, err := hello.SayHello(context.Background(), msg)
	if err != nil {
		log.Println("Could not send SayHello")
		return "", err
	}
	return r.GetMessage(), nil
}

////

type HelloService struct {
	db *db.UserDb
	gRPCService.UnimplementedHelloWorldServiceServer
}

func NewHelloService(db *db.UserDb) HelloService {
	return HelloService{db: db}
}

func (hello *HelloService) SayHello(_ context.Context, req *gRPCService.HelloWorldRequest) (*gRPCService.HelloWorldResponse, error) {
	log.Println("In the service !")

	return &gRPCService.HelloWorldResponse{Message: "Hello !"}, nil
}
