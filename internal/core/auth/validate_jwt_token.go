package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sahay-shashank/mongodb-server/internal/core/models"
)

func ValidateJWT(tokenString string) (*models.CustomClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.CustomClaim); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
