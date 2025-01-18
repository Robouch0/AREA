//
// EPITECH PROJECT, 2024
// AREA
// File description:
// commentsHandler
//

package drive

import (
	conv_utils "area/utils/convUtils"
	http_utils "area/utils/httpUtils"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	createCommentURL = "https://www.googleapis.com/drive/v3/files/%s/comments"
	deleteCommentURL = "https://www.googleapis.com/drive/v3/files/%s/comments/%s"
	listCommentURL   = "https://www.googleapis.com/drive/v3/files/%s/comments"
	getCommentURL    = "https://www.googleapis.com/drive/v3/files/%s/comments/%s"
	updateCommentURL = "https://www.googleapis.com/drive/v3/files/%s/comments/%s"
)

type DriveListComments struct {
	Comments []DriveComment `json:"comments,omitempty"`
}

func GetComment(accessToken, fileID, commentID, fields string) (*DriveComment, error) {
	url := fmt.Sprintf(getCommentURL, fileID, commentID)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
	request.Header.Add("Accept", "application/json")
	q := request.URL.Query()
	q.Set("fields", fields)
	request.URL.RawQuery = q.Encode()

	resp, err := http_utils.SendHttpRequest(request, 200)
	if err != nil {
		return nil, err
	}
	list, err := conv_utils.IoReaderToStruct[DriveComment](&resp.Body)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func ListComments(accessToken, fileID, fields string) (*DriveListComments, error) {
	url := fmt.Sprintf(listCommentURL, fileID)
	log.Println(url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
	request.Header.Add("Accept", "application/json")
	q := request.URL.Query()
	q.Set("fields", fields)
	request.URL.RawQuery = q.Encode()

	resp, err := http_utils.SendHttpRequest(request, 200)
	if err != nil {
		return nil, err
	}
	list, err := conv_utils.IoReaderToStruct[DriveListComments](&resp.Body)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func CreateCommentOnFile(accessToken, fileID, fields string, comment DriveComment) (*DriveComment, error) {
	url := fmt.Sprintf(createCommentURL, fileID)
	b, err := json.Marshal(&comment)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to convert the content to bytes"))
	}

	postRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	postRequest.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
	postRequest.Header.Add("Accept", "application/json")
	q := postRequest.URL.Query()
	q.Set("fields", fields)
	postRequest.URL.RawQuery = q.Encode()

	resp, err := http_utils.SendHttpRequest(postRequest, 200)
	if err != nil {
		return nil, err
	}
	drive, err := conv_utils.IoReaderToStruct[DriveComment](&resp.Body)
	if err != nil {
		return nil, err
	}
	return drive, nil
}

func DeleteCommentOnFile(accessToken, fileID, commentID, fields string) error {
	url := fmt.Sprintf(deleteCommentURL, fileID, commentID)
	postRequest, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	postRequest.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
	postRequest.Header.Add("Accept", "application/json")
	q := postRequest.URL.Query()
	q.Set("fields", fields)
	postRequest.URL.RawQuery = q.Encode()

	_, err = http_utils.SendHttpRequest(postRequest, 204)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCommentOnFile(accessToken, fileID, commentID, fields string, comment DriveComment) (*DriveComment, error) {
	url := fmt.Sprintf(updateCommentURL, fileID, commentID)
	b, err := json.Marshal(&comment)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to convert the content to bytes"))
	}

	postRequest, err := http.NewRequest("PATCH", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	postRequest.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
	postRequest.Header.Add("Accept", "application/json")
	q := postRequest.URL.Query()
	q.Set("fields", fields)
	postRequest.URL.RawQuery = q.Encode()

	resp, err := http_utils.SendHttpRequest(postRequest, 200)
	if err != nil {
		return nil, err
	}
	drive, err := conv_utils.IoReaderToStruct[DriveComment](&resp.Body)
	if err != nil {
		return nil, err
	}
	return drive, nil
}
