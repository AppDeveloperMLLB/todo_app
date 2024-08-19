package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AppDeveloperMLLB/todo_app/apperrors"
	"github.com/AppDeveloperMLLB/todo_app/common"
	"github.com/AppDeveloperMLLB/todo_app/controllers/services"
	"github.com/AppDeveloperMLLB/todo_app/models"
	"github.com/gorilla/mux"
)

// TodoController - Todoに関するコントローラ
type TodoController struct {
	service services.TodoService
}

// NewTodoController - TodoControllerのコンストラクタ
func NewTodoController(s services.TodoService) *TodoController {
	return &TodoController{service: s}
}

// GetTodoHandler - GET /todo/{todo_id}のハンドラ
func (c *TodoController) GetTodoHandler(w http.ResponseWriter, req *http.Request) {
	todoID, err := strconv.Atoi(mux.Vars(req)["todo_id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "path parameter must be number")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	todo, err := c.service.GetTodoService(todoID)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

// TodoListHandler - GET /todoのハンドラ
func (c *TodoController) TodoListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "queryparam must be number")
			apperrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		page = 1
	}

	var perPage int
	if perPageQuery, ok := queryMap["per_page"]; ok && len(perPageQuery) > 0 {
		var err error
		perPage, err = strconv.Atoi(perPageQuery[0])
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "queryparam must be number")
			apperrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		perPage = 10
	}

	todoList, err := c.service.GetTodoListService(page, perPage)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(todoList)
}

// CreateTodoHandler - POST /todoのハンドラ
func (c *TodoController) CreateTodoHandler(w http.ResponseWriter, req *http.Request) {
	var reqTodo models.Todo
	if err := json.NewDecoder(req.Body).Decode(&reqTodo); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	userID := common.GetUserID(req.Context())
	reqTodo.UserID = userID

	todo, err := c.service.CreateTodoService(reqTodo)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (c *TodoController) UpdateTodoHandler(w http.ResponseWriter, req *http.Request) {
	var reqTodo models.Todo
	if err := json.NewDecoder(req.Body).Decode(&reqTodo); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	todo, err := c.service.UpdateTodoService(reqTodo)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (c *TodoController) DeleteTodoHandler(w http.ResponseWriter, req *http.Request) {
	todoID, err := strconv.Atoi(mux.Vars(req)["todo_id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "path parameter must be number")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	userID := common.GetUserID(req.Context())

	err = c.service.DeleteTodoService(todoID, userID)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// // ArticleHandler - GET /article/idのハンドラ
// func (c *ArticleController) ArticleHandler(w http.ResponseWriter, req *http.Request) {
// 	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
// 	if err != nil {
// 		err = apperrors.BadParam.Wrap(err, "path parameter must be number")
// 		apperrors.ErrorHandler(w, req, err)
// 		return
// 	}

// 	article, err := c.service.GetArticleService(articleID)
// 	if err != nil {
// 		apperrors.ErrorHandler(w, req, err)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(article)
// }

// // ArticleNiceHandler - POST /article/niceのハンドラ
// func (c *ArticleController) ArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
// 	var reqArticle models.Article
// 	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
// 		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
// 		apperrors.ErrorHandler(w, req, err)
// 		return
// 	}

// 	if c.isMatchUser(req, reqArticle) {
// 		c.handleNotMatchUser(w, req)
// 		return
// 	}

// 	article, err := c.service.PostNiceService(reqArticle)
// 	if err != nil {
// 		apperrors.ErrorHandler(w, req, err)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(article)
// }

// func (c *ArticleController) isMatchUser(req *http.Request, article models.Article) bool {
// 	authedUserID := common.GetUserID(req.Context())
// 	return article.UserID == authedUserID
// }

// func (c *ArticleController) handleNotMatchUser(w http.ResponseWriter, req *http.Request) {
// 	err := apperrors.NotMatchUser.Wrap(errors.New("not match user"), "invalid parameter")
// 	apperrors.ErrorHandler(w, req, err)
// }
