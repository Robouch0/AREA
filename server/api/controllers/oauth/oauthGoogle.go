//
// EPITECH PROJECT, 2024
// AREA
// File description:
// getEmail
//

package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/jwtauth/v5"
)

type GoogleOAuth struct {
	AccessTokenURL  string
	EmailRequestURL string
}

type googleUserInfo struct {
	Email string `json:"email"`
}

type googleAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
}

// AccessTokenBody: `{ "client_id" : "%s", "client_secret" : "%s", "code" : "%s", redirect_uri: "%s", grant_type: "authorization_code" }`,

func (google *GoogleOAuth) GetAccessToken(OAuthCode *OAuthRequest) (*OAuthAccessInfos, error) {
	requestBodyMap := map[string]string{
		"client_id":     os.Getenv(fmt.Sprintf("%s_ID", strings.ToUpper(OAuthCode.Service))),
		"client_secret": os.Getenv(fmt.Sprintf("%s_SECRET", strings.ToUpper(OAuthCode.Service))),
		"code":          OAuthCode.Code,
		"redirect_uri":  OAuthCode.RedirectURI,
		"grant_type":    "authorization_code",
	}

	resp, err := AccessTokenPost(OAuthCode, google.AccessTokenURL, requestBodyMap)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var googleOAuthResponse googleAccessTokenResponse
	err = json.Unmarshal(body, &googleOAuthResponse)
	if err != nil {
		return nil, err
	}
	return &OAuthAccessInfos{
		AccessToken: googleOAuthResponse.AccessToken,
		Scope:       googleOAuthResponse.Scope,
	}, nil
}

func (google *GoogleOAuth) HandleUserTokens(oauthInfo OAuthAccessInfos, w *http.ResponseWriter, JwtTok *jwtauth.JWTAuth) error {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", google.EmailRequestURL, nil) // GetUserMail request
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
	var googleUserInfo googleUserInfo
	err = json.NewDecoder(result.Body).Decode(&googleUserInfo)
	if err != nil {
		return err
	}
	user, err := CreateUserWithEmail(*w, googleUserInfo.Email, JwtTok)
	if err == nil { // if no error was found it means there is a new user
		CreateToken(*w, user, oauthInfo.AccessToken, "google") // If the token just need to be updated then
	}
	return nil
}
