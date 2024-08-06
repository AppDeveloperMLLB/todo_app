package repositories

import (
	"database/sql"

	"github.com/AppDeveloperMLLB/todo_app/models"
)

// SelectTodoList - Todo一覧を取得する
func SelectTodoList(db *sql.DB, page int, perPage int, status string) ([]models.Todo, error) {
	const sqlStr = `
		SELECT * FROM todos LIMIT $1 OFFSET $2;
	`

	rows, err := db.Query(sqlStr, perPage, page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todoArray := make([]models.Todo, 0)
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(
			&todo.ID,
			&todo.UserID,
			&todo.Title,
			&todo.Description,
			&todo.Status,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		todoArray = append(todoArray, todo)
	}

	return todoArray, nil
}

// InsertArticle - 記事を登録する
// func InsertArticle(db *sql.DB, article models.Article) (int, error) {
// 	const sqlStr = `
// 		INSERT INTO articles (user_id, title, contents, username, nice, created_at)
// 		VALUES ($1, $2, $3, $4, 0, now()) returning id;
// 	`

// 	var newArticleID int
// 	err := db.QueryRow(sqlStr, article.UserID, article.Title, article.Contents, article.UserName).Scan(&newArticleID)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return newArticleID, nil

// 	// Postgresqlだと以下のコードは使えない
// 	// const sqlStr = `
// 	// 	INSERT INTO articles (title, contents, username, nice, created_at)
// 	// 	VALUES ($1, $2, $3, 0, now());
// 	// `

// 	// result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName);
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }

// 	// fmt.Println(result.LastInsertId())
// 	// fmt.Println(result.RowsAffected())
// }

// // SelectArticleList - 記事一覧を取得する
// func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
// 	const sqlStr = `
// 		SELECT * FROM articles LIMIT 10 OFFSET $1;
// 	`

// 	// 取得できるデータが0か1の場合、QueryRowを使う
// 	// Queryとの違いは、Queryはデータが0軒の時はエラーにならないが、
// 	// QueryRowの場合、Errメソッドからエラーを取得できる
// 	// row := db.QueryRow(sqlStr, page-1)
// 	// if err := row.Err(); err != nil {
// 	// 	fmt.Println(err)
// 	// 	return nil, err
// 	// }

// 	// var article models.Article
// 	// err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdAt)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return nil, err
// 	// }

// 	// var createdAt sql.NullTime
// 	// if createdAt.Valid {
// 	// 	article.CreatedAt = createdAt.Time
// 	// }

// 	rows, err := db.Query(sqlStr, page-1)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	articleArray := make([]models.Article, 0)
// 	for rows.Next() {
// 		var article models.Article
// 		var createdAt sql.NullTime
// 		err := rows.Scan(&article.ID, &article.UserID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdAt)
// 		if createdAt.Valid {
// 			article.CreatedAt = createdAt.Time
// 		}
// 		if err != nil {
// 			return nil, err
// 		}

// 		articleArray = append(articleArray, article)
// 	}

// 	return articleArray, nil
// }

// // SelectArticleDetail - 記事詳細を取得する
// func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
// 	const sqlStr = `
// 		SELECT *
// 		FROM articles
// 		WHERE id = $1;
// 	`

// 	row := db.QueryRow(sqlStr, articleID)
// 	if err := row.Err(); err != nil {
// 		return models.Article{}, err
// 	}

// 	var article models.Article
// 	var createdAt sql.NullTime

// 	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdAt)
// 	if err != nil {
// 		return models.Article{}, err
// 	}

// 	if createdAt.Valid {
// 		article.CreatedAt = createdAt.Time
// 	}

// 	return article, nil
// }

// // UpdateNiceNum - いいね数を更新する
// func UpdateNiceNum(db *sql.DB, articleID int) error {
// 	const sqlGetNice = `
// 		SELECT nice
// 		FROM articles
// 		WHERE id = $1;
// 	`

// 	const sqlUpdateNice = `
// 		UPDATE articles
// 		SET nice = $1
// 		WHERE id = $2;
// 	`

// 	row := db.QueryRow(sqlGetNice)
// 	if err := row.Err(); err != nil {
// 		return err
// 	}

// 	var nice int
// 	err := row.Scan(&nice)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = db.Exec(sqlUpdateNice, articleID, nice+1)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
