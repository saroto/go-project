package main

import (
	"go/goRoutine/config"
	"go/goRoutine/models"
)

func init() {
	config.LoadEnvVariable()
	config.DbConnection()
}

func main() {
	config.DB.AutoMigrate(&models.User{}, &models.Token{}, &models.Otp{})

}
