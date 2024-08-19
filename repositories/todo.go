package repositories

import (
	"database/sql"

	"github.com/AppDeveloperMLLB/todo_app/models"
)

// SelectTodo - Todoを取得する
func SelectTodo(db *sql.DB, todoID int) (models.Todo, error) {
	const sqlStr = `
		SELECT * FROM todos WHERE id = $1;
	`

	row := db.QueryRow(sqlStr, todoID)
	if err := row.Err(); err != nil {
		return models.Todo{}, err
	}

	var todo models.Todo
	err := row.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

// SelectTodoList - Todo一覧を取得する
func SelectTodoList(db *sql.DB, page int, perPage int, status string) ([]models.Todo, error) {
	const sqlStr = `
		SELECT * FROM todos LIMIT $1 OFFSET $2;
	`

	offset := (page - 1) * perPage
	if page == 0 {
		offset = 0
	}

	rows, err := db.Query(sqlStr, perPage, offset)
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

// CreateTodo - Todoを登録する
func CreateTodo(db *sql.DB, todo models.Todo) (models.Todo, error) {
	const sqlStr = `
		INSERT INTO todos (user_id, title, description, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, now(), now()) returning id, created_at, updated_at;
	`
	var newTodoID int
	var createdAt, updatedAt sql.NullTime
	err := db.QueryRow(sqlStr, todo.UserID, todo.Title, todo.Description, todo.Status).Scan(&newTodoID, &createdAt, &updatedAt)
	if err != nil {
		return models.Todo{}, err
	}

	todo.ID = newTodoID
	todo.CreatedAt = createdAt.Time
	todo.UpdatedAt = updatedAt.Time
	return todo, nil
}

// UpdateTodo - Todoを更新する
func UpdateTodo(db *sql.DB, todo models.Todo) (models.Todo, error) {
	const sqlStr = `
		UPDATE todos
		SET title = $1, description = $2, status = $3, updated_at = now()
		WHERE id = $4 returning updated_at;
	`

	var updatedAt sql.NullTime
	err := db.QueryRow(sqlStr, todo.Title, todo.Description, todo.Status, todo.ID).Scan(&updatedAt)
	if err != nil {
		return models.Todo{}, err
	}

	todo.UpdatedAt = updatedAt.Time
	return todo, nil
}

// DeleteTodo - Todoを削除する
func DeleteTodo(db *sql.DB, todoID int) error {
	const sqlStr = `
		DELETE FROM todos WHERE id = $1;
	`

	_, err := db.Exec(sqlStr, todoID)
	if err != nil {
		return err
	}

	return nil
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
