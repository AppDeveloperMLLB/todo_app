package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/AppDeveloperMLLB/todo_app/apperrors"
	"github.com/AppDeveloperMLLB/todo_app/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
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
func HandleGoogleCallback(db *sql.DB, state string, code string) (models.User, error) {
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

	bearerToken, ok := token.Extra("id_token").(string)
	if !ok {
		log.Println("No id_token field in oauth2 token.")
		return models.User{}, apperrors.NoIdToken.Wrap(nil, "No id_token field in oauth2 token")
	}

	tokenValidator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		err = apperrors.CreateValidatorFailed.Wrap(
			err,
			"internal auth error",
		)
		return models.User{}, err
	}

	clientID := os.Getenv("CLIENT_ID")
	payload, err := tokenValidator.Validate(context.Background(), bearerToken, clientID)
	if err != nil {
		log.Println("validate failed")
		err = apperrors.Unauthorized.Wrap(err, "invalid token")
		return models.User{}, err
	}

	googleUserID, ok := payload.Claims["sub"].(string)
	if !ok {
		err = apperrors.Unauthorized.Wrap(errors.New("invalid token"), "invalid token")
		return models.User{}, err
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
	// userinfoのIDとTokenから取得するIDが異なるため、Tokenから取得したIDを使用する
	user.ID = googleUserID
	user.Token = bearerToken

	err = updateUser(db, user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func updateUser(db *sql.DB, user models.User) error {
	getUserSQL := "SELECT * FROM users WHERE google_user_id = $1;"

	var userID int
	var existUser models.User
	var createdAt sql.NullTime
	var updatedAt sql.NullTime
	err := db.QueryRow(getUserSQL, user.ID).Scan(
		&userID,
		&existUser.ID,
		&existUser.Email,
		&existUser.Name,
		&existUser.Picture,
		&createdAt,
		&updatedAt,
	)
	// エラーがなく、ユーザーのデータが存在するのでなにもしない
	if err == nil {
		return nil
	}

	// データがないエラーではなく、別のエラーが発生した場合はエラーを返す
	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	// ユーザーが存在しない場合は新規登録する
	insertUserSQL := `
		INSERT INTO users (
			google_user_id, email, username, profile_picture, created_at, updated_at
		) VALUES (
		 	$1, $2, $3, $4, NOW(), NOW()
		) RETURNING id;
	`

	var newID int
	err = db.QueryRow(insertUserSQL, user.ID, user.Email, user.Name, user.Picture).Scan(&newID)
	if err != nil {
		return err
	}

	return nil
}
