package entities

import "github.com/google/uuid"

type (
	User struct {
		ID uuid.UUID `json:"id"`
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Username string `json:"username"`
		PasswordHash string `json:"password_hash"`
		RoleID uuid.UUID `json:"role_id"`
		CompanyID uuid.UUID `json:"company_id"`
		Points int `json:"points"`
	}

  Role struct {
		ID uuid.UUID `json:"id"`
		Title string `json:"title"`
	}
)