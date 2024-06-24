package models

import "github.com/google/uuid"

type (
	RegisterUserDTO struct {
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
		RoleID    uuid.UUID `json:"role_id"`
		CompanyID uuid.UUID `json:"company_id"`
	}

	LoginUserDTO struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
	}

	JwtUserDTO struct {
		ID uuid.UUID `json:"id"`
		Username string `json:"username"`
		RoleTitle string `json:"role"`
	}

	InsertUserDTO struct {
		ID uuid.UUID `json:"id"`
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		PasswordHash  string `json:"password_hash"`
		RoleID    uuid.UUID `json:"role_id"`
		CompanyID uuid.UUID `json:"company_id"`
	}

	UpdateUserDTO struct {
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
	}

	DeveloperProjectInfoDTO struct {
		ID uuid.UUID `json:"id"`
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		RoleID    uuid.UUID `json:"role_id"`
		CompanyID uuid.UUID `json:"company_id"`
		Points int `json:"points"`
	}
)