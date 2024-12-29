//
// EPITECH PROJECT, 2024
// AREA
// File description:
// huggingFaceReactions
//

package huggingFace_server

import (
	hfType "area/gRPC/api/hugging_face/hfTypes"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	http_utils "area/utils/httpUtils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	createRepoURL           = "https://huggingface.co/api/repos/create"
	deleteRepoURL           = "https://huggingface.co/api/repos/delete"
	changeRepoVisibilityURL = "https://huggingface.co/api/repos/%s/%s/settings"
	moveRepoURL             = "https://huggingface.co/api/repos/move"
)

func (hfServ *HuggingFaceService) CreateRepository(ctx context.Context, req *gRPCService.CreateHFRepoReq) (*gRPCService.CreateHFRepoReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, hfServ.tokenDb, "HuggingFaceService", "hf")
	if err != nil {
		return nil, err
	}
	b, errBytes := json.Marshal(req)
	if errBytes != nil {
		return nil, status.Errorf(codes.DataLoss, "Invalid request sent for repo creation")
	}
	postRequest, err := http.NewRequest("POST", createRepoURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, status.Errorf(codes.Canceled, "Could not create the post request")
	}
	postRequest.Header = http_utils.GetDefaultBearerHTTPHeader(tokenInfo.AccessToken)

	_, err = http_utils.SendHttpRequest(postRequest, 200)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (hfServ *HuggingFaceService) DeleteRepository(ctx context.Context, req *gRPCService.DeleteHFRepoReq) (*gRPCService.DeleteHFRepoReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, hfServ.tokenDb, "HuggingFaceService", "hf")
	if err != nil {
		return nil, err
	}
	b, errBytes := json.Marshal(req)
	if errBytes != nil {
		return nil, status.Errorf(codes.DataLoss, "Invalid request sent for repo deletion")
	}
	deleteRequest, err := http.NewRequest("DELETE", deleteRepoURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, status.Errorf(codes.Canceled, "Could not create the delete request")
	}
	deleteRequest.Header = http_utils.GetDefaultBearerHTTPHeader(tokenInfo.AccessToken)

	_, err = http_utils.SendHttpRequest(deleteRequest, 204)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (hfServ *HuggingFaceService) ChangeRepoVisibility(ctx context.Context, req *gRPCService.ChangeHFRepoReq) (*gRPCService.ChangeHFRepoReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, hfServ.tokenDb, "HuggingFaceService", "hf")
	if err != nil {
		return nil, err
	}
	repoVisibility := hfType.RepoVisibility{
		Private: req.Private,
	}
	b, errBytes := json.Marshal(&repoVisibility)
	if errBytes != nil {
		return nil, status.Errorf(codes.DataLoss, "Invalid request sent for repo visibility")
	}

	putRequest, err := http.NewRequest("PUT", fmt.Sprintf(changeRepoVisibilityURL, req.Type, req.RepoId), bytes.NewBuffer(b))
	if err != nil {
		return nil, status.Errorf(codes.Canceled, "Could not create the put request")
	}

	putRequest.Header = http_utils.GetDefaultBearerHTTPHeader(tokenInfo.AccessToken)
	_, err = http_utils.SendHttpRequest(putRequest, 204)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (hfServ *HuggingFaceService) MoveRepo(ctx context.Context, req *gRPCService.MoveHFRepoReq) (*gRPCService.MoveHFRepoReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, hfServ.tokenDb, "HuggingFaceService", "hf")
	if err != nil {
		return nil, err
	}
	b, errBytes := json.Marshal(req)
	if errBytes != nil {
		return nil, status.Errorf(codes.DataLoss, "Invalid request sent for repo move")
	}
	postRequest, err := http.NewRequest("POST", moveRepoURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, status.Errorf(codes.Canceled, "Could not create the post request")
	}
	postRequest.Header = http_utils.GetDefaultBearerHTTPHeader(tokenInfo.AccessToken)

	_, err = http_utils.SendHttpRequest(postRequest, 200)
	if err != nil {
		return nil, err
	}
	return req, nil
}
