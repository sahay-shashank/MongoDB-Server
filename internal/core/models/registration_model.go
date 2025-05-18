package models

type RegistrationRequest struct {
	OrgName string `json:"org_name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Service string `json:"service" validate:"required"`
}

type RegistrationResponse struct {
	APIKey string `json:"api_key"`
}
