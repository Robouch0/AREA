//
// EPITECH PROJECT, 2025
// AREA
// File description:
// miroService
//

package miro_server

import (
	"area/db"
	gRPCService "area/protogen/gRPC/proto"
	"cmp"

	"google.golang.org/grpc"
)

type MiroService struct {
	tokenDb *db.TokenDb
	miroDb  *db.MiroDB
	// Database for miro webhooks
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedMiroServiceServer
}

func NewMiroService() (*MiroService, error) {
	tokenDb, errTok := db.InitTokenDb()
	miroDb, errMiro := db.InitMiroDb()
	if err := cmp.Or(errTok, errMiro); err != nil {
		return nil, err
	}
	miro := &MiroService{
		tokenDb:      tokenDb,
		miroDb:       miroDb,
		reactService: nil,
	}
	return miro, nil
}

func (miro *MiroService) InitReactClient(conn *grpc.ClientConn) {
	miro.reactService = gRPCService.NewReactionServiceClient(conn)
}
