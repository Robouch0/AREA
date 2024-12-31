//
// EPITECH PROJECT, 2024
// AREA
// File description:
// watch
//

package drive

import (
	"area/utils"
	http_utils "area/utils/httpUtils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	watchFileURL    = "https://www.googleapis.com/drive/v3/files/%s/watch"
	watchChangesURL = "https://www.googleapis.com/drive/v3/changes/watch"
)

type DriveChannel struct {
	Payload bool   `json:"payload,omitempty"`
	ID      string `json:"id,omitempty"`
	Address string `json:"address,omitempty"`
	Type    string `json:"type,omitempty"`
}

// Faire les types et les fonctions
// Faire le tri côté gRPC pour avoir le bon fileID

func WatchFile(accessToken, fileID, channelID string, actionID uint) (*DriveChannel, error) {
	rawURL, err := utils.GetEnvParameter("WEBHOOK_URL")
	if err != nil {
		return nil, err
	}
	wURL := fmt.Sprintf(rawURL, "google", "watchFile", actionID)

	watchBody := &DriveChannel{
		Payload: true,
		ID:      channelID,
		Address: wURL,
		Type:    "web_hook",
	}

	b, err := json.Marshal(&watchBody)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(watchFileURL, fileID)
	postRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	postRequest.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
	postRequest.Header.Add("Accept", "application/json")
	resp, err := http_utils.SendHttpRequest(postRequest, 200)
	if err != nil {
		return nil, err
	}
	watchResBody, err := utils.IoReaderToStruct[DriveChannel](&resp.Body)
	if err != nil {
		return nil, err
	}
	return watchResBody, nil
}

func WatchChanges(accessToken, channelID string, actionID uint) (*DriveChannel, error) {
	rawURL, err := utils.GetEnvParameter("WEBHOOK_URL")
	if err != nil {
		return nil, err
	}
	wURL := fmt.Sprintf(rawURL, "google", "watchFile", actionID)

	watchBody := &DriveChannel{
		Payload: true,
		ID:      channelID,
		Address: wURL,
		Type:    "web_hook",
	}

	b, err := json.Marshal(&watchBody)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(watchChangesURL)
	postRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	postRequest.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
	postRequest.Header.Add("Accept", "application/json")
	resp, err := http_utils.SendHttpRequest(postRequest, 200)
	if err != nil {
		return nil, err
	}
	watchResBody, err := utils.IoReaderToStruct[DriveChannel](&resp.Body)
	if err != nil {
		return nil, err
	}
	return watchResBody, nil
}
