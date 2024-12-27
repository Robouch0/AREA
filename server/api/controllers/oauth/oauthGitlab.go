//
// EPITECH PROJECT, 2024
// AREA
// File description:
// oauthGitlab
//

package oauth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/jwtauth/v5"
)

type GitlabOAuth struct {
	AccessTokenURL  string
	EmailRequestURL string
}

type gitlabAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
}

type gitlabUser struct {
	Email string `json:"email"`
}

func (git *GitlabOAuth) GetAccessToken(OAuthCode *OAuthRequest) (*OAuthAccessInfos, error) {
	requestBodyMap := map[string]string{
		"client_id":     os.Getenv(fmt.Sprintf("%s_ID", strings.ToUpper(OAuthCode.Service))),
		"client_secret": os.Getenv(fmt.Sprintf("%s_SECRET", strings.ToUpper(OAuthCode.Service))),
		"code":          OAuthCode.Code,
		"redirect_uri":  OAuthCode.RedirectURI,
		"grant_type":    "authorization_code",
	}

	resp, err := AccessTokenPost(OAuthCode, git.AccessTokenURL, requestBodyMap, NewContentTypeHeader("application/json"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var gitOAuthResponse gitlabAccessTokenResponse
	err = json.Unmarshal(body, &gitOAuthResponse)
	if err != nil {
		return nil, err
	}
	return &OAuthAccessInfos{
		AccessToken: gitOAuthResponse.AccessToken,
		Scope:       gitOAuthResponse.Scope,
	}, nil
}

func (git *GitlabOAuth) HandleUserTokens(oauthInfo OAuthAccessInfos, w *http.ResponseWriter, JwtTok *jwtauth.JWTAuth) error {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", git.EmailRequestURL, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", oauthInfo.AccessToken))

	q := request.URL.Query()
	q.Set("access_token", oauthInfo.AccessToken)
	request.URL.RawQuery = q.Encode()

	result, err := client.Do(request)
	if err != nil {
		return err
	}

	var gitUserInfos gitlabUser
	err = json.NewDecoder(result.Body).Decode(&gitUserInfos)
	if err != nil {
		log.Println(err)
		return err
	}
	user, err := CreateUserWithEmail(*w, gitUserInfos.Email, JwtTok)
	if err == nil { // if no error was found it means there is a new user
		CreateToken(*w, user, oauthInfo.AccessToken, "gitlab") // If the token just need to be updated then ?
	}
	return nil
}
