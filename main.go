package main

import (
	"modules/database"
	"modules/models"
	"modules/routes"
)

func main() {
	db := database.ConnectDB()
	db.AutoMigrate(&models.User{})

	// r := routes.CreateApiUrl(db)
	// // Listen and server on 0.0.0.0:8080
	// r.Run(":8080")

	r := routes.SetupRoute(db)

	api := routes.CreateApiUrl(r)

	api.Run(":8080")
}
