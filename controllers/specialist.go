package controllers

import (
	"github.com/gin-gonic/gin"
)

func CreateSpecialist(c *gin.Context) {
	c.JSON(200, "Specialist created successfully")
}
