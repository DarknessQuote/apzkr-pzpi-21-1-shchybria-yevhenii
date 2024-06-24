package entities

import "github.com/google/uuid"

type Company struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Owner string `json:"owner"`
	Email string `json:"email"`
}