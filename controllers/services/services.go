package services

import (
	"github.com/AppDeveloperMLLB/todo_app/models"
)

// TodoService - Todoに関するサービスのインターフェース
type TodoService interface {
	//PostArticleService(article models.Article) (models.Article, error)
	//GetArticleListService(page int) ([]models.Article, error)
	GetTodoListService(page int, perPage int) ([]models.Todo, error)
	//PostNiceService(article models.Article) (models.Article, error)
}

// AuthService - 認証に関するサービスのインターフェース
type AuthService interface {
	LoginService() string
	GoogleCallbackService(state string, code string) (models.User, error)
}
