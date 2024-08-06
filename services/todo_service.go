package services

import (
	"github.com/AppDeveloperMLLB/todo_app/apperrors"
	"github.com/AppDeveloperMLLB/todo_app/models"
	"github.com/AppDeveloperMLLB/todo_app/repositories"
)

// func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
// 	var article models.Article
// 	var commentList []models.Comment
// 	var articleGetErr, commentGetErr error

// 	type articleResult struct {
// 		article models.Article
// 		err     error
// 	}
// 	articleChan := make(chan articleResult)
// 	defer close(articleChan)

// 	go func(ch chan<- articleResult, db *sql.DB, articleId int) {
// 		article, err := repositories.SelectArticleDetail(db, articleId)
// 		ch <- articleResult{article: article, err: err}
// 	}(articleChan, s.db, articleID)

// 	type commentResult struct {
// 		commentList []models.Comment
// 		err         error
// 	}
// 	commentChan := make(chan commentResult)
// 	defer close(commentChan)

// 	go func(ch chan<- commentResult, db *sql.DB, articleId int) {
// 		commentList, err := repositories.SelectCommentList(db, articleId)
// 		ch <- commentResult{commentList: commentList, err: err}
// 	}(commentChan, s.db, articleID)

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case ar := <-articleChan:
// 			article, articleGetErr = ar.article, ar.err
// 		case cr := <-commentChan:
// 			commentList, commentGetErr = cr.commentList, cr.err
// 		}
// 	}

// 	if articleGetErr != nil {
// 		if errors.Is(articleGetErr, sql.ErrNoRows) {
// 			err := apperrors.NAData.Wrap(articleGetErr, "no data")
// 			return models.Article{}, err
// 		}

// 		err := apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
// 		return models.Article{}, err
// 	}

// 	if commentGetErr != nil {
// 		err := apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
// 		return models.Article{}, err
// 	}

// 	article.CommentList = append(article.CommentList, commentList...)
// 	return article, nil
// }

// func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
// 	var article models.Article
// 	var commentList []models.Comment
// 	var articleGetErr, commentGetErr error

// 	// ロックを行うようの変数
// 	var amu sync.Mutex
// 	var cmu sync.Mutex

// 	// WaitGroupを使って、追加したカウンタが0になるまでwaitすることで処理を待てる
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go func(db *sql.DB, articleId int) {
// 		defer wg.Done()
// 		amu.Lock()
// 		article, articleGetErr = repositories.SelectArticleDetail(db, articleId)
// 		amu.Unlock()
// 	}(s.db, articleID)

// 	go func(db *sql.DB, articleId int) {
// 		defer wg.Done()

// 		cmu.Lock()
// 		commentList, commentGetErr = repositories.SelectCommentList(db, articleId)
// 		cmu.Unlock()
// 	}(s.db, articleID)

// 	wg.Wait()

// 	if articleGetErr != nil {
// 		if errors.Is(articleGetErr, sql.ErrNoRows) {
// 			err := apperrors.NAData.Wrap(articleGetErr, "no data")
// 			return models.Article{}, err
// 		}

// 		err := apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
// 		return models.Article{}, err
// 	}

// 	if commentGetErr != nil {
// 		err := apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
// 		return models.Article{}, err
// 	}

// 	article.CommentList = append(article.CommentList, commentList...)
// 	return article, nil
// }

// func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
// 	id, err := repositories.InsertArticle(s.db, article)
// 	if err != nil {
// 		err = apperrors.InsertDataField.Wrap(err, "fail to record data")
// 		return models.Article{}, err
// 	}

// 	article.ID = id
// 	return article, nil
// }

// GetTodoListService
func (s *MyAppService) GetTodoListService(page int, perPage int) ([]models.Todo, error) {
	list, err := repositories.SelectTodoList(s.db, page, perPage, "all")
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(list) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return list, nil
}

// func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
// 	err := repositories.UpdateNiceNum(s.db, article.ID)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			err = apperrors.NAData.Wrap(err, "not exist target article")
// 			return models.Article{}, err
// 		}

// 		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")
// 		return models.Article{}, err
// 	}

// 	return article, nil
// }
