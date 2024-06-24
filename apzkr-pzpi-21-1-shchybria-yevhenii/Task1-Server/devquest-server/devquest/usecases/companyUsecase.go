package usecases

import (
	"devquest-server/devquest/domain/entities"
	"devquest-server/devquest/domain/models"
	"devquest-server/devquest/domain/repositories"
	"errors"

	"github.com/google/uuid"
)

type CompanyUsecase struct {
	companyRepo repositories.CompanyRepo
}

func NewCompanyUsecase(companyRepo repositories.CompanyRepo) *CompanyUsecase {
	return &CompanyUsecase{companyRepo: companyRepo}
}

func (c *CompanyUsecase) FindAllCompanies() ([]*entities.Company, error) {
	allCompanies, err := c.companyRepo.GetAllCompanies()
	if err != nil {
		return nil, err
	}

	return allCompanies, nil
}

func (c *CompanyUsecase) FindCompanyByID(companyID uuid.UUID) (*entities.Company, error) {
	company, err := c.companyRepo.GetCompanyByID(companyID)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (c *CompanyUsecase) CreateNewCompany(companyForInsert *models.InsertCompanyDTO) (*entities.Company, error) {
	newCompany := &entities.Company{
		ID: uuid.New(),
		Name: companyForInsert.Name,
		Owner: companyForInsert.Owner,
		Email: companyForInsert.Email,
	}

	company, err := c.companyRepo.AddCompany(*newCompany)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (c *CompanyUsecase) UpdateCompany(companyID uuid.UUID, updatedCompany models.UpdateCompanyDTO) error {
	companyToUpdate, err := c.companyRepo.GetCompanyByID(companyID)
	if err != nil {
		return err
	}
	if companyToUpdate == nil {
		return errors.New("company does not exist")
	}

	companyToUpdate.Name = updatedCompany.Name
	companyToUpdate.Owner = updatedCompany.Owner
	companyToUpdate.Email = updatedCompany.Email

	err = c.companyRepo.UpdateCompany(companyToUpdate)
	return err
}

func (c *CompanyUsecase) DeleteCompany(companyID uuid.UUID) error {
	err := c.companyRepo.DeleteCompany(companyID)
	return err
}
