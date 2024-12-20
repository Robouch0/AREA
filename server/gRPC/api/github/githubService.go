//
// EPITECH PROJECT, 2024
// AREA
// File description:
// githubService
//

package github

import (
	gRPCService "area/protogen/gRPC/proto"
	"area/utils"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type GithubService struct {
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedGithubServiceServer
}

func NewGithubService() GithubService {
	return GithubService{reactService: nil}
}

// Update a repository content
func (git *GithubService) UpdateFile(_ context.Context, req *gRPCService.UpdateRepoFile) (*gRPCService.UpdateRepoFile, error) {
	if req.Owner == "" || req.Repo == "" || req.Path == "" || req.Message == "" || req.Content == "" {
		return nil, errors.New("Some required parameters are empty")
	}
	bearerTok, err := utils.GetEnvParameterToBearer("API_GITHUB")
	if err != nil {
		return nil, err
	}

	fileInfos, err := git.getRepositoryFileInfos(bearerTok, git.createFileInfos(req.Owner, req.Repo, req.Path))
	if err != nil {
		return nil, err
	}

	req.Sha = fileInfos.Sha // Sha of the current state of the file
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://api.github.com/repos/%v/%v/contents/%v", req.Owner, req.Repo, req.Path)
	putRequest, err := http.NewRequest("PUT", url, bytes.NewBuffer(b))
	putRequest.Header = utils.GetDefaultHTTPHeader(bearerTok)
	putRequest.Header.Add("Accept", "application/vnd.github+json")

	cli := &http.Client{}
	resp, err := cli.Do(putRequest)
	if err != nil {
		return nil, err
	}
	if resp.Status != "200 OK" {
		return nil, errors.New(resp.Status)
	}
	log.Println("Here: ", resp.Body) // Do something with it
	return req, nil
}

func (git *GithubService) UpdateRepository(_ context.Context, req *gRPCService.UpdateRepoInfos) (*gRPCService.UpdateRepoInfos, error) {
	if req.Owner == "" || req.Repo == "" {
		return nil, errors.New("Some required parameters are empty")
	}

	bearerTok, err := utils.GetEnvParameterToBearer("API_GITHUB")
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://api.github.com/repos/%v/%v", req.Owner, req.Repo)
	pathRequest, err := http.NewRequest("PATCH", url, bytes.NewBuffer(b))
	pathRequest.Header = utils.GetDefaultHTTPHeader(bearerTok)
	pathRequest.Header.Add("Accept", "application/vnd.github+json")

	cli := &http.Client{}
	resp, err := cli.Do(pathRequest)
	if err != nil {
		return nil, err
	}
	if resp.Status != "200 OK" {
		return nil, errors.New(resp.Status)
	}
	log.Println(resp.Body) // Do something with it
	return req, nil
}

func (git *GithubService) DeleteFile(_ context.Context, req *gRPCService.DeleteRepoFile) (*gRPCService.DeleteRepoFile, error) {
	if req.Owner == "" || req.Repo == "" || req.Path == "" || req.Message == "" {
		return nil, errors.New("Some required parameters are empty")
	}

	url := fmt.Sprintf("https://api.github.com/repos/%v/%v/contents/%v", req.Owner, req.Repo, req.Path)
	bearerTok, err := utils.GetEnvParameterToBearer("API_GITHUB")
	if err != nil {
		return nil, err
	}

	fileInfos, err := git.getRepositoryFileInfos(bearerTok, git.createFileInfos(req.Owner, req.Repo, req.Path))
	if err != nil {
		return nil, err
	}

	req.Sha = fileInfos.Sha
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	pathRequest, err := http.NewRequest("DELETE", url, bytes.NewBuffer(b))
	pathRequest.Header = utils.GetDefaultHTTPHeader(bearerTok)
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
