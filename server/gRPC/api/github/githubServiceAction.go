//
// EPITECH PROJECT, 2024
// AREA [WSL: Ubuntu]
// File description:
// githubServiceAction
//

package github

import (
	githubtypes "area/gRPC/api/github/githubTypes"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"area/utils"
	grpcutils "area/utils/grpcUtils"
	http_utils "area/utils/httpUtils"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	TwebHookURL = "https://api.github.com/repos/%v/%v/hooks"
)

func (hfServ *GithubService) storeNewWebHook(
	tokenInfo *models.Token,
	req *gRPCService.GitWebHookInfo,
	repoAction string,
) error {
	_, err := hfServ.GithubDb.StoreNewGithub(&models.Github{
		ActionID:      uint(req.ActionId),
		UserID:        uint(tokenInfo.UserID),
		Activated:     true,
		RepoOwner:      req.Owner,
		RepoName:      req.Repo,
		RepoAction:    repoAction,
	})
	return err
}

func (git *GithubService) createWebHook(tokenInfo *models.Token, webhookReq *githubtypes.GitWebHookRequest, owner string, repo string) error {
	b, err := json.Marshal(webhookReq)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to convert the content to bytes"))
	}

	webHookURL := fmt.Sprintf(TwebHookURL, owner, repo)
	postRequest, err := http.NewRequest("POST", webHookURL, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	log.Println("Bearer " + tokenInfo.AccessToken)
	postRequest.Header.Set("authorization", "Bearer "+tokenInfo.AccessToken)
	postRequest.Header.Add("Accept", "application/vnd.github+json")
	_, err = http_utils.SendHttpRequest(postRequest, 200)
	if err != nil {
		return err
	}
	return nil
}

func (git *GithubService) CreatePushWebhook(ctx context.Context, req *gRPCService.GitWebHookInfo) (*gRPCService.GitWebHookInfo, error) {
	if req.Owner == "" || req.Repo == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid argument for webhook repo")
	}
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, git.tokenDb, "GithubService", "github")
	if err != nil {
		return nil, err
	}
	envWebhookUrl, err := utils.GetEnvParameter("WEBHOOK_URL")
	if err != nil {
		return nil, err
	}

	err = git.createWebHook(tokenInfo, &githubtypes.GitWebHookRequest{
		Event: []string{"push"},
		Config: githubtypes.GithubConfig{Url: fmt.Sprintf(envWebhookUrl, "github", req.ActionId), Content: "json" },
	}, req.Owner, req.Repo)
	if err != nil {
		return nil, err
	}
	if err := git.storeNewWebHook(tokenInfo, req, "push"); err != nil {
		return nil, err
	}
	return req, nil
}

func (git *GithubService) TriggerWebHook(ctx context.Context, req *gRPCService.GithubWebHookTriggerReq) (*gRPCService.GithubWebHookTriggerReq, error) {
	act, err := git.GithubDb.GetGithubByActionID(uint(req.ActionId))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "No such action with id %d", req.ActionId)
	}

	// var hfPayload hfType.HFWebHookResponse
	// if json.Unmarshal(req.Payload, &hfPayload) != nil {
	// 	return nil, status.Errorf(codes.InvalidArgument, "Invalid Payload received")
	// }
	// if hfPayload.Event.Action == act.RepoAction && hfPayload.Event.Scope == act.RepoScope {
	// 	if hfPayload.Event.Action == "discussion" && hfPayload.Discussion.IsPullRequest != act.IsPullRequest {
	// 		return nil, status.Errorf(codes.InvalidArgument, "Received Discussion event with incorrect value for IsPullRequest")
	// 	}
	reactCtx := grpcutils.CreateContextFromUserID(int(act.UserID))
	_, err = git.reactService.LaunchReaction(
		reactCtx,
		&gRPCService.LaunchRequest{ActionId: int64(act.ActionID), PrevOutput: req.Payload},
	)
	if err != nil {
			return nil, status.Errorf(codes.Internal, "Could not handle action's reaction")
		}
	// }
	return req, nil
}
