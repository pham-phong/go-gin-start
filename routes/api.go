package routes

import (
	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
	"modules/controllers"
	"modules/middleware"
	// cors "github.com/rs/cors/wrapper/gin"
)

// var Router *gin.Engine

// func CreateApiUrl(db *gorm.DB) *gin.Engine {
// 	Router := gin.New()

// 	Router.Use(cors.AllowAll())
// 	// Provide db variable to controllers
// 	Router.Use(func(c *gin.Context) {
// 		c.Set("db", db)
// 		c.Next()
// 	})
// 	// Router of the API
// 	r := Router.Group("/api")
// 	{
// 		r.POST("/login/", controllers.Login)

// 		r.Use(middleware.AuthorizeJWT())
// 		{
// 			r.GET("/users", controllers.GetUser)
// 			r.POST("/user", controllers.SaveUser)
// 			r.GET("/user/:id", controllers.FindUser)
// 			r.PUT("/user/:id", controllers.UpdateUser)
// 			r.DELETE("/logout", controllers.Logout)
// 		}

// 	}
// 	return Router
// }

func CreateApiUrl(r *gin.Engine) *gin.Engine {
	r.POST("/api/login", controllers.Login)
	r.POST("/api/register", controllers.Register)

	r.GET("api/get-shortlink", controllers.GetShortlinks)
	r.POST("/create-shortlink", controllers.CreateShortLink)
	r.NoRoute(controllers.HandleShortUrlRedirect)

	auth := r.Group("/api/auth")
	auth.Use(middleware.AuthorizeJWT())
	{
		auth.GET("/user", controllers.ShowUser)
		auth.GET("/users", controllers.GetUser)
		auth.POST("/user", controllers.SaveUser)
		auth.GET("/user/:id", controllers.FindUser)
		auth.PUT("/user/:id", controllers.UpdateUser)
		auth.DELETE("/logout", controllers.Logout)

		// auth.POST("/create-shortlink", controllers.CreateShortLink)
	}
	return r
}
