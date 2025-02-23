//
// EPITECH PROJECT, 2024
// AREA
// File description:
// huggingFaceService
//

package huggingFace_server

import (
	"area/db"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"

	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"google.golang.org/grpc"
)

const (
	googleModelURL = "https://api-inference.huggingface.co/models/google/gemma-2-2b-it"
)

type HuggingFaceService struct {
	tokenDb      *db.TokenDb
	hfDb         *db.HuggingFaceDB
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedHuggingFaceServiceServer
}

func NewHuggingFaceService() (*HuggingFaceService, error) {
	tokenDb, err := db.InitTokenDb()
	if err != nil {
		return nil, err
	}
	hfDb, err := db.InitHuggingFaceDb()
	if err != nil {
		return nil, err
	}

	return &HuggingFaceService{tokenDb: tokenDb, hfDb: hfDb, reactService: nil}, nil
}

func (hf *HuggingFaceService) InitReactClient(conn *grpc.ClientConn) {
	hf.reactService = gRPCService.NewReactionServiceClient(conn)
}

func (hfServ *HuggingFaceService) LaunchTextGeneration(ctx context.Context, req *gRPCService.TextGenerationReq) (*gRPCService.TextGenerationRes, error) {
	userID, errClaim := grpcutils.GetUserIdFromContext(ctx, "HFService")
	if errClaim != nil {
		return nil, errClaim
	}

	tokenInfo, err := hfServ.tokenDb.GetUserTokenByProvider(int64(userID), "hf")
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	postRequest, err := http.NewRequest("POST", googleModelURL, bytes.NewBuffer(b))
	postRequest.Header.Set("Authorization", tokenInfo.AccessToken)
	postRequest.Header.Add("Content-Type", "application/json;charset=UTF-8")
	postRequest.Header.Add("Accept", "application/json")

	cli := &http.Client{}
	resp, err := cli.Do(postRequest)
	if err != nil {
		return nil, err
	}
	if resp.Status != "200 OK" {
		return nil, errors.New(resp.Status)
	}
	textGenRes := &[]gRPCService.TextGenerationRes{}
	if err = json.NewDecoder(resp.Body).Decode(textGenRes); err != nil {
		return nil, err
	}
	if len(*textGenRes) == 0 {
		return nil, errors.New("No generated text received")
	}
	return &(*textGenRes)[0], nil
}
