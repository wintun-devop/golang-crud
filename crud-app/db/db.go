package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectPostgresDB() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	// connstring := "user=postgres dbname=postgres password='*****' host=localhost port=5432 sslmode=disable"
	dbString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=5432 sslmode=disable", dbUser, dbName, dbPassword, dbHost)
	connstring := dbString
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("db", db)
	return db
}
