package models

import "github.com/google/uuid"

type (
	CreateAchievementDTO struct {
		Name      string `json:"name"`
		Description string `json:"description"`
		Points    int    `json:"points"`
		ProjectID uuid.UUID `json:"project_id"`
	}

	UpdateAchievementDTO struct {
		Name      string `json:"name"`
		Description string `json:"description"`
		Points    int    `json:"points"`
	}
)