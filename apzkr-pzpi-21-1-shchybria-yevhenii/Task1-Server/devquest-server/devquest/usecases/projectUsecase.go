package usecases

import (
	"devquest-server/devquest/domain/entities"
	"devquest-server/devquest/domain/models"
	"devquest-server/devquest/domain/repositories"
	"errors"

	"github.com/google/uuid"
)

type ProjectUsecase struct {
	projectRepo repositories.ProjectRepo
	userRepo repositories.UserRepo
	companyRepo repositories.CompanyRepo
}

func NewProjectUsecase(pRepo repositories.ProjectRepo, uRepo repositories.UserRepo, cRepo repositories.CompanyRepo) *ProjectUsecase {
	return &ProjectUsecase{projectRepo: pRepo, userRepo: uRepo, companyRepo: cRepo}
}

func (p *ProjectUsecase) GetProjectByID(projectID uuid.UUID) (*entities.Project, error) {
	project, err := p.projectRepo.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (p *ProjectUsecase) CreateNewProject(newProject models.CreateProjectDTO) error {
	existingCompany, err := p.companyRepo.GetCompanyByID(newProject.CompanyID)
	if err != nil {
		return err
	}
	if existingCompany == nil {
		return errors.New("company does not exist")
	}

	isManager, err := p.userRepo.CheckUserRole(newProject.ManagerID, "Manager")
	if err != nil {
		return err
	}
	if !isManager {
		return errors.New("no permission to create projects")
	}

	addProject := &entities.Project{
		ID: uuid.New(),
		Name: newProject.Name,
		Description: newProject.Description,
		CompanyID: newProject.CompanyID,
		ManagerID: newProject.ManagerID,
	}

	err = p.projectRepo.AddProject(*addProject)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectUsecase) UpdateProject(projectID uuid.UUID, updatedProject models.UpdateProjectDTO) error {
	projectToUpdate, err := p.projectRepo.GetProjectByID(projectID)
	if err != nil {
		return err
	}
	if projectToUpdate == nil {
		return errors.New("project does not exist")
	}

	err = p.projectRepo.UpdateProject(projectID, updatedProject)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectUsecase) DeleteProject(projectID uuid.UUID) error {
	err := p.projectRepo.DeleteProject(projectID)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectUsecase) GetManagerProjects(managerID uuid.UUID) ([]*entities.Project, error) {
	isManager, err := p.userRepo.CheckUserRole(managerID, "Manager")
	if err != nil {
		return nil, err
	}
	if !isManager {
		return nil, errors.New("incorrect role")
	}

	managerProjects, err := p.projectRepo.GetProjectsOfManager(managerID)
	if err != nil {
		return nil, err
	}

	return managerProjects, nil
}

func (p *ProjectUsecase) GetDeveloperProjects(developerID uuid.UUID) ([]*entities.Project, error) {
	isDeveloper, err := p.userRepo.CheckUserRole(developerID, "Developer")
	if err != nil {
		return nil, err
	}
	if !isDeveloper {
		return nil, errors.New("incorrect role")
	}

	developerProjects, err := p.projectRepo.GetProjectsOfDeveloper(developerID)
	if err != nil {
		return nil, err
	}

	return developerProjects, nil
}

func (p *ProjectUsecase) GetProjectDevelopers(projectID uuid.UUID) ([]*models.DeveloperProjectInfoDTO, error) {
	existingProject, err := p.projectRepo.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}
	if existingProject == nil {
		return nil, errors.New("project does not exist")
	}

	projectDevelopers, err := p.projectRepo.GetProjectDevelopers(projectID)
	if err != nil {
		return nil, err
	}

	return projectDevelopers, nil
}

func (p *ProjectUsecase) AddDeveloperToProject(projectID uuid.UUID, developerID uuid.UUID) error {
	existingProject, err := p.projectRepo.GetProjectByID(projectID)
	if err != nil {
		return err
	}
	if existingProject == nil {
		return errors.New("project does not exist")
	}

	isDeveloper, err := p.userRepo.CheckUserRole(developerID, "Developer")
	if err != nil {
		return err
	}
	if !isDeveloper {
		return errors.New("only developers can be added to projects")
	}

	isOnProject, err := p.projectRepo.CheckDeveloperOnProject(projectID, developerID)
	if err != nil {
		return err
	}
	if isOnProject {
		return errors.New("developer is already on the project")
	}

	err = p.projectRepo.AddDeveloperToProject(projectID, developerID)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectUsecase) RemoveDeveloperFromProject(projectID uuid.UUID, developerID uuid.UUID) error {
	existingProject, err := p.projectRepo.GetProjectByID(projectID)
	if err != nil {
		return err
	}
	if existingProject == nil {
		return errors.New("project does not exist")
	}

	isDeveloper, err := p.userRepo.CheckUserRole(developerID, "Developer")
	if err != nil {
		return err
	}
	if !isDeveloper {
		return errors.New("only developers can be removed from projects")
	}

	isOnProject, err := p.projectRepo.CheckDeveloperOnProject(projectID, developerID)
	if err != nil {
		return err
	}
	if !isOnProject {
		return errors.New("developer is not on the project")
	}

	err = p.projectRepo.RemoveDeveloperFromProject(projectID, developerID)
	if err != nil {
		return err
	}

	return nil
}