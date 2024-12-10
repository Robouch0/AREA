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

func (git *GithubService) UpdateRepository(_ context.Context, req *gRPCService.UpdateRepoReq) (*gRPCService.UpdateRepoReq, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%v/%v", req.Owner, req.Repo)
	bearerTok, err := utils.GetEnvParameterToBearer("API_GITHUB")
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	postRequest, err := http.NewRequest("PATCH", url, bytes.NewBuffer(b))
	postRequest.Header.Set("Authorization", bearerTok)
	postRequest.Header.Add("Content-Type", "application/json;charset=UTF-8")
	postRequest.Header.Add("Accept", "application/vnd.github+json")

	cli := &http.Client{}
	resp, err := cli.Do(postRequest)
	if err != nil {
		return nil, err
	}
	if resp.Status != "200 OK" {
		return nil, errors.New(resp.Status)
	}
	log.Println(resp.Body) // Do something with it
	return req, nil
}
