package models

import "github.com/google/uuid"

type (
	CreateProjectDTO struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		CompanyID   uuid.UUID `json:"company_id"`
		ManagerID   uuid.UUID `json:"manager_id"`
	}

	UpdateProjectDTO struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)