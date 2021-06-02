package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	cors "github.com/rs/cors/wrapper/gin"
)

func SetupRoute(db *gorm.DB) *gin.Engine {
	// init
	r := gin.New()

	r.Use(cors.AllowAll())
	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	return r
}
