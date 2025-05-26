package controllers

import (
	"github.com/gin-gonic/gin"
)

func CreateSpecialist(c *gin.Context) {
	c.JSON(200, "Specialist created successfully")
}

func GetSpecialist(c *gin.Context) {
	c.JSON(200, "Specialist get 2 successfully")
}
