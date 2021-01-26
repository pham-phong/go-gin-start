package api

import (
	Usercontrollers "modules/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateApiUrl() {
	Router = gin.Default()
	// auth of the API
	r := Router.Group("/api")
	{
		r.GET("/user", Usercontrollers.GetUser)
	}
}
