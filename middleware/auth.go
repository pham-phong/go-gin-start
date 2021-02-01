package middleware

import (
	"github.com/gin-gonic/gin"
	"modules/auth"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenAuth, err := auth.ExtractTokenMetadata(c.Request)
		if err != nil {
			respondWithError(c, 401, "Unauthorized")
			return
		}
		_, err = auth.FetchAuth(tokenAuth)
		if err != nil {
			respondWithError(c, 401, "Unauthorized")
			return
		}
		c.Next()
	}
}
