package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AppDeveloperMLLB/todo_app/apperrors"
	"github.com/AppDeveloperMLLB/todo_app/controllers/services"
)

// TodoController - Todoに関するコントローラ
type TodoController struct {
	service services.TodoService
}

// NewTodoController - TodoControllerのコンストラクタ
func NewTodoController(s services.TodoService) *TodoController {
	return &TodoController{service: s}
}

// PostArticleHandler - POST /articleのハンドラ
// func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
// 	var reqArticle models.Article
// 	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
// 		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
// 		apperrors.ErrorHandler(w, req, err)
// 		return
// 	}

// 	authedUserID := common.GetUserID(req.Context())
// 	userName := common.GetUserName(req.Context())
// 	reqArticle.UserID = authedUserID
// 	reqArticle.UserName = userName

// 	article, err := c.service.PostArticleService(reqArticle)
// 	if err != nil {
// 		apperrors.ErrorHandler(w, req, err)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(article)
// }

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
