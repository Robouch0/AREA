//
// EPITECH PROJECT, 2025
// AREA
// File description:
// oauthMiro
//

package oauth

import (
	"area/utils"
	http_utils "area/utils/httpUtils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/jwtauth/v5"
)

type MiroOAuth struct {
	AccessTokenURL  string
	EmailRequestURL string
}

type miroAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
}

type miroUserInfo struct {
	Email string `json:"email"`
}

func (Miro *MiroOAuth) GetAccessToken(OAuthCode *OAuthRequest) (*OAuthAccessInfos, error) {
	req, err := http.NewRequest("POST", Miro.AccessTokenURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	q := req.URL.Query()
	q.Set("client_id", os.Getenv(fmt.Sprintf("%s_ID", strings.ToUpper(OAuthCode.Service))))
	q.Set("client_secret", os.Getenv(fmt.Sprintf("%s_SECRET", strings.ToUpper(OAuthCode.Service))))
	q.Set("code", OAuthCode.Code)
	q.Set("redirect_uri", OAuthCode.RedirectURI)
	q.Set("grant_type", "authorization_code")
	req.URL.RawQuery = q.Encode()

	resp, err := http_utils.SendHttpRequest(req, 200)
	if err != nil {
		return nil, err
	}
	miroOAuthResponse, err := utils.IoReaderToStruct[miroAccessTokenResponse](&resp.Body)
	if err != nil {
		return nil, err
	}
	return &OAuthAccessInfos{
		AccessToken: miroOAuthResponse.AccessToken,
		Scope:       miroOAuthResponse.Scope,
	}, nil
}

func (Miro *MiroOAuth) HandleUserTokens(oauthInfo OAuthAccessInfos, w *http.ResponseWriter, JwtTok *jwtauth.JWTAuth) error {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", Miro.EmailRequestURL, nil) // GetUserMail request
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", oauthInfo.AccessToken))
	request.URL.Query().Add("personFields", "emailAddresses")

	result, err := client.Do(request)
	if err != nil {
		return err
	}
	if result.StatusCode != 200 {
		log.Println(result.Status)
		return errors.New("Invalid status of the request")
	}
	var miroUserInfo miroUserInfo
	err = json.NewDecoder(result.Body).Decode(&miroUserInfo)
	if err != nil {
		return err
	}
	user, err := CreateUserWithEmail(*w, miroUserInfo.Email, JwtTok)
	if err == nil { // if no error was found it means there is a new user
		CreateToken(*w, user, oauthInfo.AccessToken, "miro") // If the token just need to be updated then
	}
	return nil
}
