//
// EPITECH PROJECT, 2024
// AREA
// File description:
// listEmails
//

package gmail

import (
	conv_utils "area/utils/convUtils"
	http_utils "area/utils/httpUtils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	listEmailsURL = "https://gmail.googleapis.com/gmail/v1/users/%s/messages"
	maxResults    = "10"

	getEmailURL = "https://gmail.googleapis.com/gmail/v1/users/%s/messages/%s"
)

type MessageListRes struct {
	Messages           []GmailMessage `json:"messages,omitempty"`
	NextPageToken      string         `json:"nextPageToken,omitempty"`
	ResultSizeEstimate int            `json:"resultSizeEstimate,omitempty"`
}

func GetListEmails(googleUserID string, accessToken string, label string) (*MessageListRes, error) {
	url := fmt.Sprintf(listEmailsURL, googleUserID)
	request, _ := http.NewRequest("GET", url, nil) // GetUserMail request
	request.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
	request.Header.Add("Accept", "application/json")

	q := request.URL.Query()
	q.Add("maxResults", maxResults)
	if label != "" {
		q.Add("labelIds", label)
	}
	request.URL.RawQuery = q.Encode()

	log.Println("Query: ", request.URL.String())
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

func GetEmail(
	googleUserID string,
	accessToken string,
	messageID string,
	format string,
	metadata string,
) (*GmailMessage, error) {
	url := fmt.Sprintf(getEmailURL, googleUserID, messageID)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
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
	message, err := conv_utils.IoReaderToStruct[GmailMessage](&result.Body)
	if err != nil {
		return nil, err
	}
	return message, nil
}
