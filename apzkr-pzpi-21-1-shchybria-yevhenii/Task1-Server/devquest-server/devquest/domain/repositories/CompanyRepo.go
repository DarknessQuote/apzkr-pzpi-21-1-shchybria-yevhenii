package repositories

import (
	"devquest-server/devquest/domain/entities"

	"github.com/google/uuid"
)

type CompanyRepo interface {
	GetAllCompanies() ([]*entities.Company, error)
	GetCompanyByID(companyID uuid.UUID) (*entities.Company, error)
	AddCompany(entities.Company) (*entities.Company, error)
	UpdateCompany(*entities.Company) error
	DeleteCompany(companyID uuid.UUID) error
}