package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"solid/order"
)

type SQLiteRepo struct {
	db *sql.DB
}

func NewSQLiteRepo(db *sql.DB) *SQLiteRepo {
	return &SQLiteRepo{db: db}
}

func (r *SQLiteRepo) Init() error {
	query := `
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		customer TEXT NOT NULL,
		products TEXT NOT NULL,
		total REAL NOT NULL,
		status TEXT NOT NULL
	)`
	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRepo) SaveOrder(order *order.Order) error {
	productsJSON, err := json.Marshal(order.Products)
	if err != nil {
		return fmt.Errorf("failed to marshal products: %w", err)
	}

	result, err := r.db.Exec(
		"INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
		order.Customer,
		string(productsJSON),
		order.Total,
		order.Status,
	)
	if err != nil {
		return fmt.Errorf("failed to save order: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}
	order.ID = int(id)

	return nil
}
