//
// EPITECH PROJECT, 2025
// AREA
// File description:
// oauthAsana
//

package oauth

import (
	http_utils "area/utils/httpUtils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/go-chi/jwtauth/v5"
)

type AsanaOAuth struct {
	AccessTokenURL  string
	EmailRequestURL string
}

type asanaAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
}

type asanaUserInfo struct {
	Email string `json:"email"`
}

type asanaResponseUser struct {
	Data asanaUserInfo `json:"data,omitempty"`
}

func (Asana *AsanaOAuth) GetAccessToken(OAuthCode *OAuthRequest) (*OAuthAccessInfos, error) {
	values := url.Values{}
	values.Set("grant_type", "authorization_code")
	values.Set("client_id", os.Getenv(fmt.Sprintf("%s_ID", strings.ToUpper(OAuthCode.Service))))
	values.Set("client_secret", os.Getenv(fmt.Sprintf("%s_SECRET", strings.ToUpper(OAuthCode.Service))))
	values.Set("code", OAuthCode.Code)
	values.Set("redirect_uri", OAuthCode.RedirectURI)

	response, resperr := http.PostForm(Asana.AccessTokenURL, values)

	if resperr != nil {
		return nil, resperr
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var AsanaOAuthResponse asanaAccessTokenResponse
	err = json.Unmarshal(body, &AsanaOAuthResponse)
	if err != nil {
		return nil, err
	}
	return &OAuthAccessInfos{
		AccessToken: AsanaOAuthResponse.AccessToken,
		Scope:       AsanaOAuthResponse.Scope,
	}, nil
}

func (Asana *AsanaOAuth) HandleUserTokens(oauthInfo OAuthAccessInfos, w *http.ResponseWriter, JwtTok *jwtauth.JWTAuth) error {
	request, _ := http.NewRequest("GET", Asana.EmailRequestURL, nil) // GetUserMail request
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", oauthInfo.AccessToken))
	request.Header.Set("Accept", "application/json")

	result, err := http_utils.SendHttpRequest(request, 200)
	if err != nil {
		return err
	}
	var AsanaUserInfo asanaResponseUser
	err = json.NewDecoder(result.Body).Decode(&AsanaUserInfo)
	if err != nil {
		log.Println("Error on email get infos", err)
		return err
	}
	user, err := CreateUserWithEmail(*w, AsanaUserInfo.Data.Email, JwtTok)
	if err == nil { // if no error was found it means there is a new user
		CreateToken(*w, user, oauthInfo.AccessToken, "asana") // If the token just need to be updated then
	}
	return nil
}
