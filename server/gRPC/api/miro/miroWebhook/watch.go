//
// EPITECH PROJECT, 2025
// AREA
// File description:
// watch
//

package miro_webhook

import (
	"area/utils"
	http_utils "area/utils/httpUtils"
	"bytes"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	createWebhookURL = "https://api.miro.com/v2-experimental/webhooks/board_subscriptions"
)

type CreateWebhookBody struct {
	BoardId     string `json:"boardId"`
	CallbackURL string `json:"callbackUrl"`
	Status      string `json:"status"`
}

type CreateWebhookResponse struct {
	ID string `json:"id"`
}

func CreateWebhook(accessToken string, body *CreateWebhookBody) (*CreateWebhookResponse, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot marshalize miro webhook body, %v", err)
	}
	req, err := http.NewRequest("POST", createWebhookURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot create miro webhook request, %v", err)
	}
	req.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
	req.Header.Set("Accept", "application/json")

	resp, err := http_utils.SendHttpRequest(req, 201)
	if err != nil {
		return nil, err
	}
	return utils.IoReaderToStruct[CreateWebhookResponse](&resp.Body)
}
