//
// EPITECH PROJECT, 2024
// AREA
// File description:
// helloService
//

package api

import (
	"area/db"
	helloworld "area/protogen/gRPC/proto"
	"context"
	"encoding/json"
	"log"

	"google.golang.org/grpc"
)

type msg struct {
	Msg string `json:"msg"`
}

type HelloServiceClient struct {
	helloworld.HelloWorldServiceClient
}

func NewHelloServiceClient(conn *grpc.ClientConn) *HelloServiceClient {
	return &HelloServiceClient{helloworld.NewHelloWorldServiceClient(conn)}
}

func (hello *HelloServiceClient) SendAction(body []byte) (string, error) {
	msg := new(msg)
	err := json.Unmarshal([]byte(body), msg)

	r, err := hello.SayHello(context.Background(), &helloworld.HelloWorldRequest{Message: msg.Msg})

	if err != nil {
		return "", err
	}
	return r.GetMessage(), nil
}

////

type HelloService struct {
	db *db.UserDb
	helloworld.UnimplementedHelloWorldServiceServer
}

func NewHelloService(db *db.UserDb) HelloService {
	return HelloService{db: db}
}

func (hello *HelloService) SayHello(_ context.Context, req *helloworld.HelloWorldRequest) (*helloworld.HelloWorldResponse, error) {
	log.Println("In the service !")

	return &helloworld.HelloWorldResponse{Message: "Hello !"}, nil
}
