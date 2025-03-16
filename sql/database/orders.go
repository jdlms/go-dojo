package database

import (
	"context"
	"database/sql"
)

type OrderRepository struct {
	Db *sql.DB
}

type Order struct {
	Id      int
	Product string
	Amount  int
}

func (r *OrderRepository) CreateTable(ctx context.Context) error {
	_, err := r.Db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS orders (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
		product TEXT,
		amount INTEGER
		)`)

	return err
}

func (r *OrderRepository) Insert(ctx context.Context, order Order) error {
	_, err := r.Db.ExecContext(ctx, "INSERT INTO orders (product, amount) VALUES (?, ?)", order.Product, order.Amount)

	return err
}

func (r *OrderRepository) GetAll(ctx context.Context) ([]Order, error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.Id, &order.Product, &order.Amount)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepository) GetById(ctx context.Context, id int) (Order, error) {
	var order Order
	err := r.Db.QueryRowContext(ctx, "SELECT * FROM orders WHERE id = ?", id).Scan(&order.Id, &order.Product, &order.Amount)
	if err != nil {
		return Order{}, err
	}
	return order, nil
}

func (r *OrderRepository) Update(ctx context.Context, order Order) error {
	_, err := r.Db.ExecContext(ctx, "UPDATE orders SET product = ?, amount = ? WHERE id = ?", order.Product, order.Amount, order.Id)
	return err
}

func (r *OrderRepository) Delete(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, "DELETE FROM orders WHERE id = ?", id)
	return err
}
