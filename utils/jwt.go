package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtSecret []byte

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

func GenerateToken(email string, userId int64) (string, error) {
	if len(jwtSecret) == 0 {
		return "", errors.New("JWT_SECRET is not set")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString(jwtSecret)
}
