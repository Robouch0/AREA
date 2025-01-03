//
// EPITECH PROJECT, 2024
// AREA
// File description:
// oauthHF
//

package oauth

import (
	"area/utils"
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/jwtauth/v5"
)

type HFOAuth struct {
	AccessTokenURL  string
	EmailRequestURL string
}

type hfAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type hfUser struct {
	Email string `json:"email"`
}

func (hf *HFOAuth) makeHTTPTokenRequest(OAuthCode *OAuthRequest) (*http.Request, error) {
	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("code", OAuthCode.Code)
	form.Set("redirect_uri", OAuthCode.RedirectURI)

	req, reqError := http.NewRequest("POST", hf.AccessTokenURL, strings.NewReader(form.Encode()))
	if reqError != nil {
		return nil, reqError
	}

	clientID, errID := utils.GetEnvParameter(fmt.Sprintf("%s_ID", strings.ToUpper(OAuthCode.Service)))
	clientSecret, errSecret := utils.GetEnvParameter(fmt.Sprintf("%s_SECRET", strings.ToUpper(OAuthCode.Service)))
	if err := cmp.Or(&errID, &errSecret); *err != nil {
		return nil, *err
	}

	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+utils.EncodeToBase64(clientID+":"+clientSecret))
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (hf *HFOAuth) GetAccessToken(OAuthCode *OAuthRequest) (*OAuthAccessInfos, error) {
	req, err := hf.makeHTTPTokenRequest(OAuthCode)
	if err != nil {
		return nil, err
	}

	response, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		return nil, resperr
	}
	if response.StatusCode != 200 {
		log.Println(response)
		return nil, errors.New(response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var spotifyOAuthTok spotifyAccessTokenResponse
	err = json.Unmarshal(body, &spotifyOAuthTok)
	if err != nil {
		return nil, err
	}
	return &OAuthAccessInfos{
		AccessToken: spotifyOAuthTok.AccessToken,
		Scope:       spotifyOAuthTok.Scope,
	}, nil
}

func (hf *HFOAuth) HandleUserTokens(oauthInfo OAuthAccessInfos, w *http.ResponseWriter, JwtTok *jwtauth.JWTAuth) error {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", hf.EmailRequestURL, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", oauthInfo.AccessToken))

	result, err := client.Do(request)
	if err != nil {
		return err
	}

	var hfUserInfo hfUser // Maybe later fetch username and other
	err = json.NewDecoder(result.Body).Decode(&hfUserInfo)
	if err != nil {
		log.Println(err)
		return err
	}
	user, err := CreateUserWithEmail(*w, hfUserInfo.Email, JwtTok)
	if err == nil { // if no error was found it means there is a new user
		CreateToken(*w, user, oauthInfo.AccessToken, "hf") // If the token just need to be updated then ?
	}
	return nil
}
