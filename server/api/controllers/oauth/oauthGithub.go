//
// EPITECH PROJECT, 2024
// AREA
// File description:
// oauthGithub
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

type GithubOAuth struct {
	AccessTokenURL  string
	EmailRequestURL string
}

type githubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
}

type githubUser struct {
	Email string `json:"email"`
}

func (git *GithubOAuth) GetAccessToken(OAuthCode *OAuthRequest) (*OAuthAccessInfos, error) {
	requestBodyMap := map[string]string{
		"client_id":     os.Getenv(fmt.Sprintf("%s_ID", strings.ToUpper(OAuthCode.Service))),
		"client_secret": os.Getenv(fmt.Sprintf("%s_SECRET", strings.ToUpper(OAuthCode.Service))),
		"code":          OAuthCode.Code,
		"redirect_uri":  OAuthCode.RedirectURI,
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
	var gitOAuthResponse githubAccessTokenResponse
	err = json.Unmarshal(body, &gitOAuthResponse)
	if err != nil {
		return nil, err
	}
	return &OAuthAccessInfos{
		AccessToken: gitOAuthResponse.AccessToken,
		Scope:       gitOAuthResponse.Scope,
	}, nil
}

func (git *GithubOAuth) HandleUserTokens(oauthInfo OAuthAccessInfos, w *http.ResponseWriter, JwtTok *jwtauth.JWTAuth) error {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", git.EmailRequestURL, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", oauthInfo.AccessToken))

	result, err := client.Do(request)

	if err != nil {
		return err
	}
	var gitUserInfos []githubUser
	err = json.NewDecoder(result.Body).Decode(&gitUserInfos)
	if err != nil {
		log.Println(err)
		return err
	}
	if len(gitUserInfos) == 0 {
		return errors.New("No user emails given from github")
	}
	user, err := CreateUserWithEmail(*w, gitUserInfos[0].Email, JwtTok)
	if err == nil { // if no error was found it means there is a new user
		CreateToken(*w, user, oauthInfo.AccessToken, "github") // If the token just need to be updated then ?
	}
	return nil
}
