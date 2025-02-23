//
// EPITECH PROJECT, 2024
// AREA
// File description:
// oauthTypes
//

package oauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/jwtauth/v5"
)

type OAuthRequest struct {
	Service     string `json:"service"`
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
}

type OAuthURLs struct {
	RedirectURL string        `json:"redirect_uri"`
	OAuth       OAuthProvider `json:"oauth_provider"`
}

type OAuthAccessInfos struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
}

type OAuthProvider interface {
	GetAccessToken(OAuthCode *OAuthRequest) (*OAuthAccessInfos, error)

	// Take the oauthInfo, the responseWritter and the JWT Token to get the informations
	// of the new user and store it in Database.
	HandleUserTokens(oauthInfo OAuthAccessInfos, w *http.ResponseWriter, JwtTok *jwtauth.JWTAuth) error
}

func NewContentTypeHeader(contentType string) http.Header {
	h := http.Header{}
	h.Set("Content-Type", contentType)
	return h
}

func AccessTokenPost(
	OAuthCode *OAuthRequest,
	accessTokenURI string,
	requestBodyMap map[string]string,
	header http.Header,
) (*http.Response, error) {
	requestJSON, err := json.Marshal(requestBodyMap)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", accessTokenURI, bytes.NewBuffer(requestJSON))
	if err != nil {
		return nil, err
	}
	req.Header = header
	req.Header.Set("Accept", "application/json")

	response, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		return nil, resperr
	}
	if response.StatusCode != 200 {
		io.Copy(os.Stderr, response.Body)
		return nil, errors.New(response.Status)
	}
	return response, nil
}
