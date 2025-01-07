//
// EPITECH PROJECT, 2025
// AREA
// File description:
// stopPubSub
//

package gmail

import (
	"area/models"
	http_utils "area/utils/httpUtils"
	"fmt"
	"net/http"
)

const (
	stopPubSubURL = "https://gmail.googleapis.com/gmail/v1/users/%s/stop"
)

func StopPubSub(tokenInfo *models.Token) error {
	url := fmt.Sprintf(stopPubSubURL, "me")

	postRequest, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	postRequest.Header.Set("Authorization", "Bearer "+tokenInfo.AccessToken)
	postRequest.Header.Add("Content-Type", "application/json;charset=UTF-8")
	postRequest.Header.Add("Accept", "application/json")
	_, err = http_utils.SendHttpRequest(postRequest, 204)
	if err != nil {
		return err
	}
	return nil
}
