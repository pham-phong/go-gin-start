package controllers

import (
	"math/rand"
	"modules/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UrlCreatRequest struct {
	Code string `json:"code"`
	Link string `json:"link"`
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

func randURL(n int, c *gin.Context) string {
	db := c.MustGet("db").(*gorm.DB)

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	// check randURL in database
	var short_link models.ShortUrl
	if err := db.Where("code = ?", b).First(&short_link).Error; err == nil {
		randURL(n, c)
	}
	return string(b)
}

func CreateShortLink(c *gin.Context) {
	// validate input
	var input UrlCreatRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// random URL
	CodeRand := randURL(10, c)
	// Create user
	short_link := models.ShortUrl{Code: CodeRand, Link: input.Link}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&short_link)
	c.JSON(http.StatusOK, gin.H{"data": short_link})
}

func HandleShortUrlRedirect(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var short_link models.ShortUrl

	code := c.Request.URL.Path[len("/"):]

	if err := db.Where("code = ?", code).First(&short_link).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.Redirect(302, short_link.Link)
}
