package controllers

import (
	"modules/auth"
	"modules/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
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

	if result := db.Create(&user); result.Error != nil {
		c.JSON(422, result.Error)
		return
	} else {
		c.JSON(http.StatusOK, result.Value)
	}
}

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	var req UserLogin

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	checkPasswordHash := CheckPasswordHash(req.Password, user.Password)

	//compare the user from the request, with the one we defined:
	if user.Email != req.Email || !checkPasswordHash {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	ts, err := auth.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	saveErr := auth.CreateAuth(user.ID, ts)

	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}

	tokens := map[string]string{
		"token":         ts.AccessToken,
		"refresh_token": ts.RefreshToken,
		// "type":          "Bearer",
	}
	c.JSON(http.StatusOK, tokens)
}

func Logout(c *gin.Context) {
	au, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	deleted, delErr := auth.DeleteAuth(au.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}
