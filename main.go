package main

import (
	"log"
	"modules/database"
	"modules/models"
	"modules/routes"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db := database.ConnectDB()
	db.AutoMigrate(&models.User{}, &models.ShortUrl{})

	r := routes.SetupRoute(db)

	api := routes.CreateApiUrl(r)

	api.Run(":8080")
}
