package repositories

import (
	"devquest-server/devquest/domain/entities"
	"devquest-server/devquest/domain/models"

	"github.com/google/uuid"
)

type AchievementRepo interface {
	GetAchievementByID(achievementID uuid.UUID) (*entities.Achievement, error)
	GetProjectAchievements(projectID uuid.UUID) ([]*entities.Achievement, error)
	GetDevelopersAchievements(developerID uuid.UUID) ([]*entities.Achievement, error)
	AddAchievement(newAchievement entities.Achievement) error
	UpdateAchievement(achievementID uuid.UUID, updatedAchievement models.UpdateAchievementDTO) error
	DeleteAchievement(achievementID uuid.UUID) error
	CheckAchievementOnProject(projectID uuid.UUID, achievementID uuid.UUID) (bool, error)
	GiveAchievementToDeveloper(achievementID uuid.UUID, developerID uuid.UUID) error
}