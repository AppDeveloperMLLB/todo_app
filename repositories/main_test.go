package repositories_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"

	"github.com/AppDeveloperMLLB/todo_app/db"
	"github.com/AppDeveloperMLLB/todo_app/settings"
	_ "github.com/lib/pq"
)

var testDB *sql.DB

func setup() error {
	settings.Initialize()
	var err error
	testDB, err = db.ConnectDB()
	if err != nil {
		log.Println("Failed to connect to the database")
		return err
	}

	err = cleanupDB()
	if err != nil {
		log.Println("Failed to cleanup the database")
		return err
	}

	err = setupTestData()
	if err != nil {
		log.Println("Failed to setup test data")
		return err
	}

	return nil
}

func teardown() {
	cleanupDB()
	testDB.Close()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	m.Run()
	teardown()
}

func setupTestData() error {
	cmd := exec.Command(
		"psql",
		"-h",
		"127.0.0.1",
		"-U",
		"test",
		"-d",
		os.Getenv("POSTGRES_DB"),
		"-f",
		"./testdata/createTable.sql")
	cmd.Env = append(os.Environ(), "PGPASSWORD=password")
	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command(
		"psql",
		"-h",
		"127.0.0.1",
		"-U",
		"test",
		"-d",
		os.Getenv("POSTGRES_DB"),
		"-f",
		"./testdata/insertTestData.sql")
	cmd.Env = append(os.Environ(), "PGPASSWORD=password")
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
func cleanupDB() error {
	cmd := exec.Command(
		"psql",
		"-h",
		"127.0.0.1",
		"-U",
		"test",
		"-d",
		os.Getenv("POSTGRES_DB"),
		"-f",
		"./testdata/cleanup.sql")
	cmd.Env = append(os.Environ(), "PGPASSWORD=password")
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
