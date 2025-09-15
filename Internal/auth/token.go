package auth

import (
	"fmt"
	"go/goRoutine/config"
	"go/goRoutine/models"
	"go/goRoutine/types"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func IssuseToken(id int, email string, name string) (string, error) {
	var tokenData *models.Token
	var expiredDate = time.Now().Add(24 * time.Hour)

	// claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"iss":   name,
	// 	"exp":   expiredDate,
	// 	"iat":   time.Now().Unix(),
	// 	"email": email,
	// 	"id":    id,
	// })
	claim := &types.Claims{
		Email: email,
		Id:    uint64(id),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    name,
			ExpiresAt: jwt.NewNumericDate(expiredDate),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal(err)
	}
	tokenData = &models.Token{
		UserId:      id,
		Token:       tokenString,
		ExpiredDate: expiredDate.Unix(),
	}
	if err = config.DB.Create(tokenData).Error; err != nil {
		log.Fatal("Error", err)
	}
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("token invalid")
	}
	return nil
}
