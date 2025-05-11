package models

type RegistrationRequest struct {
	OrgName string `json:"org_name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
}

type RegistrationResponse struct {
	APIKey string `json:"api_key"`
}
