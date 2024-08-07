package services

import (
	"os"

	"github.com/AppDeveloperMLLB/todo_app/models"
	"github.com/AppDeveloperMLLB/todo_app/repositories"
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

// LoginService - 認証に関するサービス
func (s *MyAppService) LoginService() string {
	oauthConf.ClientID = os.Getenv("CLIENT_ID")
	oauthConf.ClientSecret = os.Getenv("CLIENT_SECRET")
	return repositories.GetAuthCodeURL()
}

// GoogleCallbackService - Googleのコールバック
func (s *MyAppService) GoogleCallbackService(state string, code string) (models.User, error) {
	return repositories.HandleGoogleCallback(s.db, state, code)
}

// func HandleGoogleCallback(state string, code string) (string, error) {
// 	//state := r.FormValue("state")
// 	if state != oauthStateString {
// 		log.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
// 		//http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return "", apperrors.InvalidOauthState.Wrap(nil, "invalid oauth state")
// 	}

// 	//code := r.FormValue("code")
// 	log.Printf("code: %s\n", code)
// 	token, err := oauthConf.Exchange(context.Background(), code)
// 	if err != nil {
// 		log.Printf("oauthConf.Exchange() failed with '%s'\n", err)
// 		//http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return "", apperrors.OauthConfExchangeFailed.Wrap(err, "oauthConf.Exchange() failed")
// 	}

// 	idToken, ok := token.Extra("id_token").(string)
// 	if !ok {
// 		log.Println("No id_token field in oauth2 token.")
// 		//http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return "", apperrors.NoIdToken.Wrap(nil, "No id_token field in oauth2 token")
// 	}
// 	log.Printf("id_token: %s\n", idToken)

// 	log.Printf("token: %s\n", token.AccessToken)
// 	client := oauthConf.Client(context.Background(), token)
// 	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
// 	if err != nil {
// 		log.Printf("client.Get() failed with '%s'\n", err)
// 		//http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return "", apperrors.ClientGetFailed.Wrap(err, "client.Get() failed")
// 	}
// 	defer userInfo.Body.Close()

// 	data, _ := io.ReadAll(userInfo.Body)
// 	//fmt.Fprintf(w, "UserInfo: %s\n", data)
// 	str := fmt.Sprintf("UserInfo: %s\n", data)
// 	return str, nil
// }
