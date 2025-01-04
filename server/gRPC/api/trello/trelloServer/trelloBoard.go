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
	"bytes"
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

	if req.BoardName == "" || req.BoardDescription == "" {
		return nil, errors.New("some required parameters are empty")
	}

	url := "https://api.trello.com/1/boards"

	postRequestBody := fmt.Sprintf(`{"name": "%s", "desc": "%s", "key": "%s", "token": "%s"}`,
		req.BoardName, req.BoardDescription, /* ou chopper l'api key trello ? dans l'env ? */, tokenInfo.AccessToken)

	postRequest, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(postRequestBody)))
	if err != nil {
		log.Println("Error when creating API call to Trello", err)
		return nil, err
	}

	postRequest.Header.Set("Content-Type", "application/json")
	cli := &http.Client{}
	resp, err := cli.Do(postRequest)
	if err != nil {
		log.Println("Error when sending API call to Trello", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("Error creating playlist:", resp.Status)
		return nil, errors.New(resp.Status)
	}
	log.Println("Playlist created")
	return &gRPCService.CreateBoardResp{}, nil
}
