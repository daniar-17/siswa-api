package controllers

import (
	"siswa-api/config"
	"siswa-api/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input models.User
	c.ShouldBindJSON(&input)

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(hashed)

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(400, gin.H{"error": "Username already exists"})
		return
	}

	c.JSON(200, gin.H{"message": "Register success"})
}

func Login(c *gin.Context) {
	var input models.User
	c.ShouldBindJSON(&input)

	var user models.User
	config.DB.Where("username = ?", input.Username).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "Wrong credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	c.JSON(200, gin.H{"token": tokenString})
}
