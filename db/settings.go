package db

import (
	"fmt"
	"os"
)

type DBSettings struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBConn     string
}

func (dbSettings DBSettings) GetDBConn() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbSettings.DBHost,
		dbSettings.DBUser,
		dbSettings.DBPassword,
		dbSettings.DBName,
	)
}

func GetDBSettings() DBSettings {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbDb := os.Getenv("POSTGRES_DB")

	return DBSettings{
		DBHost:     dbHost,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbDb,
	}
}
