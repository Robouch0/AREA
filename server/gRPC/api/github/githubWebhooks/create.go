//
// EPITECH PROJECT, 2025
// AREA
// File description:
// githubWebhooks
//

package github_webhooks

import (
	"area/models"
	http_utils "area/utils/httpUtils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GithubConfig struct {
	Url     string `json:"url"`
	Content string `json:"content_type"`
}

type GitWebHookRequest struct {
	Event  []string     `json:"events"`
	Active bool         `json:"active"`
	Config GithubConfig `json:"config"`
}

func SendCreateWebHook(
	tokenInfo *models.Token,
	owner string,
	repo string,
	url string,
	webhookReq *GitWebHookRequest,
) error {
	b, err := json.Marshal(webhookReq)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to convert the content to bytes"))
	}

	webHookURL := fmt.Sprintf(url, owner, repo)
	postRequest, err := http.NewRequest("POST", webHookURL, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	postRequest.Header.Set("Authorization", "Bearer "+tokenInfo.AccessToken)
	postRequest.Header.Add("Accept", "application/vnd.github+json")
	_, err = http_utils.SendHttpRequest(postRequest, 201)
	if err != nil {
		return err
	}
	return nil
}
