package usecases

import (
	"devquest-server/devquest/domain/entities"
	"devquest-server/devquest/domain/models"
	"devquest-server/devquest/domain/repositories"
	"errors"

	"github.com/google/uuid"
)

type AchievementUsecase struct {
	achievementRepo repositories.AchievementRepo
	projectRepo     repositories.ProjectRepo
	userRepo repositories.UserRepo
}

func NewAchievementUsecase(aRepo repositories.AchievementRepo, pRepo repositories.ProjectRepo, uRepo repositories.UserRepo) *AchievementUsecase {
	return &AchievementUsecase{achievementRepo: aRepo, projectRepo: pRepo, userRepo: uRepo}
}

func (a *AchievementUsecase) GetProjectAchievements(projectID uuid.UUID) ([]*entities.Achievement, error) {
	existingProject, err := a.projectRepo.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}
	if existingProject == nil {
		return nil, errors.New("project does not exist")
	}

	projectAchievements, err := a.achievementRepo.GetProjectAchievements(projectID)
	if err != nil {
		return nil, err
	}

	return projectAchievements, nil
}

func (a *AchievementUsecase) GetDeveloperAchievements(developerID uuid.UUID) ([]*entities.Achievement, error) {
	developerAchievements, err := a.achievementRepo.GetDevelopersAchievements(developerID)
	if err != nil {
		return nil, err
	}

	return developerAchievements, nil
}

func (a *AchievementUsecase) AddAchievementToProject(newAchievement models.CreateAchievementDTO) error {
	existingProject, err := a.projectRepo.GetProjectByID(newAchievement.ProjectID)
	if err != nil {
		return err
	}
	if existingProject == nil {
		return errors.New("project does not exist")
	}

	addAchievement := &entities.Achievement{
		ID: uuid.New(),
		Name: newAchievement.Name,
		Description: newAchievement.Description,
		Points: newAchievement.Points,
		ProjectID: newAchievement.ProjectID,
	}

	err = a.achievementRepo.AddAchievement(*addAchievement)
	if err != nil {
		return err
	}

	return nil
}

func (a *AchievementUsecase) UpdateAchievement(achievementID uuid.UUID, updatedAchievement models.UpdateAchievementDTO) error {
	achievementToUpdate, err := a.achievementRepo.GetAchievementByID(achievementID)
	if err != nil {
		return err
	}
	if achievementToUpdate == nil {
		return errors.New("achievement does not exist")
	}

	err = a.achievementRepo.UpdateAchievement(achievementID, updatedAchievement)
	if err != nil {
		return err
	}

	return nil
}

func (a *AchievementUsecase) DeleteAchievement(achievementID uuid.UUID) error {
	err := a.achievementRepo.DeleteAchievement(achievementID)
	if err != nil {
		return err
	}

	return nil
}

func (a *AchievementUsecase) GiveAchievementToDeveloper(projectID uuid.UUID, achievementID uuid.UUID, developerID uuid.UUID) error {
	existingProject, err := a.projectRepo.GetProjectByID(projectID)
	if err != nil {
		return err
	}
	if existingProject == nil {
		return errors.New("project does not exist")
	}

	isDeveloper, err := a.userRepo.CheckUserRole(developerID, "Developer")
	if err != nil {
		return err
	}
	if !isDeveloper {
		return errors.New("only developers can get achievements")
	}

	achievementIsOnProject, err := a.achievementRepo.CheckAchievementOnProject(projectID, achievementID)
	if err != nil {
		return err
	}
	if !achievementIsOnProject {
		return errors.New("achievement is from other project")
	}

	developerIsOnProject, err := a.projectRepo.CheckDeveloperOnProject(projectID, developerID)
	if err != nil {
		return err
	}
	if !developerIsOnProject {
		return errors.New("developer is from other project")
	}

	err = a.achievementRepo.GiveAchievementToDeveloper(achievementID, developerID)
	if err != nil {
		return err
	}

	achievement, err := a.achievementRepo.GetAchievementByID(achievementID)
	if err != nil {
		return err
	}

	err = a.projectRepo.UpdateDeveloperProjectPoints(projectID, developerID, achievement.Points)
	if err != nil {
		return err
	}

	return nil
}