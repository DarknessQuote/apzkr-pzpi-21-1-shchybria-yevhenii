package postgres

import (
	"context"
	"database/sql"
	"devquest-server/devquest/domain/entities"
	"devquest-server/devquest/domain/repositories"
	"devquest-server/devquest/infrastructure"

	"github.com/google/uuid"
)

type companyPostgresRepo struct {
	db infrastructure.Database
}

func NewCompanyPostgresRepo(db infrastructure.Database) repositories.CompanyRepo {
	return &companyPostgresRepo{db: db}
}

func (c *companyPostgresRepo) GetAllCompanies() ([]*entities.Company, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id, name, owner, email
		FROM companies
		ORDER BY name
	`

	rows, err := c.db.GetDB().QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []*entities.Company
	for rows.Next() {
		var company entities.Company

		err := rows.Scan(&company.ID, &company.Name, &company.Owner, &company.Email)
		if err != nil {
			return nil, err
		}

		companies = append(companies, &company)
	}

	return companies, nil
}

func (c *companyPostgresRepo) GetCompanyByID(companyID uuid.UUID) (*entities.Company, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id, name, owner, email
		FROM companies
		WHERE id = $1
	`

	row := c.db.GetDB().QueryRowContext(ctx, query, companyID)

	var company entities.Company
	err := row.Scan(&company.ID, &company.Name, &company.Owner, &company.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
			return nil, err
	}

	return &company, nil
}

func (c *companyPostgresRepo) AddCompany(company entities.Company) (*entities.Company, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.db.GetDBTimeout())
	defer cancel()

	execute := `
		INSERT INTO companies
		(id, name, owner, email)
		VALUES ($1, $2, $3, $4)
	`

	_, err := c.db.GetDB().ExecContext(ctx, execute, company.ID, company.Name, company.Owner, company.Email)
	if err != nil {
			return nil, err
	}

	return &company, nil
}

func (c *companyPostgresRepo) UpdateCompany(company *entities.Company) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.db.GetDBTimeout())
	defer cancel()

	execute := `
		UPDATE companies
		SET name = $1, owner = $2, email = $3
		WHERE id = $4
	`

	_, err := c.db.GetDB().ExecContext(ctx, execute, company.Name, company.Owner, company.Email, company.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c *companyPostgresRepo) DeleteCompany(companyID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.db.GetDBTimeout())
	defer cancel()

	execute := `
		DELETE FROM companies
		WHERE id = $1
	`

	_, err := c.db.GetDB().ExecContext(ctx, execute, companyID)
	if err != nil {
		return err
	}

	return nil
}