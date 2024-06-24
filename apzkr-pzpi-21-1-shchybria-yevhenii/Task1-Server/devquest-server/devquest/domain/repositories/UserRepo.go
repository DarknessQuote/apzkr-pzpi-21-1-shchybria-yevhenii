package repositories

import (
	"devquest-server/devquest/domain/entities"
	"devquest-server/devquest/domain/models"

	"github.com/google/uuid"
)

type UserRepo interface {
	GetUserByID(id uuid.UUID) (*entities.User, error)
	GetUserByUsername(username string) (*entities.User, error)
	GetDevelopersByCompany(companyID uuid.UUID) ([]*entities.User, error)
	InsertUser(user *models.InsertUserDTO) error
	CheckUserRole(userID uuid.UUID, roleTitle string) (bool, error)
	
	GetRolesForRegistration() ([]*entities.Role, error)
	GetRoleByID(roleID uuid.UUID) (*entities.Role, error)
}