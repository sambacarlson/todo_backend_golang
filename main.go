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
	CONF := struct {
		BASE_URL string
		DB_URL   string
	}{
		BASE_URL: fmt.Sprintf("localhost:%s", os.Getenv("API_PORT")),
		DB_URL:   os.Getenv("DB_URI"),
	}
	if err != nil {
		log.Fatal("error loading env file")
		return
	}
	db, err := sql.Open("postgres", CONF.DB_URL)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Printf("db here;; %v", db)

	// register routes
	_ = presentation.Route(CONF.BASE_URL)
}
