package controller

import (
	auth "go/goRoutine/Internal/otpCode"
	"go/goRoutine/config"
	"go/goRoutine/mail"
	"go/goRoutine/models"
	"go/goRoutine/types"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestOtpCode(c *gin.Context) {
	var req types.OTPReqBody
	// var otpModel models.Otp
	// Parse and validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are require"})
		return
	}
	otp := auth.GenereateOtpCode(6)
	otpCodeParse := models.Otp{
		UserID:    uint(req.Id),
		OtpCode:   otp,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
	if err := config.DB.Create(&otpCodeParse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to stop otp",
		})
		return
	}
	err := mail.SendMail(req.Email, "OTP Code", otp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to sent email",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Email sent successfully",
	})
}

func VerifyOtpCode(c *gin.Context) {
	var otpModel models.Otp
	var reqBody types.VerifyOtpCode

	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are require"})
		return
	}

	if err := config.DB.Where("user_id = ? AND otp_code =?", reqBody.UserId, reqBody.OtpCode).First(&otpModel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP code"})
		return
	}
	if time.Now().After(otpModel.ExpiresAt) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "OTP code has expired"})
		return
	}
	config.DB.Delete(&otpModel)
	c.JSON(http.StatusOK, gin.H{"message": "OTP code verified successfully"})
}
