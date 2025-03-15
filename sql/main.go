package main

import (
	"database/sql"
	"log"

	"github.com/jdlms/go-dojo/sql/database"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./shop.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	orderRepository := &database.OrderRepository{Db: db}

	err = orderRepository.CreateTable()
	if err != nil {
		log.Fatal("Error creating orders table", err)
	}
}
