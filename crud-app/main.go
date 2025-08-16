package main

import (
	"crud-app/db"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	appSecretKey := os.Getenv("APP_SECRET_KEY")
	fmt.Println("app secret key", appSecretKey)
	db.ConnectPostgresDB()

}
