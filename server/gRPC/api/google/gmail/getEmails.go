//
// EPITECH PROJECT, 2024
// AREA
// File description:
// listEmails
//

package gmail

import (
	"area/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	listEmailsURL = "https://gmail.googleapis.com/gmail/v1/users/%s/messages"
	maxResults    = "50"

	getEmailURL = "https://gmail.googleapis.com/gmail/v1/users/%s/messages/%s"
)

type MessageListRes struct {
	Messages           []GmailMessage `json:"messages,omitempty"`
	NextPageToken      string         `json:"nextPageToken,omitempty"`
	ResultSizeEstimate int            `json:"resultSizeEstimate,omitempty"`
}

func GetListEmails(googleUserID string, accessToken string) (*MessageListRes, error) {
	url := fmt.Sprintf(listEmailsURL, googleUserID)
	request, _ := http.NewRequest("GET", url, nil) // GetUserMail request
	request.Header = utils.GetDefaultBearerHTTPHeader(accessToken)
	request.Header.Add("Accept", "application/json")

	request.URL.Query().Set("maxResults", maxResults)

	client := &http.Client{}
	result, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if result.StatusCode != 200 {
		io.Copy(os.Stderr, result.Body)
		return nil, status.Errorf(codes.Aborted, result.Status)
	}
	messageList := &MessageListRes{}
	err = json.NewDecoder(result.Body).Decode(messageList)
	if err != nil {
		return nil, err
	}
	return messageList, nil
}

func GetEmail( // Look for later if it will not be called getEmailHeader
	googleUserID string,
	accessToken string,
	messageID string,
	format string,
	metadata string,
) (*GmailMessage, error) {
	url := fmt.Sprintf(getEmailURL, googleUserID, messageID)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header = utils.GetDefaultBearerHTTPHeader(accessToken)
	request.Header.Add("Accept", "application/json")
	request.URL.Query().Set("format", format)
	request.URL.Query().Add("metadataHeaders", metadata)

	client := &http.Client{}
	result, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if result.StatusCode != 200 {
		io.Copy(os.Stderr, result.Body)
		return nil, status.Errorf(codes.Aborted, result.Status)
	}
	message := &GmailMessage{}
	err = json.NewDecoder(result.Body).Decode(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
