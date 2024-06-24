package entities

import "github.com/google/uuid"

type Achievement struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Points int `json:"points"`
	ProjectID uuid.UUID `json:"project_id"`
}