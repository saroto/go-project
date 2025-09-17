package main

import (
	"go/goRoutine/config"
	"go/goRoutine/controller"
	"go/goRoutine/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariable()
	config.DbConnection()
}

func main() {

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	api := router.Group("/api")
	api.POST("/register", controller.RegisterUser)
	api.POST("/login", controller.Login)
	api.POST("/otp", controller.RequestOtpCode)
	api.POST("/verify-otp", controller.VerifyOtpCode)
	authRouter := api.Group("/auth", middleware.AuthMiddleware())
	authRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
