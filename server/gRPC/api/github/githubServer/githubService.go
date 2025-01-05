//
// EPITECH PROJECT, 2024
// AREA
// File description:
// githubService
//

package github

import (
	"area/db"
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

	"google.golang.org/grpc"
)

type GithubService struct {
	tokenDb      *db.TokenDb
	GithubDb     *db.GithubDB
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedGithubServiceServer
}

func NewGithubService() (*GithubService, error) {
	tokenDb, err := db.InitTokenDb()

	githubDb, err := db.InitGithubDb()
	return &GithubService{tokenDb: tokenDb, GithubDb: githubDb, reactService: nil}, err
}

func (git *GithubService) InitReactClient(conn *grpc.ClientConn) {
	git.reactService = gRPCService.NewReactionServiceClient(conn)
}

// Update a repository content
func (git *GithubService) UpdateFile(ctx context.Context, req *gRPCService.UpdateRepoFile) (*gRPCService.UpdateRepoFile, error) {
    if req.Owner == "" || req.Repo == "" || req.Path == "" || req.Message == "" || req.Content == "" {
        return nil, errors.New("Some required parameters are empty")
    }

    tokenInfo, err := grpcutils.GetTokenByContext(ctx, git.tokenDb, "GithubService", "github")
    if err != nil {
        log.Println("User is not registered to github")
        return nil, err
    }

    fileInfos, err := git.getRepositoryFileInfos(tokenInfo.AccessToken, git.createFileInfos(req.Owner, req.Repo, req.Path))
    if err != nil {
        return nil, err
    }

    req.Sha = fileInfos.Sha // Sha of the current state of the file
    req.Content = utils.EncodeToBase64(req.Content)
    b, err := json.Marshal(req)
    if err != nil {
        return nil, err
    }
    url := fmt.Sprintf("https://api.github.com/repos/%v/%v/contents/%v", req.Owner, req.Repo, req.Path)
    putRequest, err := http.NewRequest("PUT", url, bytes.NewBuffer(b))
    putRequest.Header = http_utils.GetDefaultBearerHTTPHeader(tokenInfo.AccessToken)
    putRequest.Header.Add("Accept", "application/vnd.github+json")

    resp, err := http_utils.SendHttpRequest(putRequest, 200)
    if err != nil {
        return nil, err
    }
    log.Println("Here: ", resp.Body) // Do something with it
    return req, nil
}

func (git *GithubService) UpdateRepository(ctx context.Context, req *gRPCService.UpdateRepoInfos) (*gRPCService.UpdateRepoInfos, error) {
	userID, errClaim := grpcutils.GetUserIdFromContext(ctx, "GithubService")
	if errClaim != nil {
		return nil, errClaim
	}

	if req.Owner == "" || req.Repo == "" {
		return nil, errors.New("Some required parameters are empty")
	}

	tokenInfo, err := git.tokenDb.GetUserTokenByProvider(int64(userID), "github")
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://api.github.com/repos/%v/%v", req.Owner, req.Repo)
	pathRequest, err := http.NewRequest("PATCH", url, bytes.NewBuffer(b))
	pathRequest.Header = http_utils.GetDefaultBearerHTTPHeader(tokenInfo.AccessToken)
	pathRequest.Header.Add("Accept", "application/vnd.github+json")

	cli := &http.Client{}
	resp, err := cli.Do(pathRequest)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		var body map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&body)
		if err != nil {
			return nil, err
		}
		log.Println(body)
		return nil, errors.New(resp.Status)
	}
	log.Println(resp.Body) // Do something with it
	return req, nil
}

func (git *GithubService) DeleteFile(ctx context.Context, req *gRPCService.DeleteRepoFile) (*gRPCService.DeleteRepoFile, error) {
	userID, errClaim := grpcutils.GetUserIdFromContext(ctx, "GithubService")
	if errClaim != nil {
		return nil, errClaim
	}

	if req.Owner == "" || req.Repo == "" || req.Path == "" || req.Message == "" {
		return nil, errors.New("Some required parameters are empty")
	}

	url := fmt.Sprintf("https://api.github.com/repos/%v/%v/contents/%v", req.Owner, req.Repo, req.Path)
	tokenInfo, err := git.tokenDb.GetUserTokenByProvider(int64(userID), "github")
	if err != nil {
		return nil, err
	}

	fileInfos, err := git.getRepositoryFileInfos(tokenInfo.AccessToken, git.createFileInfos(req.Owner, req.Repo, req.Path))
	if err != nil {
		return nil, err
	}

	req.Sha = fileInfos.Sha
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	pathRequest, err := http.NewRequest("DELETE", url, bytes.NewBuffer(b))
	pathRequest.Header = http_utils.GetDefaultBearerHTTPHeader(tokenInfo.AccessToken)
	pathRequest.Header.Add("Accept", "application/vnd.github+json")

	cli := &http.Client{}
	resp, err := cli.Do(pathRequest)
	if err != nil {
		return nil, err
	}
	if resp.Status != "200 OK" {
		return nil, errors.New(resp.Status)
	}

	log.Println(resp.Body)
	return req, nil
}

