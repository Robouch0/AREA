//
// EPITECH PROJECT, 2024
// AREA
// File description:
// trelloService
//

package asana_server

import (
	"area/db"
	gRPCService "area/protogen/gRPC/proto"
)

type AsanaService struct {
	tokenDb      *db.TokenDb
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedAsanaServiceServer
}

func NewAsanaService() (*AsanaService, error) {
	tokenDb, err := db.InitTokenDb()

	return &AsanaService{tokenDb: tokenDb, reactService: nil}, err
}
