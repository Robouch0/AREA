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
	"io"
	"net/http"
	"strconv"

	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	pushWebHookURL   = "https://www.gitlab.com/api/v4/projects/%v/hooks"
	updateWebHookURL = "https://www.gitlab.com/api/v4/projects/%v/hooks/%v"
)

type GitLabWebhookPayload struct {
	Hookid uint32 `json:"id"`
}

func (git *GitlabService) storeNewWebHook(
	tokenInfo *models.Token,
	req *gRPCService.GitlabWebHookInfo,
	repoAction models.GlAction,
	actionType models.GlType,
	hookid uint32,
) error {
	_, err := git.gitlabDb.StoreNewGitlab(&models.Gitlab{
		ActionID:   uint(req.ActionId),
		UserID:     uint(tokenInfo.UserID),
		Activated:  true,
		RepoID:     req.Id,
		RepoAction: repoAction,
		ActionType: actionType,
		HookID: uint(hookid),
	})
	return err
}

func (git *GitlabService) createWebHook(tokenInfo *models.Token, webhookReq *gitlabtypes.GitlabWebHookRequest, project_id string) (*GitLabWebhookPayload, error) {
	b, err := json.Marshal(webhookReq)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to convert the content to bytes"))
	}

	webHookURL := fmt.Sprintf(pushWebHookURL, project_id)
	postRequest, err := http.NewRequest("POST", webHookURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	postRequest.Header.Add("Content-Type", "application/json;charset=UTF-8")
	q := postRequest.URL.Query()
	q.Set("access_token", tokenInfo.AccessToken)
	postRequest.URL.RawQuery = q.Encode()
	resp, err := http_utils.SendHttpRequest(postRequest, 201)
	if err != nil {
		return nil, err
	}
	var payload GitLabWebhookPayload
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error : reading body %s", err)
	}
	err = json.Unmarshal(bytes, &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
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

	resp, err := git.createWebHook(tokenInfo, &gitlabtypes.GitlabWebHookRequest{
		Url:       fmt.Sprintf(envWebhookUrl, "gitlab", models.GlPush, req.ActionId),
		PushEvent: true,
	}, req.Id)
	if err != nil {
		return nil, err
	}
	if err := git.storeNewWebHook(tokenInfo, req, models.GlPush, models.GlEmpty, resp.Hookid); err != nil {
		return nil, err
	}
	return req, nil
}

func (git *GitlabService) CreateIssueWebhook(ctx context.Context, req *gRPCService.GitlabWebHookInfo) (*gRPCService.GitlabWebHookInfo, error) {
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

	resp, err := git.createWebHook(tokenInfo, &gitlabtypes.GitlabWebHookRequest{
		Url:         fmt.Sprintf(envWebhookUrl, "gitlab", models.GlIssue, req.ActionId),
		IssuesEvent: true,
	}, req.Id)
	if err != nil {
		return nil, err
	}
	if err := git.storeNewWebHook(tokenInfo, req, models.GlIssue, models.Glopen, resp.Hookid); err != nil {
		return nil, err
	}
	return req, nil
}

func (git *GitlabService) CreateTagWebhook(ctx context.Context, req *gRPCService.GitlabWebHookInfo) (*gRPCService.GitlabWebHookInfo, error) {
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

	resp, err := git.createWebHook(tokenInfo, &gitlabtypes.GitlabWebHookRequest{
		Url:      fmt.Sprintf(envWebhookUrl, "gitlab", models.GlTag, req.ActionId),
		TagEvent: true,
	}, req.Id)
	if err != nil {
		return nil, err
	}
	if err := git.storeNewWebHook(tokenInfo, req, models.GlTag, models.GlEmpty, resp.Hookid); err != nil {
		return nil, err
	}
	return req, nil
}

func (git *GitlabService) CreateReleaseWebhook(ctx context.Context, req *gRPCService.GitlabWebHookInfo) (*gRPCService.GitlabWebHookInfo, error) {
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

	resp, err := git.createWebHook(tokenInfo, &gitlabtypes.GitlabWebHookRequest{
		Url:          fmt.Sprintf(envWebhookUrl, "gitlab", models.GlRelease, req.ActionId),
		ReleaseEvent: true,
	}, req.Id)
	if err != nil {
		return nil, err
	}
	if err := git.storeNewWebHook(tokenInfo, req, models.GlRelease, models.Glcreate, resp.Hookid); err != nil {
		return nil, err
	}
	return req, nil
}

func (git *GitlabService) CreateMergeEventWebhook(ctx context.Context, req *gRPCService.GitlabWebHookInfo) (*gRPCService.GitlabWebHookInfo, error) {
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

	resp, err := git.createWebHook(tokenInfo, &gitlabtypes.GitlabWebHookRequest{
		Url:        fmt.Sprintf(envWebhookUrl, "gitlab", models.GlMergeR, req.ActionId),
		MergeEvent: true,
	}, req.Id)
	if err != nil {
		return nil, err
	}
	if err := git.storeNewWebHook(tokenInfo, req, models.GlMergeR, models.Glopen, resp.Hookid); err != nil {
		return nil, err
	}
	return req, nil
}

func (git *GitlabService) TriggerWebHook(ctx context.Context, req *gRPCService.GitlabWebHookTriggerReq) (*gRPCService.GitlabWebHookTriggerReq, error) {
	act, err := git.gitlabDb.GetGitlabByActionID(uint(req.ActionId))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "No such action with id %d", req.ActionId)
	}

	var gitpayload gitlabtypes.GitlabWebHookResponse
	if json.Unmarshal(req.Payload, &gitpayload) != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Payload received")
	}
	if (gitpayload.EventName != "" && gitpayload.EventName == string(act.RepoAction) && strconv.Itoa(gitpayload.Project.ProjectId) == act.RepoID) ||
		(gitpayload.EventType != "" && gitpayload.EventType == string(act.RepoAction) && strconv.Itoa(gitpayload.Project.ProjectId) == act.RepoID) ||
		(gitpayload.ObjectKind != "" && gitpayload.ObjectKind == string(act.RepoAction) && strconv.Itoa(gitpayload.Project.ProjectId) == act.RepoID) {
		if gitpayload.Object.Action == string(act.ActionType) || gitpayload.Action == string(act.ActionType) {
			gitoutput := &gitlabtypes.GitlabPayloadResponse{
				EventName: gitpayload.EventName + gitpayload.EventType + gitpayload.ObjectKind,
				ProjectId: gitpayload.Project.ProjectId,
			}
			b, err := json.Marshal(gitoutput)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Could not handle action's reaction")
			}
			reactCtx := grpcutils.CreateContextFromUserID(int(act.UserID))
			_, err = git.reactService.LaunchReaction(
				reactCtx,
				&gRPCService.LaunchRequest{ActionId: int64(act.ActionID), PrevOutput: b},
			)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Could not handle action's reaction")
			}
		}
	}
	return req, nil
}
