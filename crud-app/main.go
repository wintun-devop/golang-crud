package main

import (
	"crud-app/db"
	"crud-app/services"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	/*
		appSecretKey := os.Getenv("APP_SECRET_KEY")
		fmt.Println("app secret key", appSecretKey)
	*/
	db.ConnectPostgresDB()
	// Create
	/*
		newProduct := services.Product{ID: utils.GenerateCUID(), Name: "Phone", Model: "Oppo A9", Price: 300.99}
		inserted, err := services.InsertProduct(newProduct)
		if err != nil {
			fmt.Println("Create error:", err)
		}
		fmt.Println(inserted)
	*/
	// Read
	/*
		p, err := servcies.GetProduct("cuidffsafufdfdf")
		if err != nil {
			fmt.Println("Read error:", err)
		} else if p != nil {
			productMap := map[string]interface{}{
				"id":         p.ID,
				"name":       p.Name,
				"model":      p.Model,
				"price":      p.Price,
				"created_at": p.CreatedAt,
				"updated_at": p.UpdatedAt,
			}
			prettyJSON, _ := utils.MapToPrettyJSON(productMap)
			// fmt.Printf(p.ID, p.Model)
			// fmt.Printf("Product: %+v\n", *p)
			fmt.Println(prettyJSON)
		}
	*/
	// Delete
	err := services.DeleteProduct("cmeehh4hd00002w9i4wtk3kif")
	if err != nil {
		fmt.Println("Delete error:", err)
	} else {
		fmt.Println("Delete success")
	}

}
