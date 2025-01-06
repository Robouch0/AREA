//
// EPITECH PROJECT, 2025
// AREA
// File description:
// githubServiceRepos
//

package github

import (
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
)

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
