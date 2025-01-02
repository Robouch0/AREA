//
// EPITECH PROJECT, 2024
// AREA
// File description:
// trashEmail
//

package gmail

import (
	"area/utils"
	"fmt"
	"io"
	"net/http"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	trashURL = "https://gmail.googleapis.com/gmail/v1/users/%s/messages/%s/%s"
)

func getTrashExpr(trash bool) string {
	if trash {
		return "trash"
	}
	return "untrash"
}

func MoveEmail(tokenAccess string, googleUserID string, messageID string, trash bool) (*GmailMessage, error) {
	url := fmt.Sprintf(trashURL, googleUserID, messageID, getTrashExpr(trash))

	postRequest, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}
	postRequest.Header.Set("Authorization", "Bearer "+tokenAccess)
	postRequest.Header.Add("Content-Type", "message/rfc822")
	postRequest.Header.Add("Accept", "application/json")

	cli := &http.Client{}
	resp, err := cli.Do(postRequest)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		io.Copy(os.Stderr, resp.Body)
		return nil, status.Errorf(codes.Aborted, resp.Status)
	}
	message, err := utils.IoReaderToStruct[GmailMessage](&resp.Body)
	if err != nil {
		return nil, err
	}
	return message, nil
}
