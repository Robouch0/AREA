//
// EPITECH PROJECT, 2024
// AREA
// File description:
// githubServiceInfos
//

package github

import (
	gRPCService "area/protogen/gRPC/proto"
	"area/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type FileInfos struct {
	Type     string `json:"type"`
	Encoding string `json:"encoding"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Content  string `json:"content"`
	Sha      string `json:"sha"`
}

func (git *GithubService) getRepositoryFileInfos(bearerTok string, req *gRPCService.UpdateRepoFile) (*FileInfos, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%v/%v/contents/%v", req.Owner, req.Repo, req.Path)
	postRequest, err := http.NewRequest("GET", url, nil)
	postRequest.Header = utils.GetDefaultHTTPHeader(bearerTok)
	postRequest.Header.Add("Accept", "application/vnd.github+json")

	cli := &http.Client{}
	resp, err := cli.Do(postRequest)
	if err != nil {
		return nil, err
	}
	if resp.Status != "200 OK" {
		return nil, errors.New(resp.Status)
	}
	fileInfos := &FileInfos{}
	err = json.NewDecoder(resp.Body).Decode(fileInfos)
	if err != nil {
		return nil, err
	}
	return fileInfos, nil
}
