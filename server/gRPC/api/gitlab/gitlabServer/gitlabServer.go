//
// EPITECH PROJECT, 2024
// AREA
// File description:
// gitlabServer
//

package gitlab_server

import (
	"area/db"
	gitlabtypes "area/gRPC/api/gitlab/gitlabTypes"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"area/utils"
	grpcutils "area/utils/grpcUtils"
	http_utils "area/utils/httpUtils"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"google.golang.org/grpc"
)

type GitlabService struct {
	tokenDb      *db.TokenDb
	gitlabDb     *db.GitlabDB
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedGitlabServiceServer
}

const (
	deleteGitlabWebHookURL = "https://www.gitlab.com/api/v4/projects/%v/hooks/%v"
)

func NewGitlabService() (*GitlabService, error) {
	tokenDb, err := db.InitTokenDb()
	if err != nil {
		return nil, err
	}

	gitlabDb, err := db.InitGitlabDb()

	return &GitlabService{tokenDb: tokenDb, gitlabDb: gitlabDb, reactService: nil}, err
}

func (git *GitlabService) InitReactClient(conn *grpc.ClientConn) {
	git.reactService = gRPCService.NewReactionServiceClient(conn)
}

func (git *GitlabService) CreateFile(ctx context.Context, req *gRPCService.CreateLabRepoFileReq) (*gRPCService.CreateLabRepoFileReq, error) {
	if req.Branch == "" || req.CommitMessage == "" || req.FilePath == "" || req.Id == "" || req.Content == "" {
		return nil, errors.New("Some required parameters are empty")
	}

	url := fmt.Sprintf("https://www.gitlab.com/api/v4/projects/%v/repository/files/%v", req.Id, req.FilePath)
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, git.tokenDb, "GitlabService", "gitlab")
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	pathRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	pathRequest.Header.Add("Content-Type", "application/json;charset=UTF-8")
	q := pathRequest.URL.Query()
	q.Set("access_token", tokenInfo.AccessToken)
	pathRequest.URL.RawQuery = q.Encode()

	resp, err := http_utils.SendHttpRequest(pathRequest, 201)
	if err != nil {
		return nil, err
	}
	log.Println(resp.Body)
	return req, nil
}

func (git *GitlabService) UpdateFile(ctx context.Context, req *gRPCService.UpdateLabRepoFileReq) (*gRPCService.UpdateLabRepoFileReq, error) {
	if req.Branch == "" || req.CommitMessage == "" || req.FilePath == "" || req.Id == "" || req.Content == "" {
		return nil, errors.New("Some required parameters are empty")
	}

	url := fmt.Sprintf("https://www.gitlab.com/api/v4/projects/%v/repository/files/%v", req.Id, req.FilePath)
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, git.tokenDb, "GitlabService", "gitlab")
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	pathRequest, err := http.NewRequest("PUT", url, bytes.NewBuffer(b))
	pathRequest.Header.Add("Content-Type", "application/json;charset=UTF-8")
	q := pathRequest.URL.Query()
	q.Set("access_token", tokenInfo.AccessToken)
	pathRequest.URL.RawQuery = q.Encode()

	resp, err := http_utils.SendHttpRequest(pathRequest, 200)
	if err != nil {
		return nil, err
	}
	log.Println(resp.Body)
	return req, nil
}

func (git *GitlabService) DeleteFile(ctx context.Context, req *gRPCService.DeleteLabRepoFileReq) (*gRPCService.DeleteLabRepoFileReq, error) {
	if req.Branch == "" || req.CommitMessage == "" || req.FilePath == "" || req.Id == "" {
		return nil, errors.New("Some required parameters are empty")
	}

	url := fmt.Sprintf("https://www.gitlab.com/api/v4/projects/%v/repository/files/%v", req.Id, req.FilePath)
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, git.tokenDb, "GitlabService", "gitlab")
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	pathRequest, err := http.NewRequest("DELETE", url, bytes.NewBuffer(b))
	pathRequest.Header.Add("Content-Type", "application/json;charset=UTF-8")
	q := pathRequest.URL.Query()
	q.Set("access_token", tokenInfo.AccessToken)
	pathRequest.URL.RawQuery = q.Encode()

	resp, err := http_utils.SendHttpRequest(pathRequest, 204)
	if err != nil {
		return nil, err
	}
	log.Println(resp.Body)
	return req, nil
}

func (git *GitlabService) MarkItemAsDone(ctx context.Context, req *gRPCService.TodoLabItemDoneReq) (*gRPCService.TodoLabItemDoneReq, error) {
	if req.Id == "" {
		return nil, errors.New("Some required parameters are empty")
	}

	url := fmt.Sprintf("https://www.gitlab.com/api/v4/todos/%v/mark_as_done", req.Id)
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, git.tokenDb, "GitlabService", "gitlab")
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	pathRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	pathRequest.Header.Add("Content-Type", "application/json;charset=UTF-8")
	q := pathRequest.URL.Query()
	q.Set("access_token", tokenInfo.AccessToken)
	pathRequest.URL.RawQuery = q.Encode()

	resp, err := http_utils.SendHttpRequest(pathRequest, 201)
	if err != nil {
		return nil, err
	}
	log.Println(resp.Body)
	return req, nil
}

