package repositories

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/AppDeveloperMLLB/todo_app/apperrors"
	"github.com/AppDeveloperMLLB/todo_app/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	authorizeEndpoint = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenEndpoint     = "https://www.googleapis.com/oauth2/v4/token"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
	oauthStateString = "random" // ランダムな文字列を使用してCSRFを防止
)

// GetAuthCodeURL - 認証コードのURLを取得する
func GetAuthCodeURL() string {
	oauthConf.ClientID = os.Getenv("CLIENT_ID")
	oauthConf.ClientSecret = os.Getenv("CLIENT_SECRET")
	return oauthConf.AuthCodeURL(oauthStateString)
}

// HandleGoogleCallback - Googleのコールバック
func HandleGoogleCallback(state string, code string) (models.User, error) {
	if state != oauthStateString {
		log.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		return models.User{}, apperrors.InvalidOauthState.Wrap(nil, "invalid oauth state")
	}

	token, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		//http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return models.User{}, apperrors.OauthConfExchangeFailed.Wrap(err, "oauthConf.Exchange() failed")
	}

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		log.Println("No id_token field in oauth2 token.")
		return models.User{}, apperrors.NoIdToken.Wrap(nil, "No id_token field in oauth2 token")
	}

	client := oauthConf.Client(context.Background(), token)
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Printf("client.Get() failed with '%s'\n", err)
		return models.User{}, apperrors.ClientGetFailed.Wrap(err, "client.Get() failed")
	}
	defer userInfo.Body.Close()

	var user models.User
	json.NewDecoder(userInfo.Body).Decode(&user)
	user.Token = idToken
	return user, nil
}
