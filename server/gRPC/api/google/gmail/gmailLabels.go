//
// EPITECH PROJECT, 2024
// AREA
// File description:
// gmailLabels
//

package gmail

import (
	"area/utils"
	http_utils "area/utils/httpUtils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	listLabelURL   = "https://gmail.googleapis.com/gmail/v1/users/%s/labels"
	createLabelURL = "https://gmail.googleapis.com/gmail/v1/users/%s/labels"
	deleteLabelURL = "https://gmail.googleapis.com/gmail/v1/users/%s/labels/%s"
	updateLabelURL = "https://gmail.googleapis.com/gmail/v1/users/%s/labels/%s"
)

// Visibility of a gmail message with this label
type GmailMessageVisibility string

const (
	Show GmailMessageVisibility = "show"
	Hide                        = "hide"
)

// Visibility of a gmail message with this label
type GmailLabelVisibility string

const (
	labelShow         GmailLabelVisibility = "labelShow"
	labelShowIfUnread                      = "labelShowIfUnread"
	labelHide                              = "labelHide"
)

type GmailLabelType string

const (
	System GmailLabelType = "system"
	User                  = "user"
)

type GmailLabel struct {
	ID                    string                 `json:"id,omitempty"`
	Name                  string                 `json:"name,omitempty"`
	MessageListVisibility GmailMessageVisibility `json:"messageListVisibility,omitempty"`
	LabelListVisibility   GmailLabelVisibility   `json:"labelListVisibility,omitempty"`
	Type                  GmailLabelType         `json:"type,omitempty"`
}

type ListGmailLabel struct {
	Labels []GmailLabel `json:"labels"`
}

func ListLabels(accessToken, userID string) (*ListGmailLabel, error) {
	url := fmt.Sprintf(listLabelURL, userID)
	request, _ := http.NewRequest("GET", url, nil) // GetUserMail request
	request.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
	request.Header.Add("Accept", "application/json")

	client := &http.Client{}
	result, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if result.StatusCode != 200 {
		io.Copy(os.Stderr, result.Body)
		return nil, status.Errorf(codes.Aborted, result.Status)
	}
	messageList := &ListGmailLabel{}
	err = json.NewDecoder(result.Body).Decode(messageList)
	if err != nil {
		return nil, err
	}
	return messageList, nil
}

func PutLabel(accessToken, userID, labelName string, putLabel GmailLabel) (*GmailLabel, error) {
	labels, err := ListLabels(accessToken, userID)
	if err != nil {
		return nil, err
	}
	for _, label := range labels.Labels {
		if label.Name == labelName {
			url := fmt.Sprintf(updateLabelURL, userID, label.ID)
			b, err := json.Marshal(&putLabel)
			if err != nil {
				return nil, err
			}
			putRequest, err := http.NewRequest("PUT", url, bytes.NewBuffer(b))
			if err != nil {
				return nil, err
			}
			putRequest.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
			putRequest.Header.Add("Accept", "application/json")
			resp, err := http_utils.SendHttpRequest(putRequest, 200)
			if err != nil {
				return nil, err
			}
			return utils.IoReaderToStruct[GmailLabel](&resp.Body)
		}
	}
	return nil, status.Errorf(codes.NotFound, "Did not found label named: %v", labelName)
}

func CreateLabel(accessToken, userID string, newLabel GmailLabel) (*GmailLabel, error) {
	url := fmt.Sprintf(createLabelURL, userID)

	b, err := json.Marshal(&newLabel)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to convert the content to bytes"))
	}

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
	lab, err := utils.IoReaderToStruct[GmailLabel](&resp.Body)
	if err != nil {
		return nil, err
	}
	return lab, nil
}

func DeleteLabel(accessToken, userID, labelName string) error {
	labels, err := ListLabels(accessToken, userID)
	if err != nil {
		return err
	}
	for _, label := range labels.Labels {
		if label.Name == labelName {
			url := fmt.Sprintf(deleteLabelURL, userID, label.ID)
			deleteRequest, err := http.NewRequest("DELETE", url, nil)
			if err != nil {
				return err
			}
			deleteRequest.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
			deleteRequest.Header.Add("Accept", "application/json")
			_, err = http_utils.SendHttpRequest(deleteRequest, 204)
			return err
		}
	}
	return status.Errorf(codes.NotFound, "Did not found label named: %v", labelName)
}
