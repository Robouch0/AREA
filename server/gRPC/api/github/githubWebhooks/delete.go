//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// delete
//

package github_webhooks

import (
	"area/models"
	http_utils "area/utils/httpUtils"
	"context"
	"fmt"
	"net/http"
)

const (
	deleteWebHookURL = "https://api.github.com/repos/%v/%v/hooks/%v"
)

func DeletePushWebhook(tokenInfo *models.Token, ctx context.Context, action *models.Github) error {
	webHookURL := fmt.Sprintf(deleteWebHookURL, action.RepoOwner, action.RepoName, action.HookId)
	postRequest, err := http.NewRequest("DELETE", webHookURL, nil)
	if err != nil {
		return err
	}
	postRequest.Header.Set("Authorization", "Bearer "+tokenInfo.AccessToken)
	postRequest.Header.Add("Accept", "application/vnd.github+json")
	_, err = http_utils.SendHttpRequest(postRequest, 204)
	if err != nil {
		return err
	}
	return nil
}
