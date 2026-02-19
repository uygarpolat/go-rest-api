package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type tokenClaims struct {
	Email  string `json:"email"`
	UserID int64  `json:"userId"`
	jwt.RegisteredClaims
}

func getJWTSecret() ([]byte, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET is not set")
	}
	return []byte(secret), nil
}

func GenerateToken(email string, userId int64) (string, error) {
	jwtSecret, err := getJWTSecret()
	if err != nil {
		return "", err
	}
	claims := tokenClaims{
		Email:  email,
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("sign token: %w", err)
	}
	return signedToken, nil
}

func VerifyToken(tokenString string) (int64, error) {
	jwtSecret, err := getJWTSecret()
	if err != nil {
		return 0, err
	}

	claims := &tokenClaims{}
	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Method.Alg())
		}
		return jwtSecret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return 0, fmt.Errorf("parse token: %w", err)
	}

	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	if claims.UserID == 0 {
		return 0, errors.New("invalid token claims")
	}

	return claims.UserID, nil
}
