package postgres

import (
	"context"
	"database/sql"
	"devquest-server/devquest/domain/entities"
	"devquest-server/devquest/domain/models"
	"devquest-server/devquest/domain/repositories"
	"devquest-server/devquest/infrastructure"

	"github.com/google/uuid"
)

type ProjectPostgresRepo struct {
	db infrastructure.Database
}

func NewProjectPostgresRepo(db infrastructure.Database) repositories.ProjectRepo {
	return &ProjectPostgresRepo{db: db}
}

func (p *ProjectPostgresRepo) GetProjectByID(projectID uuid.UUID) (*entities.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), p.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id, name, description, company_id, manager_id
		FROM projects
		WHERE id = $1
	`

	row := p.db.GetDB().QueryRowContext(ctx, query, projectID)

	var project entities.Project
	err := row.Scan(&project.ID, &project.Name, &project.Description, &project.CompanyID, &project.ManagerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
			return nil, err
	}

	return &project, nil
}

func (p *ProjectPostgresRepo) AddProject(project entities.Project) error {
	ctx, cancel := context.WithTimeout(context.Background(), p.db.GetDBTimeout())
	defer cancel()

	execute := `
		INSERT INTO projects
		(id, name, description, company_id, manager_id)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := p.db.GetDB().ExecContext(ctx, execute, project.ID, project.Name, project.Description, project.CompanyID, project.ManagerID)
	if err != nil {
			return err
	}

	return nil
}

func (p *ProjectPostgresRepo) UpdateProject(projectID uuid.UUID, updatedProject models.UpdateProjectDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), p.db.GetDBTimeout())
	defer cancel()

	execute := `
		UPDATE projects
		SET name = $1, description = $2
		WHERE id = $3
	`

	_, err := p.db.GetDB().ExecContext(ctx, execute, updatedProject.Name, updatedProject.Description, projectID)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectPostgresRepo) DeleteProject(projectID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), p.db.GetDBTimeout())
	defer cancel()

	execute := `
		DELETE FROM projects
		WHERE id = $1
	`

	_, err := p.db.GetDB().ExecContext(ctx, execute, projectID)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectPostgresRepo) GetProjectsOfManager(managerID uuid.UUID) ([]*entities.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), p.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id, name, description, company_id, manager_id
		FROM projects
		WHERE manager_id = $1
		ORDER BY name
	`

	rows, err := p.db.GetDB().QueryContext(ctx, query, managerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*entities.Project
	for rows.Next() {
		var project entities.Project

		err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.CompanyID, &project.ManagerID)
		if err != nil {
			return nil, err
		}

		projects = append(projects, &project)
	}

	return projects, nil
}

func (p *ProjectPostgresRepo) GetProjectsOfDeveloper(developerID uuid.UUID) ([]*entities.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), p.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id, name, description, company_id, manager_id
		FROM projects
		WHERE id in (SELECT project_id FROM projects_users WHERE developer_id = $1)
		ORDER BY name
	`

	rows, err := p.db.GetDB().QueryContext(ctx, query, developerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*entities.Project
	for rows.Next() {
		var project entities.Project

		err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.CompanyID, &project.ManagerID)
		if err != nil {
			return nil, err
		}

		projects = append(projects, &project)
	}

	return projects, nil
}

func (p *ProjectPostgresRepo) GetProjectDevelopers(projectID uuid.UUID) ([]*models.DeveloperProjectInfoDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), p.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT u.id, u.username, u.first_name, u.last_name, u.role_id, u.company_id, pu.points
		FROM users u
		LEFT JOIN projects_users pu ON u.id = pu.developer_id
		WHERE pu.project_id = $1
		ORDER BY u.username
	`

	rows, err := p.db.GetDB().QueryContext(ctx, query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var developers []*models.DeveloperProjectInfoDTO
	for rows.Next() {
		var developer models.DeveloperProjectInfoDTO

		err := rows.Scan(&developer.ID, &developer.Username, &developer.FirstName, &developer.LastName, &developer.RoleID, &developer.CompanyID, &developer.Points)
		if err != nil {
			return nil, err
		}

		developers = append(developers, &developer)
	}

	return developers, nil
}

func (p *ProjectPostgresRepo) CheckDeveloperOnProject(projectID uuid.UUID, developerID uuid.UUID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), p.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT developer_id
		FROM projects_users
		WHERE project_id = $1 AND developer_id = $2
	`

	row := p.db.GetDB().QueryRowContext(ctx, query, projectID, developerID)

	var developerOnProjectID uuid.UUID
	err := row.Scan(&developerOnProjectID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
			return false, err
	}

	return true, nil
}

func (p *ProjectPostgresRepo) AddDeveloperToProject(projectID uuid.UUID, developerID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), p.db.GetDBTimeout())
	defer cancel()

	execute := `
		INSERT INTO projects_users
		(project_id, developer_id, points)
		VALUES ($1, $2, 0)
	`

	_, err := p.db.GetDB().ExecContext(ctx, execute, projectID, developerID)
	if err != nil {
			return err
	}

	return nil
}

func (p *ProjectPostgresRepo) RemoveDeveloperFromProject(projectID uuid.UUID, developerID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), p.db.GetDBTimeout())
	defer cancel()

	execute := `
		DELETE FROM projects_users
		WHERE project_id = $1 AND developer_id = $2
	`

	_, err := p.db.GetDB().ExecContext(ctx, execute, projectID, developerID)
	if err != nil {
			return err
	}

	return nil
}

func (p *ProjectPostgresRepo) UpdateDeveloperProjectPoints(projectID uuid.UUID, developerID uuid.UUID, points int) error {
	ctx, cancel := context.WithTimeout(context.Background(), p.db.GetDBTimeout())
	defer cancel()

	execute := `
		UPDATE projects_users
		SET points = points + $1
		WHERE project_id = $2 AND developer_id = $3
	`

	_, err := p.db.GetDB().ExecContext(ctx, execute, points, projectID, developerID)
	if err != nil {
		return err
	}

	return nil
}