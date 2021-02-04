package main

import (
	"modules/database"
	"modules/models"
	"modules/routes"
)

func main() {
	db := database.ConnectDB()
	db.AutoMigrate(&models.User{}, &models.ShortUrl{})

	r := routes.SetupRoute(db)

	api := routes.CreateApiUrl(r)

	api.Run(":8080")
}
