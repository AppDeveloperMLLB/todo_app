package services

import (
	"github.com/AppDeveloperMLLB/todo_app/models"
)

// TodoService - Todoに関するサービスのインターフェース
type TodoService interface {
	GetTodoService(todoID int) (models.Todo, error)
	GetTodoListService(page int, perPage int) ([]models.Todo, error)
	CreateTodoService(todo models.Todo) (models.Todo, error)
	UpdateTodoService(todo models.Todo) (models.Todo, error)
	DeleteTodoService(todoID int, userID string) error
}

// AuthService - 認証に関するサービスのインターフェース
type AuthService interface {
	LoginService() string
	GoogleCallbackService(state string, code string) (models.User, error)
}
