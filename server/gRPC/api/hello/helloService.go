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
	"log"
)

type HelloService struct {
	db *db.UserDb
	gRPCService.UnimplementedHelloWorldServiceServer
}

func NewHelloService(db *db.UserDb) HelloService {
	return HelloService{db: db}
}

func (hello *HelloService) SayHello(_ context.Context, req *gRPCService.HelloWorldRequest) (*gRPCService.HelloWorldResponse, error) {
	log.Println("In the service !")

	return &gRPCService.HelloWorldResponse{Message: req.Message}, nil
}
