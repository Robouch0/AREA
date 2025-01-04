//
// EPITECH PROJECT, 2024
// AREA
// File description:
// trelloClientReaction
//

package trello_client

import (
	IServ "area/gRPC/api/serviceInterface"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
)

func (trelloCli *TrelloClient) createBoard(ingredients map[string]any, _ []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}

	var createReq gRPCService.CreateBoardReq
	err = json.Unmarshal(jsonString, &createReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err = trelloCli.cc.CreateBoard(ctx, &createReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Board created"}, nil
}
