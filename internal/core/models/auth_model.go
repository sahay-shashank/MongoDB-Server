package models

import "github.com/golang-jwt/jwt/v5"

type AuthRequest struct {
	APIKey string `json:"api_key" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type CustomClaim struct {
	TenantID string `json:"tenant_id"`
	Service  string `json:"service"`
	jwt.RegisteredClaims
}
