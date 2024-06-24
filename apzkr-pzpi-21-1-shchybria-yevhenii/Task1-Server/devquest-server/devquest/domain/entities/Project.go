package entities

import "github.com/google/uuid"

type Project struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	CompanyID uuid.UUID `json:"company_id"`
	ManagerID uuid.UUID `json:"manager_id"`
}