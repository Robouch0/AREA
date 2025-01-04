//
// EPITECH PROJECT, 2024
// AREA
// File description:
// trello Board service
//

package trello_server

import (
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (trello *TrelloService) CreateBoard(ctx context.Context, req *gRPCService.CreateBoardReq) (*gRPCService.CreateBoardResp, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, trello.tokenDb, "TrelloService", "trello")
	if err != nil {
		return nil, err
	}
	return &gRPCService.CreateBoardResp{}, nil
}
