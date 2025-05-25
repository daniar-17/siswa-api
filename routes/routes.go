package routes

import (
	"siswa-api/controllers"
	"siswa-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/students")
	auth.Use(middleware.JWTAuth())
	{
		auth.POST("", controllers.CreateStudent)
		auth.GET("", controllers.GetStudents)
		auth.PUT("/:id", controllers.UpdateStudent)
		auth.DELETE("/:id", controllers.DeleteStudent)

		auth.GET("/specialist", controllers.CreateSpecialist)
	}
}
