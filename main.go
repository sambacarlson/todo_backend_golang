package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sambacarlson/todo_backend_golang/src/logic"
	"github.com/sambacarlson/todo_backend_golang/src/persistence"
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
	db, err := sql.Open("postgres", "user=todosuperuser password=mhXYtq3xBvw5Cmw dbname=mytodo_db sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//New database instance
	pers, _ := persistence.NewConnection(db)
	logics, _ := logic.NewLogic(pers)

	// register routes
	_ = presentation.Route(CONF.BASE_URL, *logics)
}
