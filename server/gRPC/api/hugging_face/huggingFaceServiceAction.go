//
// EPITECH PROJECT, 2024
// AREA
// File description:
// huggingFaceAction
//

package huggingFace

import (
	hfType "area/gRPC/api/hugging_face/HFTypes"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"area/utils"
	grpcutils "area/utils/grpcUtils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
 * Actions:
 * - If someone push reaction (action: update, scope: repo.content)
 * - If some create a new discussion (action: create, scope: discussion, discussion.isPullRequest: false)
 * - If some create a new PR (action: create, scope: discussion, discussion.isPullRequest: true)
 */
const (
	webHookURL = "https://huggingface.co/api/settings/webhooks"
)

func (hfServ *HuggingFaceService) createWebHook(tokenInfo *models.Token, webhookReq *hfType.HFWebHookRequest) error {
	b, err := json.Marshal(webhookReq)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to convert the content to bytes"))
	}
	postRequest, err := http.NewRequest("POST", webHookURL, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	postRequest.Header.Set("Authorization", "Bearer "+tokenInfo.AccessToken)
	postRequest.Header.Add("Accept", "application/json")
	_, err = utils.SendHttpRequest(postRequest, 200)
	if err != nil {
		return err
	}
	return nil
}

func (hfServ *HuggingFaceService) CreateRepoUpdateWebHook(ctx context.Context, req *gRPCService.HFWebHookInfo) (*gRPCService.HFWebHookInfo, error) {
	if req.Name == "" || req.Type == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid argument for webhook repo")
	}
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, hfServ.tokenDb, "HuggingFaceService", "hf")
	if err != nil {
		return nil, err
	}
	envWebhookUrl, err := utils.GetEnvParameter("WEBHOOK_URL")
	if err != nil {
		return nil, err
	}

	err = hfServ.createWebHook(tokenInfo, &hfType.HFWebHookRequest{
		Watched: []hfType.HFRepo{{Type: req.Type, Name: req.Name}},
		Url:     fmt.Sprintf(envWebhookUrl, "hf", req.ActionId),
		Domains: []string{"repo"},
	})
	if err != nil {
		return nil, err
	}
	// Put to database
	return req, nil
}

func (hfServ *HuggingFaceService) CreateNewPRWebHook(ctx context.Context, req *gRPCService.HFWebHookInfo) (*gRPCService.HFWebHookInfo, error) {
	if req.Name == "" || req.Type == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid argument for webhook repo")
	}
	_, err := grpcutils.GetTokenByContext(ctx, hfServ.tokenDb, "HuggingFaceService", "hf")
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (hfServ *HuggingFaceService) CreateNewDiscussionWebHook(ctx context.Context, req *gRPCService.HFWebHookInfo) (*gRPCService.HFWebHookInfo, error) {
	if req.Name == "" || req.Type == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid argument for webhook repo")
	}
	_, err := grpcutils.GetTokenByContext(ctx, hfServ.tokenDb, "HuggingFaceService", "hf")
	if err != nil {
		return nil, err
	}
	return req, nil
}
