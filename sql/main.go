package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/jdlms/go-dojo/sql/database"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./shop.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	orderRepository := &database.OrderRepository{Db: db}

	err = orderRepository.CreateTable(ctx)
	if err != nil {
		log.Fatal("Error creating orders table", err)
	}

	err = orderRepository.Insert(ctx, database.Order{Product: "Laptop", Amount: 10})
	if err != nil {
		log.Fatal("Error inserting order:", err)
	}

	err = orderRepository.Insert(ctx, database.Order{Product: "Keyboard", Amount: 50})
	if err != nil {
		log.Fatal("Error inserting order:", err)
	}

	orders, err := orderRepository.GetAll(ctx)
	if err != nil {
		log.Fatal("Error getting orders:", err)
	}

	log.Println(orders)

	order, err := orderRepository.GetById(ctx, orders[0].Id)
	if err != nil {
		log.Fatal("Error getting order:", err)
	}

	order.Amount = 1500
	err = orderRepository.Update(ctx, order)
	if err != nil {
		log.Fatal("Error updating order:", err)
	}

	orders, err = orderRepository.GetAll(ctx)
	if err != nil {
		log.Fatal("Error getting orders:", err)
	}

	log.Println(orders)

	err = orderRepository.Delete(ctx, order.Id)
	if err != nil {
		log.Fatal("Error deleting order:", err)
	}

	orders, err = orderRepository.GetAll(ctx)
	if err != nil {
		log.Fatal("Error getting orders:", err)
	}

	log.Println(orders)

}

// from https://codingwithpatrik.dev/posts/go-golang-sql/
