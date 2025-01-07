//
// EPITECH PROJECT, 2025
// AREA
// File description:
// stop
//

package drive

import (
	"area/models"
	http_utils "area/utils/httpUtils"
	"bytes"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	stopWatchURL = "https://www.googleapis.com/drive/v3/channels/stop"
)

func StopWatchDrive(tokenInfo *models.Token, channelID, resourceID string) error {
	b, err := json.Marshal(&DriveChannel{
		ID:         channelID,
		ResourceId: resourceID,
	})
	if err != nil {
		return status.Errorf(codes.DataLoss, "Invalid data sent")
	}
	postRequest, err := http.NewRequest("POST", stopWatchURL, bytes.NewBuffer(b))
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
