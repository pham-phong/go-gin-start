package mappings

import (
	"modules/controllers"

	"github.com/gin-gonic/gin"
)
var Router *gin.Engine
func CreateUrlMappings() {
	Router = gin.Default()
	Router.Use(controllers.Cors())
	// auth of the API
	auth := Router.Group("/auth")
	{
		auth.GET("/users/:id", controllers.GetUserDetail)
		auth.GET("/users/", controllers.GetUser)
		auth.POST("/login/", controllers.Login)
		auth.PUT("/users/:id", controllers.UpdateUser)
		auth.POST("/users", controllers.PostUser)
	}
}