package repositories

import (
	"devquest-server/devquest/domain/entities"
	"devquest-server/devquest/domain/models"

	"github.com/google/uuid"
)

type ProjectRepo interface {
	GetProjectByID(projectID uuid.UUID) (*entities.Project, error)
	AddProject(newProject entities.Project) error
	UpdateProject(projectID uuid.UUID, updatedProject models.UpdateProjectDTO) error
	DeleteProject(projectID uuid.UUID) error
	
	GetProjectsOfManager(managerID uuid.UUID) ([]*entities.Project, error)
	GetProjectsOfDeveloper(developerID uuid.UUID) ([]*entities.Project, error)
	GetProjectDevelopers(projectID uuid.UUID) ([]*models.DeveloperProjectInfoDTO, error)
	CheckDeveloperOnProject(projectID uuid.UUID, developerID uuid.UUID) (bool, error)
	AddDeveloperToProject(projectID uuid.UUID, developerID uuid.UUID) error
	RemoveDeveloperFromProject(projectID uuid.UUID, developerID uuid.UUID) error
	UpdateDeveloperProjectPoints(projectID uuid.UUID, developerID uuid.UUID, points int) error
}