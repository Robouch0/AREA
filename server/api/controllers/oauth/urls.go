//
// EPITECH PROJECT, 2024
// AREA
// File description:
// urls
//

package oauth

import (
	http_utils "area/utils/httpUtils"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const (
	googleScopes string = "https://www.googleapis.com/auth/gmail.modify " +
		"https://www.googleapis.com/auth/userinfo.email " +
		"https://mail.google.com/ " +
		"https://www.googleapis.com/auth/drive " +
		"https://www.googleapis.com/auth/drive.file " +
		"https://www.googleapis.com/auth/drive.appdata"
	hfScopes     = "openid profile email read-repos write-repos manage-repos write-discussions read-billing"
	gitlabScopes = "api read_api read_user read_repository write_repository openid profile email"
	miroScopes   = "boards-read boards-write identity:read identity:write team:read team:write"
)

func createOAuthRedirect(consentURL, provider, redirectURI string) string {
	return fmt.Sprintf(consentURL, os.Getenv(fmt.Sprintf("%s_ID", strings.ToUpper(provider))), redirectURI)
}

func CreateOAuthURLS() map[string]OAuthURLs {
	oauthUrls := make(map[string]OAuthURLs)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	oauthUrls["github"] = OAuthURLs{
		RedirectURL: "https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&scope=user:email,repo,workflow&state=github", // Mettre plus de droits précis
		OAuth: &GithubOAuth{
			AccessTokenURL:  "https://github.com/login/oauth/access_token",
			EmailRequestURL: "https://api.github.com/user/emails",
		},
	}
	oauthUrls["google"] = OAuthURLs{
		RedirectURL: fmt.Sprintf(
			"https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&redirect_uri=%s&scope=%s&response_type=code&state=google", "%s", "%s", googleScopes),
		OAuth: &GoogleOAuth{
			AccessTokenURL:  "https://oauth2.googleapis.com/token",
			EmailRequestURL: "https://www.googleapis.com/oauth2/v1/userinfo",
		},
	}
	oauthUrls["discord"] = OAuthURLs{
		RedirectURL: "https://discord.com/oauth2/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=identify+email+bot+applications.commands+guilds&state=discord",
		OAuth: &DiscordOAuth{
			AccessTokenURL:  "https://discord.com/api/oauth2/token",
			EmailRequestURL: "https://discord.com/api/users/@me",
		},
	}
	oauthUrls["spotify"] = OAuthURLs{
		RedirectURL: `https://accounts.spotify.com/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=user-modify-playback-state playlist-modify-public playlist-modify-private&state=spotify`, // Maybe add state later
		OAuth: &SpotifyOAuth{
			AccessTokenURL:  "https://accounts.spotify.com/api/token",
			EmailRequestURL: "https://api.spotify.com/v1/me",
		},
	}
	oauthUrls["hf"] = OAuthURLs{
		RedirectURL: fmt.Sprintf(
			`https://huggingface.co/oauth/authorize?client_id=%s&response_type=code&prompt=consent&redirect_uri=%s&scope=%s&state=hf`, "%s", "%s", hfScopes),
		OAuth: &HFOAuth{
			AccessTokenURL:  "https://huggingface.co/oauth/token",
			EmailRequestURL: "https://huggingface.co/api/whoami-v2",
		},
	}
	oauthUrls["gitlab"] = OAuthURLs{
		RedirectURL: fmt.Sprintf(
			`https://gitlab.com/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code&state=gitlab&scope=%s`, "%s", "%s", gitlabScopes),
		OAuth: &GitlabOAuth{
			AccessTokenURL:  "https://gitlab.com/oauth/token",
			EmailRequestURL: "https://gitlab.com/oauth/userinfo",
		},
	}
	oauthUrls["asana"] = OAuthURLs{
		RedirectURL: "https://app.asana.com/-/oauth_authorize?response_type=code&client_id=%s&code_challenge_method=S256&code_challenge=test&redirect_uri=%s&state=asana&scope=default", // Normally default is enough
		OAuth: &AsanaOAuth{
			AccessTokenURL:  "https://app.asana.com/-/oauth_token",
			EmailRequestURL: "https://app.asana.com/api/1.0/users/me",
		},
	}
	oauthUrls["miro"] = OAuthURLs{
		RedirectURL: fmt.Sprintf(
			"https://miro.com/oauth/authorize?response_type=code&client_id=%s&redirect_uri=%s&state=miro&scope=%s", "%s", "%s", miroScopes),
		OAuth: &MiroOAuth{
			AccessTokenURL:  "https://api.miro.com/v1/oauth/token",
			EmailRequestURL: "https://api.miro.com/v1/users/me",
		},
	}
	return oauthUrls
}

// Sign-up OAuth godoc
//
// @Summary      Get an oauth url for a service
// @Description  Get the oauth redirect url for a service
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param redirect_uri query string true "Redirect URL for the oauth"
// @Param service path string true "Name of the service to use oauth with"
// @Success      200  {object}  string
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Router       /oauth/{service} [get]
func GetUrl(OAuthURLs map[string]OAuthURLs) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		OAuthservice := chi.URLParam(r, "service")
		log.Println(OAuthservice)
		OAuthRedirectURI := r.URL.Query().Get("redirect_uri") // Rajouter dans le swagger les urls params

		log.Println("Service: ", OAuthservice)
		if serviceUrls, ok := OAuthURLs[OAuthservice]; ok {
			url := createOAuthRedirect(serviceUrls.RedirectURL, OAuthservice, OAuthRedirectURI)
			log.Println(url)
			if url != "" {
				w.WriteHeader(200)
				w.Write([]byte(url))
			} else {
				http_utils.WriteHTTPResponseErr(&w, 400, "Service does not exist")
			}
			return
		}
		http_utils.WriteHTTPResponseErr(&w, 404, fmt.Sprintf("Service %s not found", OAuthservice))
	}
}
