package controllers

import (
	"siswa-api/config"
	"siswa-api/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateStudent(c *gin.Context) {
	var student models.Student
	student.Nama = c.PostForm("nama")
	student.NIS = c.PostForm("nis")

	file, _ := c.FormFile("foto")
	if file != nil {
		path := fmt.Sprintf("uploads/%s", file.Filename)
		c.SaveUploadedFile(file, path)
		student.Foto = path
	}

	config.DB.Create(&student)
	c.JSON(200, student)
}

func GetStudents(c *gin.Context) {
	var students []models.Student
    nama := c.Query("nama")

    query := config.DB

    if nama != "" {
        query = query.Where("nama LIKE ?", "%"+nama+"%")
    }

    if err := query.Find(&students).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, students)
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	config.DB.First(&student, id)

	c.ShouldBind(&student)
	student.Nama = c.PostForm("nama")
	student.NIS = c.PostForm("nis")

	file, _ := c.FormFile("foto")
	if file != nil {
		path := fmt.Sprintf("uploads/%s", file.Filename)
		c.SaveUploadedFile(file, path)
		student.Foto = path
	}

	config.DB.Save(&student)
	c.JSON(200, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Student{}, id)
	c.JSON(200, gin.H{"message": "Deleted"})
}
