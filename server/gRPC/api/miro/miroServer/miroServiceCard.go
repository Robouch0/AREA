//
// EPITECH PROJECT, 2025
// AREA
// File description:
// miroServiceCard
//

package miro_server

import (
	mirotypes "area/gRPC/api/miro/miroTypes"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"
)

const (
	newCardURL = "https://api.miro.com/v2/boards/%s/cards"
)

type CreateCardNoteReq struct {
	Description string `json:"description"`
	Title       string `json:"title"`
}

func (miro *MiroService) CreateCardItem(ctx context.Context, req *gRPCService.CreateCardReq) (*gRPCService.CreateCardResp, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, miro.tokenDb, "MiroService", "miro")
	if err != nil {
		return nil, err
	}

	err = sendCreationReq[CreateCardNoteReq](tokenInfo, req.BoardId, newCardURL, mirotypes.MiroGenericBody[CreateCardNoteReq]{
		Data: CreateCardNoteReq{
			Description: req.Description,
			Title:       req.Title,
		},
	})
	return &gRPCService.CreateCardResp{
		Id:          req.BoardId,
		Description: req.Description,
		Title:       req.Title,
	}, err
}
