//
// EPITECH PROJECT, 2024
// AREA
// File description:
// huggingFaceAction
//

package huggingFace

import (
	hfType "area/gRPC/api/hugging_face/hfTypes"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"area/utils"
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
	webHookURL = "https://huggingface.co/api/settings/webhooks"
)

func (hfServ *HuggingFaceService) storeNewWebHook(
	tokenInfo *models.Token,
	req *gRPCService.HFWebHookInfo,
	repoAction,
	repoScope string,
	isPR bool,
) error {
	_, err := hfServ.hfDb.StoreNewHF(&models.HuggingFace{
		ActionID:      uint(req.ActionId),
		UserID:        uint(tokenInfo.UserID),
		Activated:     true,
		RepoType:      req.Type,
		RepoName:      req.Name,
		RepoAction:    repoAction,
		RepoScope:     repoScope, // Maybe add few scope possible as response
		IsPullRequest: isPR,
	})
	return err
}

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
	_, err = http_utils.SendHttpRequest(postRequest, 200)
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
	if err := hfServ.storeNewWebHook(tokenInfo, req, "update", "repo.content", false); err != nil {
		return nil, err
	}
	return req, nil
}

func (hfServ *HuggingFaceService) CreateNewPRWebHook(ctx context.Context, req *gRPCService.HFWebHookInfo) (*gRPCService.HFWebHookInfo, error) {
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
		Domains: []string{"discussion"},
	})
	if err != nil {
		return nil, err
	}
	if err := hfServ.storeNewWebHook(tokenInfo, req, "create", "discussion", true); err != nil {
		return nil, err
	}
	return req, nil
}

func (hfServ *HuggingFaceService) CreateNewDiscussionWebHook(ctx context.Context, req *gRPCService.HFWebHookInfo) (*gRPCService.HFWebHookInfo, error) {
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
		Domains: []string{"discussion"},
	})
	if err != nil {
		return nil, err
	}
	if err := hfServ.storeNewWebHook(tokenInfo, req, "create", "discussion", false); err != nil {
		return nil, err
	}
	return req, nil
}

/* WebHook Trigger Functions */

func (hfServ *HuggingFaceService) TriggerWebHook(ctx context.Context, req *gRPCService.WebHookTriggerReq) (*gRPCService.WebHookTriggerReq, error) {
	act, err := hfServ.hfDb.GetHfByActionID(uint(req.ActionId))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "No such action with id %d", req.ActionId)
	}

	var hfPayload hfType.HFWebHookResponse
	if json.Unmarshal(req.Payload, &hfPayload) != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Payload received")
	}
	if hfPayload.Event.Action == act.RepoAction && hfPayload.Event.Scope == act.RepoScope {
		if hfPayload.Event.Action == "discussion" && hfPayload.Discussion.IsPullRequest != act.IsPullRequest {
			return nil, status.Errorf(codes.InvalidArgument, "Received Discussion event with incorrect value for IsPullRequest")
		}
		reactCtx := grpcutils.CreateContextFromUserID(int(act.UserID))
		_, err := hfServ.reactService.LaunchReaction(
			reactCtx,
			&gRPCService.LaunchRequest{ActionId: int64(act.ActionID), PrevOutput: req.Payload},
		)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Could not handle action's reaction")
		}
	}
	return req, nil
}
