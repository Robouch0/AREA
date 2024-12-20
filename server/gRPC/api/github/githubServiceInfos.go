//
// EPITECH PROJECT, 2024
// AREA
// File description:
// githubServiceInfos
//

package github

import (
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

type RepoFileInfo struct {
	Owner string `json:"owner"`
	Repo  string `json:"repo"`
	Path  string `json:"path"`
}

func (git *GithubService) createFileInfos(Owner string, Repo string, Path string) *RepoFileInfo {
	file := new(RepoFileInfo)

	file.Owner = Owner
	file.Repo = Repo
	file.Path = Path
	return file
}

func (git *GithubService) getRepositoryFileInfos(accessToken string, req *RepoFileInfo) (*FileInfos, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%v/%v/contents/%v", req.Owner, req.Repo, req.Path)
	postRequest, err := http.NewRequest("GET", url, nil)
	postRequest.Header = utils.GetDefaultBearerHTTPHeader(accessToken)
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
