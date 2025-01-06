//
// EPITECH PROJECT, 2025
// AREA
// File description:
// miroServiceText
//

package miro_server

import (
	mirotypes "area/gRPC/api/miro/miroTypes"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"
)

const (
	newTextURL = "https://api.miro.com/v2/boards/%s/texts"
)

type CreateTextNoteReq struct {
	Content string `json:"content"`
}

func (miro *MiroService) CreateTextItem(ctx context.Context, req *gRPCService.CreateTextReq) (*gRPCService.CreateTextResp, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, miro.tokenDb, "MiroService", "miro")
	if err != nil {
		return nil, err
	}

	err = sendCreationReq[CreateTextNoteReq](tokenInfo, req.BoardId, newTextURL, mirotypes.MiroGenericBody[CreateTextNoteReq]{
		Data: CreateTextNoteReq{
			Content: req.Content,
		},
	})
	return &gRPCService.CreateTextResp{
		Id:      req.BoardId,
		Content: req.Content,
	}, err
}
