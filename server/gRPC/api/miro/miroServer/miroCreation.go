//
// EPITECH PROJECT, 2025
// AREA
// File description:
// miroCreation
//

package miro_server

import (
	mirotypes "area/gRPC/api/miro/miroTypes"
	"area/models"
	http_utils "area/utils/httpUtils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func sendCreationReq[T any](
	tokenInfo *models.Token,
	boardID string,
	formatURL string,
	body mirotypes.MiroGenericBody[T],
) error {
	url := fmt.Sprintf(formatURL, boardID)

	b, err := json.Marshal(&body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header = http_utils.GetDefaultHTTPHeader("Bearer " + tokenInfo.AccessToken)
	req.Header.Set("accept", "application/json")
	_, err = http_utils.SendHttpRequest(req, 201)
	if err != nil {
		return err
	}
	return nil
}
