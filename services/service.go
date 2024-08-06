package services

import "database/sql"

// MyAppService - サービスの構造体
type MyAppService struct {
	db *sql.DB
}

// NewMyAppService - MyAppServiceのコンストラクタ
func NewMyAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}
