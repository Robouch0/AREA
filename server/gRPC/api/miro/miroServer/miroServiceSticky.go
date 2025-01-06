//
// EPITECH PROJECT, 2025
// AREA
// File description:
// miroServiceSticky
//

package miro_server

import (
	mirotypes "area/gRPC/api/miro/miroTypes"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"
)

const (
	newStickyNoteURL = "https://api.miro.com/v2/boards/%s/sticky_notes"
)

type CreateStickyNoteReq struct {
	Content string `json:"content"`
	Shape   string `json:"shape"`
}

func (miro *MiroService) CreateStickyNote(ctx context.Context, req *gRPCService.CreateStickyNoteReq) (*gRPCService.CreateStickyNoteResp, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, miro.tokenDb, "MiroService", "miro")
	if err != nil {
		return nil, err
	}

	err = sendCreationReq[CreateStickyNoteReq](tokenInfo, req.BoardId, newStickyNoteURL, mirotypes.MiroGenericBody[CreateStickyNoteReq]{
		Data: CreateStickyNoteReq{
			Content: req.Content,
			Shape:   req.Shape,
		},
	})
	return &gRPCService.CreateStickyNoteResp{
		Id:      req.BoardId,
		Content: req.Content,
		Shape:   req.Shape,
	}, err
}
