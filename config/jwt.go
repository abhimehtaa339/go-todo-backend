package config

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"time"
)

var jwtSecret = GetEnv("JWT_SECRET", "")

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
	accessSigned, err := accessToken.SignedString([]byte(jwtSecret))
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
	refreshSigned, err := refreshToken.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Printf("Error generating refresh signed token: %v", err)
		return nil, err
	}

	return &TokenPair{
		accessSigned,
		refreshSigned,
	}, nil
}
