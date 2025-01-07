//
// EPITECH PROJECT, 2025
// AREA
// File description:
// update
//

package github_webhooks

import (
	"area/models"
	"area/utils"
	http_utils "area/utils/httpUtils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	updateWebHookURL = "https://api.github.com/repos/%v/%v/hooks/%v"
)

func UpdatePushWebhook(tokenInfo *models.Token, ctx context.Context, action *models.Github, activated bool) error {
	envWebhookUrl, err := utils.GetEnvParameter("WEBHOOK_URL")
	if err != nil {
		return err
	}

	b, err := json.Marshal(&GitWebHookRequest{
		Event:  []string{string(action.RepoAction)},
		Active: activated,
		Config: GithubConfig{Url: fmt.Sprintf(envWebhookUrl, "github", action.RepoAction, action.ActionID), Content: "json"}})
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to convert the content to bytes"))
	}

	webHookURL := fmt.Sprintf(updateWebHookURL, action.RepoOwner, action.RepoName)
	postRequest, err := http.NewRequest("PATCH", webHookURL, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	postRequest.Header.Set("Authorization", "Bearer "+tokenInfo.AccessToken)
	postRequest.Header.Add("Accept", "application/vnd.github+json")
	_, err = http_utils.SendHttpRequest(postRequest, 200)
	if err != nil {
		return err
	}
	return nil
}
