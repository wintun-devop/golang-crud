package services

import (
	"crud-app/db"
	"database/sql"

	_ "github.com/lib/pq"
)

type Product struct {
	ID        string
	Name      string
	Model     string
	Price     float64
	CreatedAt string
	UpdatedAt string
}

/*
func InsertProduct(p Product) error {
	_, err := db.ConnectPostgresDB().Exec(
		"INSERT INTO products (id,name, model, price) VALUES ($1, $2, $3, $4)",
		p.ID, p.Name, p.Model, p.Price,
	)
	return err
}
*/

func InsertProduct(p Product) (Product, error) {
	var inserted Product
	err := db.ConnectPostgresDB().QueryRow(
		`INSERT INTO products (id, name, model, price)
         VALUES ($1, $2, $3, $4)
         RETURNING id, name, model, price`,
		p.ID, p.Name, p.Model, p.Price,
	).Scan(&inserted.ID, &inserted.Name, &inserted.Model, &inserted.Price)

	return inserted, err
}

// Read
func GetProduct(id string) (*Product, error) {
	row := db.ConnectPostgresDB().QueryRow(
		"SELECT id, name, model, price, created_at, updated_at FROM products WHERE id = $1", id,
	)
	var p Product
	err := row.Scan(&p.ID, &p.Name, &p.Model, &p.Price, &p.CreatedAt, &p.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &p, err
}

// Update
func UpdateProduct(p Product) error {
	_, err := db.ConnectPostgresDB().Exec(
		"UPDATE products SET name = $1, model = $2, price = $3 WHERE id = $4",
		p.Name, p.Model, p.Price, p.ID,
	)
	return err
}

// Delete
func DeleteProduct(id string) error {
	_, err := db.ConnectPostgresDB().Exec("DELETE FROM products WHERE id = $1", id)
	return err
}
