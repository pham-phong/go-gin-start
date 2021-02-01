package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"modules/controllers"
)

var Router *gin.Engine

func CreateApiUrl(db *gorm.DB) {
	Router = gin.Default()

	// Provide db variable to controllers
	Router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	// Router of the API
	r := Router.Group("/api")
	{
		r.GET("/user", controllers.GetUser)
		r.POST("/user", controllers.SaveUser)
		r.GET("/user/:id", controllers.FindUser)
		r.PUT("/user/:id", controllers.UpdateUser)

		r.POST("/login/", controllers.Login)

	}
}
