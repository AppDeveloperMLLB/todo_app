package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	dbSettings := GetDBSettings()
	dbConn := dbSettings.GetDBConn()
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
