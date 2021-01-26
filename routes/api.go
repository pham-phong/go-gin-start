package api

import (
	"modules/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
		r.GET("/user", Usercontrollers.GetUser)
		r.POST("/user", Usercontrollers.SaveUser)
		r.GET("/user/:id", Usercontrollers.FindUser)
		r.PUT("/user/:id", Usercontrollers.UpdateUser)

		// r.POST("/login/", Usercontrollers.login)
		
	}
}
