package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AppDeveloperMLLB/todo_app/api"
	"github.com/AppDeveloperMLLB/todo_app/db"
	"github.com/AppDeveloperMLLB/todo_app/settings"
	_ "github.com/lib/pq"
)

func main() {
	settings.Initialize()

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	r := api.NewRouter(db)

	fmt.Println("Server is running on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
