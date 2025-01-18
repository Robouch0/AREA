//
// EPITECH PROJECT, 2024
// AREA
// File description:
// watchRequest
//

package gmail

import (
	"area/models"
	"area/utils"
	conv_utils "area/utils/convUtils"
	http_utils "area/utils/httpUtils"
	"bytes"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	watchRequestURL = "https://www.googleapis.com/gmail/v1/users/me/watch"
)

/* WebHook request */

type WatchRequestBody struct {
	TopicName           string   `json:"topicName"`
	LabelIds            []string `json:"labelIds"`
	LabelFilterBehavior string   `json:"labelFilterBehavior"`
}

type WatchReponseBody struct {
	HistoryID  string `json:"history_id"`
	Expiration string `json:"expiration"`
}

/* WebHook response */

type GmailPayload struct {
	EmailAddress string  `json:"emailAddress"`
	HistoryId    float64 `json:"historyId,omitempty"`
}

type PubSubMessage struct {
	Data        string `json:"data"`
	MessageId   string `json:"messageId"`
	PublishTime string `json:"publishTime"`
}

type PubSubPayload struct {
	Message PubSubMessage `json:"message"`

	Subscription string `json:"subscription"`
}

/* Webhook request function */

func SendWatchMeRequest(tokenInfo *models.Token) (*WatchReponseBody, error) {
	topic, err := utils.GetEnvParameter("GOOGLE_TOPIC_NAME")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "No 'GOOGLE_TOPIC_NAME' in environment")
	}
	watchBody := &WatchRequestBody{
		TopicName:           topic,
		LabelIds:            []string{"UNREAD", "INBOX"},
		LabelFilterBehavior: "INCLUDE",
	}

	b, err := json.Marshal(&watchBody)
	if err != nil {
		return nil, err
	}
	postRequest, err := http.NewRequest("POST", watchRequestURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	postRequest.Header.Set("Authorization", "Bearer "+tokenInfo.AccessToken)
	postRequest.Header.Add("Content-Type", "application/json;charset=UTF-8")
	postRequest.Header.Add("Accept", "application/json")
	resp, err := http_utils.SendHttpRequest(postRequest, 200)
	if err != nil {
		return nil, err
	}
	watchResBody, err := conv_utils.IoReaderToStruct[WatchReponseBody](&resp.Body)
	if err != nil {
		return nil, err
	}
	return watchResBody, nil
}
