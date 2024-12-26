//
// EPITECH PROJECT, 2024
// AREA
// File description:
// sign-oauth
//

package controllers

import (
	"area/api/controllers/oauth"
	"area/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func OAuthRoutes(JwtTok *jwtauth.JWTAuth) chi.Router {
	OAuthRouter := chi.NewRouter()
	OAuthURLs := oauth.CreateOAuthURLS()
	db.InitTokenDb()
	OAuthRouter.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(JwtTok))
		r.Use(jwtauth.Authenticator(JwtTok))
		r.Post("/connect/", oauth.Connect(OAuthURLs))
	})

	OAuthRouter.Get("/{service}", oauth.GetUrl(OAuthURLs))

	OAuthRouter.Post("/", oauth.LoginOAuth(JwtTok, OAuthURLs))
	return OAuthRouter
}
