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
)

type GitlabService struct {
	tokenDb      *db.TokenDb
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedGitlabServiceServer
}

func NewGitlabService() (*GitlabService, error) {
	tokenDb, err := db.InitTokenDb()
	if err != nil {
		return nil, err
	}

	return &GitlabService{tokenDb: tokenDb, reactService: nil}, err
}

func (git *GitlabService) InitReactClient(conn *grpc.ClientConn) {
	git.reactService = gRPCService.NewReactionServiceClient(conn)
}

func (git *GitlabService) CreateFile(ctx context.Context, req *gRPCService.CreateLabRepoFile) (*gRPCService.CreateLabRepoFile, error) {
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

	resp, err := http_utils.SendHttpRequest(pathRequest, 201);
	if err != nil {
		return nil, err
	}
	log.Println(resp.Body)
	return req, nil
}

func (git *GitlabService) UpdateFile(ctx context.Context, req *gRPCService.UpdateLabRepoFile) (*gRPCService.UpdateLabRepoFile, error) {
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

	resp, err := http_utils.SendHttpRequest(pathRequest, 200);
	if err != nil {
		return nil, err
	}
	log.Println(resp.Body)
	return req, nil
}

func (git *GitlabService) DeleteFile(ctx context.Context, req *gRPCService.DeleteLabRepoFile) (*gRPCService.DeleteLabRepoFile, error) {
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

	resp, err := http_utils.SendHttpRequest(pathRequest, 204);
	if err != nil {
		return nil, err
	}
	log.Println(resp.Body)
	return req, nil
}

func (git *GitlabService) MarkItemAsDone(ctx context.Context, req *gRPCService.TodoLabItemDone) (*gRPCService.TodoLabItemDone, error) {
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

	resp, err := http_utils.SendHttpRequest(pathRequest, 201);
	if err != nil {
		return nil, err
	}
	log.Println(resp.Body)
	return req, nil
}


func (git *GitlabService) MarkAllItemAsDone(ctx context.Context, req *gRPCService.AllTodoLabItemDone) (*gRPCService.AllTodoLabItemDone, error) {
	url := fmt.Sprintf("https://www.gitlab.com/api/v4/todos/mark_as_done")
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

	resp, err := http_utils.SendHttpRequest(pathRequest, 204);
	if err != nil {
		return nil, err
	}
	log.Println(resp.Body)
	return req, nil
}