package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sambacarlson/todo_backend_golang/src/presentation"
)

func main() {
	err := godotenv.Load()
	Const := struct {
		BASE_URL string
		DB_URL   string
	}{
		BASE_URL: fmt.Sprintf("localhost:%s", os.Getenv("API_PORT")),
		DB_URL:   os.Getenv("DB_URI"),
	}
	if err != nil {
		log.Fatal("Error loading env file")
		return
	}
	db, err := sql.Open("postgres", Const.DB_URL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_ = presentation.Route(Const.BASE_URL)
}
