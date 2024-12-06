//
// EPITECH PROJECT, 2024
// AREA
// File description:
// helloClient
//

package hello

import (
	IServ "area/gRPC/api/serviceInterface"
	gRPCService "area/protogen/gRPC/proto"
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"
)

type HelloServiceClient struct {
	gRPCService.HelloWorldServiceClient
}

func NewHelloServiceClient(conn *grpc.ClientConn) *HelloServiceClient {
	return &HelloServiceClient{gRPCService.NewHelloWorldServiceClient(conn)}
}

func (hello *HelloServiceClient) SendAction(body map[string]any, actionID int) (*IServ.ActionResponseStatus, error) {
	if msg, ok := body["msg"]; ok {
		_, err := hello.SayHello(context.Background(), &gRPCService.HelloWorldRequest{Message: msg.(string)})
		if err != nil {
			log.Println("Could not send SayHello")
			return nil, err
		}
		return &IServ.ActionResponseStatus{Description: "Hello !", ActionID: actionID}, nil
	}
	return nil, errors.New("Incorrect body with no msg")
}
