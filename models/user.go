package models

// User - ユーザの構造体
type User struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Token   string `json:"token"`
}
