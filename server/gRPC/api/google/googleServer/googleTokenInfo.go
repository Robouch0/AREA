//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleTokenInfo
//

package google_server

import (
	"area/utils"
	http_utils "area/utils/httpUtils"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	tokenInfoURL = "https://oauth2.googleapis.com/tokeninfo"
)

type GoogleTokenInfo struct {
	// First contents like iss, sub are ignored feel free to add it.

	Email         string `json:"email,omitempty"`
	EmailVerified bool   `json:"email_verified,omitempty"`
}

func GetTokenInfo(token string) (*GoogleTokenInfo, error) {
	req, err := http.NewRequest("GET", tokenInfoURL, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Invalid request: %v", err)
	}
	q := req.URL.Query()
	q.Add("id_token", token)
	req.URL.RawQuery = q.Encode()

	resp, err := http_utils.SendHttpRequest(req, 200)
	if err != nil {
		return nil, err
	}
	gTokInfo, err := utils.IoReaderToStruct[GoogleTokenInfo](&resp.Body)
	if err != nil {
		return nil, err
	}
	return gTokInfo, nil
}
