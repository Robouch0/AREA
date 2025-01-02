//
// EPITECH PROJECT, 2024
// AREA
// File description:
// sharedDriveHandler
//

package drive

import (
	"area/utils"
	http_utils "area/utils/httpUtils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IT IS NOT FREE : (

const (
	createDriveURL = "https://www.googleapis.com/drive/v3/drives"
	listDriveURL   = "https://www.googleapis.com/drive/v3/drives"
	deleteDriveURL = "https://www.googleapis.com/drive/v3/drives/%s"
	patchDriveURL  = "https://www.googleapis.com/drive/v3/drives/%s"
)

type ListSharedDrive struct {
	NextPageToken string        `json:"nextPageToken,omitempty"`
	Kind          string        `json:"kind,omitempty"`
	Drives        []SharedDrive `json:"drives,omitempty"`
}

func ListSharedDriveHandler(accessToken string, useDomainAdminAccess bool) (*ListSharedDrive, error) {
	request, err := http.NewRequest("GET", listDriveURL, nil)
	if err != nil {
		return nil, err
	}
	request.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
	request.Header.Add("Accept", "application/json")
	q := request.URL.Query()
	q.Set("useDomainAdminAccess", strconv.FormatBool(useDomainAdminAccess))
	request.URL.RawQuery = q.Encode()

	resp, err := http_utils.SendHttpRequest(request, 200)
	if err != nil {
		return nil, err
	}
	list, err := utils.IoReaderToStruct[ListSharedDrive](&resp.Body)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func CreateSharedDrive(accessToken string, sharedDrive SharedDrive) (*SharedDrive, error) {
	b, err := json.Marshal(&sharedDrive)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to convert the content to bytes"))
	}

	postRequest, err := http.NewRequest("POST", createDriveURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	postRequest.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
	postRequest.Header.Add("Accept", "application/json")
	q := postRequest.URL.Query()
	q.Set("requestId", "Area-"+uuid.NewString())
	postRequest.URL.RawQuery = q.Encode()

	resp, err := http_utils.SendHttpRequest(postRequest, 200)
	if err != nil {
		return nil, err
	}
	drive, err := utils.IoReaderToStruct[SharedDrive](&resp.Body)
	if err != nil {
		return nil, err
	}
	return drive, nil
}

func DeleteSharedDrive(accessToken, driveName string, useDomainAdminAccess bool) error {
	list, err := ListSharedDriveHandler(accessToken, useDomainAdminAccess)
	if err != nil {
		return err
	}
	for _, d := range list.Drives {
		if d.Name == driveName {
			url := fmt.Sprintf(deleteDriveURL, d.ID)
			request, err := http.NewRequest("DELETE", url, nil)
			if err != nil {
				return err
			}
			request.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
			request.Header.Add("Accept", "application/json")
			q := request.URL.Query()
			q.Set("useDomainAdminAccess", strconv.FormatBool(useDomainAdminAccess))
			request.URL.RawQuery = q.Encode()

			_, err = http_utils.SendHttpRequest(request, 200)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return status.Errorf(codes.InvalidArgument, "No such drive with name %v", driveName)
}

func UpdateSharedDrive(accessToken, driveName string, putDrive SharedDrive, useDomainAdminAccess bool) (*SharedDrive, error) {
	list, err := ListSharedDriveHandler(accessToken, useDomainAdminAccess)
	if err != nil {
		return nil, err
	}
	for _, d := range list.Drives {
		if d.Name == driveName {
			url := fmt.Sprintf(patchDriveURL, d.ID)
			b, err := json.Marshal(&putDrive)
			if err != nil {
				return nil, err
			}
			putRequest, err := http.NewRequest("PATCH", url, bytes.NewBuffer(b))
			if err != nil {
				return nil, err
			}
			putRequest.Header = http_utils.GetDefaultBearerHTTPHeader(accessToken)
			putRequest.Header.Add("Accept", "application/json")
			resp, err := http_utils.SendHttpRequest(putRequest, 200)
			if err != nil {
				return nil, err
			}
			return utils.IoReaderToStruct[SharedDrive](&resp.Body)
		}
	}
	return nil, status.Errorf(codes.NotFound, "Did not found label named: %v", driveName)
}
