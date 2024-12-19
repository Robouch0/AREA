//
// EPITECH PROJECT, 2024
// AREA
// File description:
// oauthDiscord
//

package oauth

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

type DiscordOAuth struct {
	AccessTokenURL  string
	EmailRequestURL string
}

func (discord *DiscordOAuth) GetAccessToken(OAuthCode *OAuthRequest) (*OAuthAccessInfos, error) {
	return nil, nil
}

func (discord *DiscordOAuth) HandleUserTokens(oauthInfo OAuthAccessInfos, w *http.ResponseWriter, JwtTok *jwtauth.JWTAuth) error {
	return nil
}
