package models

type (
	InsertCompanyDTO struct {
		Name  string `json:"name"`
		Owner string `json:"owner"`
		Email string `json:"email"`
	}

	UpdateCompanyDTO struct {
		Name  string `json:"name"`
		Owner string `json:"owner"`
		Email string `json:"email"`
	}
)