//
// EPITECH PROJECT, 2024
// AREA
// File description:
// oauthDiscord
//

package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/go-chi/jwtauth/v5"
)

type DiscordOAuth struct {
	AccessTokenURL  string
	EmailRequestURL string
}

type discordAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type discordUser struct {
	Email string `json:"email"`
}

func (discord *DiscordOAuth) GetAccessToken(OAuthCode *OAuthRequest) (*OAuthAccessInfos, error) {
	values := url.Values{}
	values.Set("grant_type", "authorization_code")
	values.Set("client_id", os.Getenv(fmt.Sprintf("%s_ID", strings.ToUpper(OAuthCode.Service))))
	values.Set("client_secret", os.Getenv(fmt.Sprintf("%s_SECRET", strings.ToUpper(OAuthCode.Service))))
	values.Set("code", OAuthCode.Code)
	values.Set("redirect_uri", OAuthCode.RedirectURI)

	response, resperr := http.PostForm(discord.AccessTokenURL, values)

	if resperr != nil {
		return nil, resperr
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var discordOAuthTok discordAccessTokenResponse
	err = json.Unmarshal(body, &discordOAuthTok)
	if err != nil {
		return nil, err
	}
	return &OAuthAccessInfos{
		AccessToken: discordOAuthTok.AccessToken,
		Scope:       discordOAuthTok.Scope,
	}, nil
}

func (discord *DiscordOAuth) HandleUserTokens(oauthInfo OAuthAccessInfos, w *http.ResponseWriter, JwtTok *jwtauth.JWTAuth) error {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", discord.EmailRequestURL, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", oauthInfo.AccessToken))

	result, err := client.Do(request)
	if err != nil {
		return err
	}

	var discordUserInfo discordUser // Maybe later fetch username and other
	err = json.NewDecoder(result.Body).Decode(&discordUserInfo)
	if err != nil {
		log.Println(err)
		return err
	}
	user, err := CreateUserWithEmail(*w, discordUserInfo.Email, JwtTok)
	if err == nil { // if no error was found it means there is a new user
		CreateToken(*w, user, oauthInfo.AccessToken, "discord") // If the token just need to be updated then ?
	}
	return nil
}
