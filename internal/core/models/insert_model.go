package models

type InsertRequest struct {
	Collection string      `json:"collection" validate:"required"`
	Data       interface{} `json:"data" validate:"required"`
}
