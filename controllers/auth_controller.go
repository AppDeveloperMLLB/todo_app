package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/AppDeveloperMLLB/todo_app/apperrors"
	"github.com/AppDeveloperMLLB/todo_app/controllers/services"
)

// AuthController - 認証に関するコントローラ
type AuthController struct {
	service services.AuthService
}

// NewAuthController - AuthControllerのコンストラクタ
func NewAuthController(s services.AuthService) *AuthController {
	return &AuthController{service: s}
}

// LoginHandler - GET /loginのハンドラ
func (c *AuthController) LoginHandler(w http.ResponseWriter, req *http.Request) {
	url := c.service.LoginService()
	http.Redirect(w, req, url, http.StatusTemporaryRedirect)
}

// CallbackHandler - GET /callbackのハンドラ
func (c *AuthController) CallbackHandler(w http.ResponseWriter, req *http.Request) {
	state := req.FormValue("state")
	code := req.FormValue("code")
	str, err := c.service.GoogleCallbackService(state, code)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
	}
	json.NewEncoder(w).Encode(str)
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
