package config

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"time"
)

var jwtSecret = []byte(GetEnv("JWT_SECRET", ""))

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateTokenPairs(userId uint) (*TokenPair, error) {
	//Access Token
	accessClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(AccessJWTExpiryDuration).Unix(),
		"type":    "access",
		"jti":     uuid.New().String(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessSigned, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		log.Printf("Error generating access signed token: %v", err)
		return nil, err
	}

	//Refresh Token
	refreshClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(RefreshJWTExpiryDuration).Unix(),
		"type":    "refresh",
		"jti":     uuid.New().String(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshSigned, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		log.Printf("Error generating refresh signed token: %v", err)
		return nil, err
	}

	return &TokenPair{
		accessSigned,
		refreshSigned,
	}, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	// Parse into claims so expiration / nbf are validated for us.
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// enforce expected signing method explicitly
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		log.Printf("token parse error: %v", err)
		return nil, err
	}

	// Verify token is valid according to claims (exp / nbf, signature)
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
