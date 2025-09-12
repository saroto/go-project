package main

import (
	"go/goRoutine/config"
	"go/goRoutine/controller"

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
	router.POST("/register", controller.RegisterUser)
	router.POST("/login", controller.Login)
	router.Run()
}
