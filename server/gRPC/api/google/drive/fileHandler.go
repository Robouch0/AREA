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
	"net/http"
)

const (
	listFilesURL = "https://www.googleapis.com/drive/v3/files"
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
