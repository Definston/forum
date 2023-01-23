package app

import (
	"database/sql"
	"log"
	"os"

	"forum/internal/controller"
	"forum/internal/repository"
	"forum/internal/service"

	_ "github.com/mattn/go-sqlite3"
)

func Run() {
	db := database()

	repository := repository.NewRepository(db)

	service := service.NewService(repository)

	serv := controller.NewServer(service)

	log.Println("http://localhost:8080")
	log.Fatalln(serv.Run(":8080"))
}

func database() *sql.DB {
	db, err := sql.Open("sqlite3", "./forum.sqlite")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	stTable, err := os.ReadFile("./configDB.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(stTable))
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
