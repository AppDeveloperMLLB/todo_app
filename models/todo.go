package models

import "time"

// Todo - Todoの構造体
type Todo struct {
	ID          int       `json:"id"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TodoResponse - Todoのレスポンス構造体
type TodoResponse struct {
	Todos []Todo `json:"todos"`
	Pages Pages  `json:"pages"`
}

// Pages - ページング情報
type Pages struct {
	Total   int `json:"total"`
	Current int `json:"current"`
	Next    int `json:"next"`
	Prev    int `json:"prev"`
}
