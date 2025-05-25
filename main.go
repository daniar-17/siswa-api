package main

import (
	"siswa-api/config"
	"siswa-api/models"
	"siswa-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// CORS middleware aktif di instance ini
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Inisialisasi database
	config.InitDB()
	config.DB.AutoMigrate(&models.User{}, &models.Student{})

	// Static file (untuk foto)
	router.Static("/uploads", "./uploads")

	// Routing
	routes.SetupRoutes(router)

	// Jalankan server
	router.Run(":8080")
}
