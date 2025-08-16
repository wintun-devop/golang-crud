package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	appSecretKey := os.Getenv("APP_SECRET_KEY")
	fmt.Println("Welcome to CRUD App.")
	fmt.Println("app secret key", appSecretKey)

}
