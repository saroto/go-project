package controller

import (
	"go/goRoutine/Internal/auth"
	"go/goRoutine/config"
	"go/goRoutine/models"
	"go/goRoutine/types"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var body types.UserRegisterRequest
	c.BindJSON(&body)
	if body.Email == "" || body.Name == "" || body.Password == "" {
		c.JSON(400, gin.H{
			"message": "All fields are requires",
		})
		return
	}
	var existingUser models.User
	if err := config.DB.Where("email = ?", body.Email).First(&existingUser).Error; err == nil {
		c.JSON(400, gin.H{
			"message": "User with this email already registered",
		})
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("err")
	}

	post := models.User{Name: body.Name, Email: body.Email, Password: string(hashPassword), CreatedAt: time.Now(), UpdatedAt: time.Now()}

	if err = config.DB.Create(&post).Error; err != nil {
		c.JSON(500, gin.H{"message": "Failed to register the user", "err": err.Error()})
		return
	}
	token, err := auth.IssuseToken(int(post.ID), post.Email, post.Name)
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to create token"})
		return
	}
	data := types.UserRegisterResponse{
		Name:  post.Name,
		Email: post.Email,
		Token: token,
	}
	c.JSON(200, gin.H{
		"message": "Successfully Register",
		"data":    data,
	})
}

func Login(c *gin.Context) {
	var user models.User
	var reqBody types.UserLoginRequest
	c.BindJSON(&reqBody)
	if reqBody.Email == "" || reqBody.Password == "" {
		c.JSON(400, gin.H{
			"message": "All fields are require",
		})
		return
	}
	if err := config.DB.Select("id", "email", "password").Where("email = ?", reqBody.Email).First(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Email is not found",
		})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password))

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Password is not found",
		})
		return
	}

}
