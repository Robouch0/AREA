//
// EPITECH PROJECT, 2025
// AREA
// File description:
// oauthMiro
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
	requestBodyMap := map[string]string{
		"client_id":     os.Getenv(fmt.Sprintf("%s_ID", strings.ToUpper(OAuthCode.Service))),
		"client_secret": os.Getenv(fmt.Sprintf("%s_SECRET", strings.ToUpper(OAuthCode.Service))),
		"code":          OAuthCode.Code,
		"redirect_uri":  OAuthCode.RedirectURI,
		"grant_type":    "authorization_code",
	}

	resp, err := AccessTokenPost(OAuthCode, Miro.AccessTokenURL, requestBodyMap, NewContentTypeHeader("application/json"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var MiroOAuthResponse miroAccessTokenResponse
	err = json.Unmarshal(body, &MiroOAuthResponse)
	if err != nil {
		return nil, err
	}
	return &OAuthAccessInfos{
		AccessToken: MiroOAuthResponse.AccessToken,
		Scope:       MiroOAuthResponse.Scope,
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
