//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleTokenInfo
//

package google_server

import (
	conv_utils "area/utils/convUtils"
	http_utils "area/utils/httpUtils"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	tokenInfoURL = "https://www.googleapis.com/oauth2/v1/userinfo"
)

type GoogleTokenInfo struct {
	Email         string `json:"email,omitempty"`
	EmailVerified bool   `json:"email_verified,omitempty"`
}

func GetTokenInfo(token string) (*GoogleTokenInfo, error) {
	req, err := http.NewRequest("GET", tokenInfoURL, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Invalid request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	q := req.URL.Query()
	q.Add("id_token", token)
	q.Add("personFields", "emailAddresses")
	req.URL.RawQuery = q.Encode()

	resp, err := http_utils.SendHttpRequest(req, 200)
	if err != nil {
		return nil, err
	}
	gTokInfo, err := conv_utils.IoReaderToStruct[GoogleTokenInfo](&resp.Body)
	if err != nil {
		return nil, err
	}
	return gTokInfo, nil
}
