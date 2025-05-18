package auth

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sahay-shashank/mongodb-server/internal/core/details"
	"github.com/sahay-shashank/mongodb-server/internal/core/models"
)

func GenerateJWTToken(tenantID string, serviceString string) details.APIDetails {
	claim := &models.CustomClaim{
		TenantID: tenantID,
		Service:  serviceString,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	result, err := token.SignedString(jwtSecret)

	if err != nil {
		log.Print(err)
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.JWTTokenFailure,
			Message:           details.GetMessage(details.JWTTokenFailure),
			AdditionalDetails: err,
		}
	}
	return details.APIDetails{
		StatusCode: details.JWTTokenSuccessful,
		Message:    details.GetMessage(details.JWTTokenSuccessful),
		AdditionalDetails: models.AuthResponse{
			Token: result,
		},
	}
}
