package controllers

import (
	"modules/auth"
	"modules/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ShowUser(c *gin.Context) {
	// c.Request.Header.Get()
	tokenAuth, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var user models.User
	if err := db.Where("id = ?", tokenAuth.UserId).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	users := []models.User{}

	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func SaveUser(c *gin.Context) {
	var req models.User
	// Validate input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//hash password
	hashedPassword, _ := HashPassword(req.Password)
	req.Password = string(hashedPassword)

	//Create user
	user := models.User{Username: req.Username, Email: req.Email, Password: req.Password}

	db := c.MustGet("db").(*gorm.DB)
	result := db.Create(&user)

	c.JSON(http.StatusOK, result)
}

/**
* GET /user/:id
* Find a user
 */
func FindUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// PATCH /user/:id
// Update a user
func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Model(&user).Updates(input)

	c.JSON(http.StatusOK, result)
}
