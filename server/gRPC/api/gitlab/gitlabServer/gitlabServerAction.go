//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// gitlabServerAction
//

package gitlab_server

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
	"net/http"

	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	pushWebHookURL = "https://api.github.com/repos/%v/%v/hooks"
	updateWebHookURL = "https://api.github.com/repos/%v/%v/hooks/%v"
)


func (git *GitlabService) storeNewWebHook(
	tokenInfo *models.Token,
	req *gRPCService.GitWebHookInfo,
	repoAction string,
) error {
	_, err := git.GithubDb.StoreNewGithub(&models.Github{
		ActionID:   uint(req.ActionId),
		UserID:     uint(tokenInfo.UserID),
		Activated:  true,
		RepoOwner:  req.Owner,
		RepoName:   req.Repo,
		RepoAction: repoAction,
	})
	return err
}

func (git *GitlabService) createWebHook(tokenInfo *models.Token, webhookReq *githubtypes.GitWebHookRequest, owner string, repo string) error {
	b, err := json.Marshal(webhookReq)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to convert the content to bytes"))
	}

	webHookURL := fmt.Sprintf(pushWebHookURL, owner, repo)
	postRequest, err := http.NewRequest("POST", webHookURL, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	postRequest.Header.Set("Authorization", "Bearer "+tokenInfo.AccessToken)
	postRequest.Header.Add("Accept", "application/vnd.github+json")
	_, err = http_utils.SendHttpRequest(postRequest, 201)
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
		Event:  []string{"push"},
		Config: githubtypes.GithubConfig{Url: fmt.Sprintf(envWebhookUrl, "github", "push", req.ActionId), Content: "json"},
	}, req.Owner, req.Repo)
	if err != nil {
		return nil, err
	}
	if err := git.storeNewWebHook(tokenInfo, req, "push"); err != nil {
		return nil, err
	}
	return req, nil
}

func (git *GithubService) UpdatePushWebhook(ctx context.Context, action *models.Github, activated bool) (error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, git.tokenDb, "GithubService", "github")
	if err != nil {
		return err
	}
	envWebhookUrl, err := utils.GetEnvParameter("WEBHOOK_URL")
	if err != nil {
		return err
	}

	b, err := json.Marshal(&githubtypes.GitWebHookRequest{
		Event:  []string{action.RepoAction},
		Active: activated,
		Config: githubtypes.GithubConfig{Url: fmt.Sprintf(envWebhookUrl, "github", action.RepoAction, action.ActionID), Content: "json"}})
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to convert the content to bytes"))
	}

	webHookURL := fmt.Sprintf(updateWebHookURL, action.RepoOwner, action.RepoName)
	postRequest, err := http.NewRequest("PATCH", webHookURL, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	postRequest.Header.Set("Authorization", "Bearer " + tokenInfo.AccessToken)
	postRequest.Header.Add("Accept", "application/vnd.github+json")
	_, err = http_utils.SendHttpRequest(postRequest, 200)
	if err != nil {
		return err
	}
	return nil
}

func (git *GithubService) TriggerWebHook(ctx context.Context, req *gRPCService.GithubWebHookTriggerReq) (*gRPCService.GithubWebHookTriggerReq, error) {
	act, err := git.GithubDb.GetGithubByActionID(uint(req.ActionId))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "No such action with id %d", req.ActionId)
	}

	var gitpayload githubtypes.GithubEvents
	if json.Unmarshal(req.Payload, &gitpayload) != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Payload received")
	}
	if len(gitpayload.Hook.Events) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "No events in github webhook payload")
	}
	if gitpayload.Hook.Events[0] == act.RepoAction {
		reactCtx := grpcutils.CreateContextFromUserID(int(act.UserID))
		_, err = git.reactService.LaunchReaction(
			reactCtx,
			&gRPCService.LaunchRequest{ActionId: int64(act.ActionID), PrevOutput: nil},
		)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Could not handle action's reaction")
		}
	}
	return req, nil
}

func (github *GithubService) SetActivateAction(ctx context.Context, req *gRPCService.SetActivateGithub) (*gRPCService.SetActivateGithub, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "github")
	if err != nil {
		return nil, err
	}
	action, err := github.GithubDb.GetGithubByActionID(uint(req.ActionId))
	if err != nil {
		return nil, err
	}
	if !req.Activated {
		err = github.UpdatePushWebhook(ctx, action, false)
	} else {
		err = github.UpdatePushWebhook(ctx, action, true)
	}
	if err != nil {
		return nil, err
	}
	_, err = github.GithubDb.SetActivateByActionID(req.Activated, userID, uint(req.ActionId))
	if err != nil {
		return nil, err
	}
	return req, nil
}
