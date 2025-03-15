package database

import "database/sql"

type OrderRepository struct {
	Db *sql.DB
}

type Order struct {
	Id      int
	Product string
	Amount  int
}

func (r *OrderRepository) CreateTable() error {
	_, err := r.Db.Exec(`CREATE TABLE IF NOT EXISTS orders (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
		product TEXT,
		amount INTEGER
		)`)

	return err

}

func (r *OrderRepository) Insert(order Order) error {
	_, err := r.Db.Exec("INSERT INTO orders (product, amount) VALUES (?, ?)", order.Product, order.Amount)

	return err
}
