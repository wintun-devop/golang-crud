package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	appSecretKey := os.Getenv("APP_SECRET_KEY")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	fmt.Println("Welcome to CRUD App.")
	fmt.Println("app secret key", appSecretKey)
	fmt.Println("dbhost", dbHost, ",dbName", dbName, "dbUser", dbUser, "dbPassword", dbPassword)
	// connStr := "user=dbUser password=yourpass dbname=yourdb sslmode=disable"
	// db.Init(connStr)
}
