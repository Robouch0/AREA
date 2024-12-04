//
// EPITECH PROJECT, 2024
// AREA
// File description:
// helloService
//

package api

import (
	"area/db"
	gRPCService "area/protogen/gRPC/proto"
	"context"
	"encoding/json"
	"log"

	"google.golang.org/grpc"
)

type msg struct {
	Msg string `json:"msg"`
}

type HelloServiceClient struct {
	gRPCService.HelloWorldServiceClient
}

func NewHelloServiceClient(conn *grpc.ClientConn) *HelloServiceClient {
	return &HelloServiceClient{gRPCService.NewHelloWorldServiceClient(conn)}
}

func (hello *HelloServiceClient) SendAction(body []byte) (string, error) {
	msg := new(msg)
	err := json.Unmarshal([]byte(body), msg)

	r, err := hello.SayHello(context.Background(), &gRPCService.HelloWorldRequest{Message: msg.Msg})

	if err != nil {
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