func (git *GitlabService) MarkAllItemAsDone(ctx context.Context, req *gRPCService.AllTodoLabItemDoneReq) (*gRPCService.AllTodoLabItemDoneReq, error) {
	url := fmt.Sprintf("https://www.gitlab.com/api/v4/todos/mark_as_done")
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, git.tokenDb, "GitlabService", "gitlab")
	if err != nil {
		return nil, err
	}

	pathRequest, err := http.NewRequest("POST", url, nil)
	pathRequest.Header.Add("Content-Type", "application/json;charset=UTF-8")
	q := pathRequest.URL.Query()
	q.Set("access_token", tokenInfo.AccessToken)
	pathRequest.URL.RawQuery = q.Encode()

	_, err = http_utils.SendHttpRequest(pathRequest, 204)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (git *GitlabService) SetActivateAction(ctx context.Context, req *gRPCService.SetActivateGitlab) (*gRPCService.SetActivateGitlab, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, git.tokenDb, "GitlabService", "gitlab")
	if err != nil {
		return nil, err
	}
	action, err := git.gitlabDb.GetGitlabByActionID(uint(req.ActionId))
	if err != nil {
		return nil, err
	}
	if !req.Activated {
		id := strconv.Itoa(int(action.HookID))
		webHookURL := fmt.Sprintf(deleteGitlabWebHookURL, action.RepoID, id)
		postRequest, err := http.NewRequest("DELETE", webHookURL, nil)
		if err != nil {
			return nil, err
		}
		postRequest.Header.Add("Content-Type", "application/json;charset=UTF-8")
		q := postRequest.URL.Query()
		q.Set("access_token", tokenInfo.AccessToken)
		postRequest.URL.RawQuery = q.Encode()
		_, err = http_utils.SendHttpRequest(postRequest, 204)
	} else {
		envWebhookUrl, err := utils.GetEnvParameter("WEBHOOK_URL")
		if err != nil {
			return nil, err
		}
		var info gitlabtypes.GitlabWebHookRequest
		if action.RepoAction == models.GlPush {
			info.Url = fmt.Sprintf(envWebhookUrl, "gitlab", models.GlPush, req.ActionId)
			info.PushEvent = true
		}
		if action.RepoAction == models.GIssue {
			info.Url = fmt.Sprintf(envWebhookUrl, "gitlab", models.GIssue, req.ActionId)
			info.IssuesEvent = true
		}
		if action.RepoAction == models.GlTag {
			info.Url = fmt.Sprintf(envWebhookUrl, "gitlab", models.GlTag, req.ActionId)
			info.TagEvent = true
		}
		if action.RepoAction == models.GlRelease {
			info.Url = fmt.Sprintf(envWebhookUrl, "gitlab", models.GlRelease, req.ActionId)
			info.ReleaseEvent = true
		}
		if action.RepoAction == models.GlMergeR {
			info.Url = fmt.Sprintf(envWebhookUrl, "gitlab", models.GlMergeR, req.ActionId)
			info.MergeEvent = true
		}
		_, err = git.createWebHook(tokenInfo, &info, action.RepoID)
	}
	if err != nil {
		return nil, err
	}
	_, err = git.gitlabDb.SetActivateByActionID(req.Activated, uint(tokenInfo.UserID), uint(req.ActionId))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (git *GitlabService) DeleteAction(ctx context.Context, req *gRPCService.DeleteGitlabActionReq) (*gRPCService.DeleteGitlabActionReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, git.tokenDb, "GitlabService", "gitlab")
	if err != nil {
		return nil, err
	}
	action, err := git.gitlabDb.GetGitlabByActionID(uint(req.ActionId))
	if err != nil {
		return nil, err
	}
	id := strconv.Itoa(int(action.HookID))
	webHookURL := fmt.Sprintf(deleteGitlabWebHookURL, action.RepoID, id)
	postRequest, err := http.NewRequest("DELETE", webHookURL, nil)
	if err != nil {
		return nil, err
	}
	postRequest.Header.Add("Content-Type", "application/json;charset=UTF-8")
	q := postRequest.URL.Query()
	q.Set("access_token", tokenInfo.AccessToken)
	postRequest.URL.RawQuery = q.Encode()
	_, err = http_utils.SendHttpRequest(postRequest, 204)
	if err != nil {
		return nil, err
	}
	return req, git.gitlabDb.DeleteByActionID(uint(tokenInfo.UserID), uint(req.ActionId))
}
