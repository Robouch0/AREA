//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleService
//

package google_server

import (
	"area/db"
	gRPCService "area/protogen/gRPC/proto"
)

const (
	sendMessageMeURL = "https://gmail.googleapis.com/gmail/v1/users/me/messages/send"
)

type GoogleService struct {
	tokenDb      *db.TokenDb
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedGoogleServiceServer
}

func NewGoogleService() (*GoogleService, error) {
	tokenDb, err := db.InitTokenDb()

	return &GoogleService{tokenDb: tokenDb, reactService: nil}, err
}
