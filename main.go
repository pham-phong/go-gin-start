package main

import (
	"modules/database"
	"modules/models"
	"modules/routes"
)

func main() {
	db := database.ConnectDB()
	db.AutoMigrate(&models.User{})

	api.CreateApiUrl(db)
	// Listen and server on 0.0.0.0:8080
	api.Router.Run(":8080")
}
