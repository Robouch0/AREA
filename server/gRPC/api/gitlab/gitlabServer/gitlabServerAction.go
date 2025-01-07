//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// gitlabServerAction
//

package gitlab_server

import (
	gitlabtypes "area/gRPC/api/gitlab/gitlabTypes"
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
	pushWebHookURL   = "https://www.gitlab.com/api/v4/projects/%v/hooks"
	updateWebHookURL = "https://www.gitlab.com/api/v4/projects/%v/hooks/%v"
)

func (git *GitlabService) storeNewWebHook(
	tokenInfo *models.Token,
	req *gRPCService.GitlabWebHookInfo,
	repoAction string,
) error {
	_, err := git.gitlabDb.StoreNewGithub(&models.Gitlab{
		ActionID:   uint(req.ActionId),
		UserID:     uint(tokenInfo.UserID),
		Activated:  true,
		RepoId:     req.Id,
		RepoAction: repoAction,
	})
	return err
}

func (git *GitlabService) createWebHook(tokenInfo *models.Token, webhookReq *gitlabtypes.GitlabWebHookRequest, project_id string) error {
	b, err := json.Marshal(webhookReq)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to convert the content to bytes"))
	}

	webHookURL := fmt.Sprintf(pushWebHookURL, project_id)
	postRequest, err := http.NewRequest("POST", webHookURL, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	postRequest.Header.Set("Authorization", "Bearer " + tokenInfo.AccessToken)
	q := postRequest.URL.Query()
	q.Set("access_token", tokenInfo.AccessToken)
	postRequest.URL.RawQuery = q.Encode()
	_, err = http_utils.SendHttpRequest(postRequest, 201)
	if err != nil {
		return err
	}
	return nil
}

func (git *GitlabService) CreatePushWebhook(ctx context.Context, req *gRPCService.GitlabWebHookInfo) (*gRPCService.GitlabWebHookInfo, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid argument for webhook repo")
	}
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, git.tokenDb, "GitlabService", "gitlab")
	if err != nil {
		return nil, err
	}
	envWebhookUrl, err := utils.GetEnvParameter("WEBHOOK_URL")
	if err != nil {
		return nil, err
	}

	err = git.createWebHook(tokenInfo, &gitlabtypes.GitlabWebHookRequest{
		Url: fmt.Sprintf(envWebhookUrl, "github", "push", req.ActionId),
		PushEvent: true,
	}, req.Id)
	if err != nil {
		return nil, err
	}
	if err := git.storeNewWebHook(tokenInfo, req, "push"); err != nil {
		return nil, err
	}
	return req, nil
}

func (git *GitlabService) TriggerWebHook(ctx context.Context, req *gRPCService.GitlabWebHookTriggerReq) (*gRPCService.GitlabWebHookTriggerReq, error) {
	act, err := git.gitlabDb.GetGithubByActionID(uint(req.ActionId))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "No such action with id %d", req.ActionId)
	}

	log.Println(req.Payload)
	// var gitpayload GitlabService.GithubEvents
	// if json.Unmarshal(req.Payload, &gitpayload) != nil {
	// 	return nil, status.Errorf(codes.InvalidArgument, "Invalid Payload received")
	// }
	// if len(gitpayload.Hook.Events) == 0 {
	// 	return nil, status.Errorf(codes.InvalidArgument, "No events in github webhook payload")
	// }
	// if gitpayload.Hook.Events[0] == act.RepoAction {
		reactCtx := grpcutils.CreateContextFromUserID(int(act.UserID))
		_, err = git.reactService.LaunchReaction(
			reactCtx,
			&gRPCService.LaunchRequest{ActionId: int64(act.ActionID), PrevOutput: nil},
		)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Could not handle action's reaction")
		}
	// }
	return req, nil
}
