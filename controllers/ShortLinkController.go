package controllers

import (
	"math/rand"
	"modules/database"
	"modules/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UrlCreatRequest struct {
	Link string `json:"link"`
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

func randString(n int, c *gin.Context) string {
	db := c.MustGet("db").(*gorm.DB)

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	// check in database
	var short_link models.ShortUrl
	if err := db.Where("code = ?", b).First(&short_link).Error; err == nil {
		randString(n, c)
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
	rand_url := randString(10, c)
	// Create user
	short_link := models.ShortUrl{Code: rand_url, Link: input.Link}

	db := c.MustGet("db").(*gorm.DB)
	result := db.Create(&short_link)

	c.JSON(http.StatusOK, result.Value)
}

func GetShortlinks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	links := []models.ShortUrl{}

	if err := db.Find(&links).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, links)
}

func Pagination(c *gin.Context, db *gorm.DB) *gorm.DB {
	links := []models.ShortUrl{}

	db.Scopes(database.Paginate(c)).Find(&links)

	return db.Scopes(database.Paginate(c)).Find(&links)
}

func HandleShortUrlRedirect(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var result models.ShortUrl

	code := c.Request.URL.Path[len("/api/"):]

	if err := db.Where("code = ?", code).First(&result).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var click int = int(result.Click) + 1
	db.Model(&result).Update("click", click)

	c.JSON(http.StatusOK, result.Link)
	// c.Redirect(302, result.Link)
}
