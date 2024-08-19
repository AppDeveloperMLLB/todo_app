package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/AppDeveloperMLLB/todo_app/api/middlewares"
	"github.com/AppDeveloperMLLB/todo_app/controllers"
	"github.com/AppDeveloperMLLB/todo_app/services"
	"github.com/gorilla/mux"
)

// NewRouter - ルーティングの設定
func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	tCon := controllers.NewTodoController(ser)
	aCon := controllers.NewAuthController(ser)
	//aCon := controllers.NewArticleController(ser)
	//cCon := controllers.NewCommentController(ser)
	//authCon := controllers.NewAuthController(ser)
	r := mux.NewRouter()
	// HTML
	r.HandleFunc("/", handleMain).Methods(http.MethodGet)

	// API
	r.HandleFunc("/login", aCon.LoginHandler).Methods(http.MethodGet)
	r.HandleFunc("/callback", aCon.CallbackHandler).Methods(http.MethodGet)
	r.HandleFunc("/v1/todo/{todo_id:[0-9]+}", tCon.GetTodoHandler).Methods(http.MethodGet)
	r.HandleFunc("/v1/todo/{todo_id:[0-9]+}", tCon.DeleteTodoHandler).Methods(http.MethodDelete)
	r.HandleFunc("/v1/todo", tCon.TodoListHandler).Methods(http.MethodGet)
	r.HandleFunc("/v1/todo", tCon.CreateTodoHandler).Methods(http.MethodPost)
	r.HandleFunc("/v1/todo", tCon.UpdateTodoHandler).Methods(http.MethodPut)

	// 404エラーハンドラー
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 not found", http.StatusNotFound)
	})
	// 405エラーハンドラー
	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
	})

	r.Use(middlewares.LoggingMiddleware)
	r.Use(middlewares.AuthMiddleware)

	return r
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html><body><a href="/login">Google Log In</a></body></html>`
	fmt.Fprint(w, htmlIndex)
}
