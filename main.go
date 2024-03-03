package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sambacarlson/todo_backend_golang/src/logic"
	"github.com/sambacarlson/todo_backend_golang/src/persistence"
	"github.com/sambacarlson/todo_backend_golang/src/presentation"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load()

	parsedPort, err := strconv.Atoi (os.Getenv("DB_PORT"));
	if err != nil {
		log.Fatal(err)
		parsedPort = 5432
	}


	CONF := struct {
		BASE_URL string
		DB_URL   string
		DB_PORT int
		DB_USER string
		DB_HOST string
		DB_NAME string
		DB_PASSWORD string
		DB_TLSMODE string
	}{
		BASE_URL: fmt.Sprintf("localhost:%s", os.Getenv("API_PORT")),
		DB_URL:   os.Getenv("DB_URI"),
		DB_PORT: parsedPort,
    DB_USER: os.Getenv("DB_USER"),
    DB_HOST: os.Getenv("DB_HOST"),
    DB_NAME: os.Getenv("DB_NAME"),
    DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_TLSMODE: "sslmode=disable",
	}
	if err != nil {
		log.Fatal("error loading env file")
		return
	}
	// db, err := sql.Open("postgres", "user=todosuperuser password=mhXYtq3xBvw5Cmw dbname=mytodo_db sslmode=disable")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	dbString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s %s", CONF.DB_HOST, CONF.DB_PORT, CONF.DB_USER, CONF.DB_PASSWORD, CONF.DB_NAME, CONF.DB_TLSMODE)
	sqlDBInstance, err := sql.Open("postgres", dbString)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDBInstance,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	defer sqlDBInstance.Close()

	//New database instance
	pers, _ := persistence.NewConnection(db)
	logics, _ := logic.NewLogic(pers)

	// register routes
	_ = presentation.Route(CONF.BASE_URL, *logics)
}
