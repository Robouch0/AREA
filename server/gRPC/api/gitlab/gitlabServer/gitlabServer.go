//
// EPITECH PROJECT, 2024
// AREA
// File description:
// gitlabServer
//

package gitlab_server

import (
	"area/db"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	http_utils "area/utils/httpUtils"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	return nil, status.Errorf(codes.Unavailable, "No Action Gitlab yet") // TODO Matthieu
}

func (git *GitlabService) DeleteAction(ctx context.Context, req *gRPCService.DeleteGitlabActionReq) (*gRPCService.DeleteGitlabActionReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, git.tokenDb, "GithubService", "github")
	if err != nil {
		return nil, err
	}
	action, err := git.gitlabDb.GetGitlabByActionID(uint(req.ActionId))
	if err != nil {
		return nil, err
	}
	webHookURL := fmt.Sprintf(pushWebHookURL, action.RepoID, action.HookID)
	postRequest, err := http.NewRequest("DELETE", webHookURL, nil)
	if err != nil {
		return nil, err
	}
	postRequest.Header.Add("Content-Type", "application/json;charset=UTF-8")
	q := postRequest.URL.Query()
	q.Set("access_token", tokenInfo.AccessToken)
	postRequest.URL.RawQuery = q.Encode()
	_, err = http_utils.SendHttpRequest(postRequest, 201)
	if err != nil {
		return nil, err
	}
	return req, git.gitlabDb.DeleteByActionID(uint(tokenInfo.UserID), uint(req.ActionId))
}
