//
// EPITECH PROJECT, 2024
// AREA
// File description:
// trelloService
//

package trello_server

import (
	"area/db"
	gRPCService "area/protogen/gRPC/proto"
)

type TrelloService struct {
	tokenDb      *db.TokenDb
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedTrelloServiceServer
}

func NewTrelloService() (*TrelloService, error) {
	tokenDb, err := db.InitTokenDb()

	return &TrelloService{tokenDb: tokenDb, reactService: nil}, err
}
