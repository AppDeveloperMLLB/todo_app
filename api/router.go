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
	r.HandleFunc("/v1/todo", tCon.TodoListHandler).Methods(http.MethodGet)
	r.HandleFunc("/callback", aCon.CallbackHandler).Methods(http.MethodGet)
	// r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	// r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	// r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleHandler).Methods(http.MethodGet)
	// r.HandleFunc("/article/nice", aCon.ArticleNiceHandler).Methods(http.MethodPost)
	// r.HandleFunc("/comment", cCon.CommentHandler).Methods(http.MethodPost)
	// r.HandleFunc("/login", authCon.LoginHandler).Methods(http.MethodGet)

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

// func handleTodo(w http.ResponseWriter, r *http.Request) {
// 	queryMap := r.URL.Query()
// 	var page int
// 	if p, ok := queryMap["page"]; ok && len(p) > 0 {
// 		var err error
// 		page, err = strconv.Atoi(p[0])
// 		if err != nil {
// 			err = apperrors.BadParam.Wrap(err, "queryparam must be number")
// 			apperrors.ErrorHandler(w, r, err)
// 			return
// 		}
// 	} else {
// 		page = 1
// 	}

// 	var perPage int
// 	if perPageQuery, ok := queryMap["per_page"]; ok && len(perPageQuery) > 0 {
// 		var err error
// 		perPage, err = strconv.Atoi(perPageQuery[0])
// 		if err != nil {
// 			err = apperrors.BadParam.Wrap(err, "queryparam must be number")
// 			apperrors.ErrorHandler(w, r, err)
// 			return
// 		}
// 	} else {
// 		perPage = 10
// 	}

// 	var status string
// 	if statusQuery, ok := queryMap["status"]; ok && len(statusQuery) > 0 {
// 		status = statusQuery[0]
// 		if status != "in_progress" && status != "todo" && status != "completed" {

// 		}
// 	} else {
// 		status = "all"
// 	}

// 	println(page)
// 	println(perPage)
// 	println(status)

// 	// articleList, err := c.service.GetArticleListService(page)
// 	// if err != nil {
// 	// 	apperrors.ErrorHandler(w, req, err)
// 	// 	return
// 	// }
// 	// var articleTestData = []models.Article{
// 	// 	models.Article{
// 	// 		ID:          1,
// 	// 		UserID:      "0001",
// 	// 		Title:       "FirstPost",
// 	// 		Contents:    "This is FirstPost",
// 	// 		UserName:    "sak",
// 	// 		NiceNum:     2,
// 	// 		CommentList: commentTestData,
// 	// 	},
// 	// 	models.Article{
// 	// 		ID:       2,
// 	// 		UserID:   "0001",
// 	// 		Title:    "SecondPost",
// 	// 		Contents: "This is SecondPost",
// 	// 		UserName: "sak",
// 	// 		NiceNum:  3,
// 	// 	},
// 	// }

// 	var res = models.TodoResponse{
// 		Todos: []models.Todo{
// 			models.Todo{
// 				ID:          1,
// 				UserID:      1,
// 				Title:       "FirstPost",
// 				Description: "This is FirstPost",
// 				Status:      "todo",
// 				CreatedAt:   time.Now(),
// 				UpdatedAt:   time.Now(),
// 			},
// 			models.Todo{
// 				ID:          2,
// 				UserID:      1,
// 				Title:       "SecondPost",
// 				Description: "This is SecondPost",
// 				Status:      "completed",
// 				CreatedAt:   time.Now(),
// 				UpdatedAt:   time.Now(),
// 			},
// 		},
// 		Pages: models.Pages{
// 			Total:   2,
// 			Current: 1,
// 			Next:    2,
// 			Prev:    0,
// 		},
// 	}

// 	json.NewEncoder(w).Encode(res)
// }
