//
// EPITECH PROJECT, 2024
// AREA
// File description:
// fileHandler
//

package drive

import (
	"area/utils"
	http_utils "area/utils/httpUtils"
	"bytes"
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
	listFilesURL  = "https://www.googleapis.com/drive/v3/files"
	createFileURL = "https://www.googleapis.com/drive/v3/files" // Metadata only
	deleteFileURL = "https://www.googleapis.com/drive/v3/files/%s"
	updateFileURL = "https://www.googleapis.com/drive/v3/files/%s" // Metadata only
	copyFileURL   = "https://www.googleapis.com/drive/v3/files/%s/copy"
)

type ListDriveFile struct {
	Files []DriveFile `json:"files,omitempty"`
}

func ListFiles(accessToken string) (*ListDriveFile, error) {
	request, err := http.NewRequest("GET", listFilesURL, nil)
	if err != nil {
		return nil, err
	}
	request.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
	request.Header.Add("Accept", "application/json")

	resp, err := http_utils.SendHttpRequest(request, 200)
	if err != nil {
		return nil, err
	}
	list, err := utils.IoReaderToStruct[ListDriveFile](&resp.Body)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// Use ?uploadType=resumable when doing upload file creation

// Create an empty file in drive
func CreateEmptyFile(accessToken string, file DriveFile) (*DriveFile, error) {
	log.Println(file)
	b, err := json.Marshal(&file)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to convert the content to bytes"))
	}

	postRequest, err := http.NewRequest("POST", createFileURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	postRequest.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
	postRequest.Header.Add("Accept", "application/json")

	resp, err := http_utils.SendHttpRequest(postRequest, 200)
	if err != nil {
		return nil, err
	}
	io.Copy(os.Stdout, postRequest.Body)
	drive, err := utils.IoReaderToStruct[DriveFile](&resp.Body)
	if err != nil {
		return nil, err
	}
	return drive, nil
}

func CopyFile(accessToken string, fileID string, file DriveFile) (*DriveFile, error) {
	url := fmt.Sprintf(copyFileURL, fileID)
	b, err := json.Marshal(&file)
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
	drive, err := utils.IoReaderToStruct[DriveFile](&resp.Body)
	if err != nil {
		return nil, err
	}
	return drive, nil
}

func DeleteFile(accessToken, fileID string) error {
	url := fmt.Sprintf(deleteFileURL, fileID)
	postRequest, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	postRequest.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
	postRequest.Header.Add("Accept", "application/json")

	_, err = http_utils.SendHttpRequest(postRequest, 204)
	if err != nil {
		return err
	}
	return nil
}

func UpdateFile(accessToken, fileID string, file DriveFile) (*DriveFile, error) {
	url := fmt.Sprintf(updateFileURL, fileID)
	b, err := json.Marshal(&file)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to convert the content to bytes"))
	}

	postRequest, err := http.NewRequest("PATCH", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	postRequest.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
	postRequest.Header.Add("Accept", "application/json")

	resp, err := http_utils.SendHttpRequest(postRequest, 200)
	if err != nil {
		return nil, err
	}
	drive, err := utils.IoReaderToStruct[DriveFile](&resp.Body)
	if err != nil {
		return nil, err
	}
	return drive, nil
}
