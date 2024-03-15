package main

import (
	"login-test/controllers"
	"login-test/middlewares"
	"login-test/models"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	err := models.ConnectDataBase()
	if err != nil {
		os.Exit(1)
	}

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := router.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	router.Run(":8000")
}
