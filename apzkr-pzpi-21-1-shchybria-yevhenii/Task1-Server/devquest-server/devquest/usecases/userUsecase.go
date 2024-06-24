package usecases

import (
	"devquest-server/devquest/domain/entities"
	"devquest-server/devquest/domain/models"
	"devquest-server/devquest/domain/repositories"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepo repositories.UserRepo
	companyRepo repositories.CompanyRepo
}

func NewUserUsecase(userRepo repositories.UserRepo, companyRepo repositories.CompanyRepo) *UserUsecase {
	return &UserUsecase{userRepo: userRepo, companyRepo: companyRepo}
}

func (u *UserUsecase) RegisterUser(userRegisterInfo models.RegisterUserDTO) (*models.JwtUserDTO, error) {
	existingUser, err := u.userRepo.GetUserByUsername(userRegisterInfo.Username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username is taken")
	}

	company, err := u.companyRepo.GetCompanyByID(userRegisterInfo.CompanyID)
	if err != nil {
		return nil, err
	}
	if company == nil {
		return nil, errors.New("company does not exist")
	}

	role, err := u.userRepo.GetRoleByID(userRegisterInfo.RoleID)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, errors.New("role does not exist")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userRegisterInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var registeredUser = models.InsertUserDTO{
		ID: uuid.New(),
		Username: userRegisterInfo.Username,
		FirstName: userRegisterInfo.FirstName,
		LastName: userRegisterInfo.LastName,
		PasswordHash: string(passwordHash),
		RoleID: userRegisterInfo.RoleID,
		CompanyID: userRegisterInfo.CompanyID,
	}

	err = u.userRepo.InsertUser(&registeredUser)
	if err != nil {
		return nil, err
	}

	return &models.JwtUserDTO{
		ID: registeredUser.ID,
		Username: registeredUser.Username,
		RoleTitle: role.Title,
	}, nil
}

func (u *UserUsecase) LoginUser(userLoginInfo models.LoginUserDTO) (*models.JwtUserDTO, error) {
	existingUser, err := u.userRepo.GetUserByUsername(userLoginInfo.Username)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.New("user does not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(userLoginInfo.Password))
	if err != nil {
		return nil, err
	}

	role, err := u.userRepo.GetRoleByID(existingUser.RoleID)
	if err != nil {
		return nil, err
	}

	return &models.JwtUserDTO{
		ID: existingUser.ID,
		Username: existingUser.Username,
		RoleTitle: role.Title,
	}, nil
}

func (u *UserUsecase) GetJwtUserByID(userID uuid.UUID) (*models.JwtUserDTO, error) {
	user, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	role, err := u.userRepo.GetRoleByID(user.RoleID)
	if err != nil {
		return nil, err
	}

	jwtUser := &models.JwtUserDTO{
		ID: user.ID,
		Username: user.Username,
		RoleTitle: role.Title,
	}

	return jwtUser, err
}

func (u *UserUsecase) GetDevelopersForManager(managerID uuid.UUID) ([]*entities.User, error) {
	manager, err := u.userRepo.GetUserByID(managerID)
	if err != nil {
		return nil, err
	}
	if manager == nil {
		return nil, errors.New("user does not exist")
	}
	
	company, err := u.companyRepo.GetCompanyByID(manager.CompanyID)
	if err != nil {
		return nil, err
	}
	if company == nil {
		return nil, errors.New("company does not exist")
	}

	developers, err := u.userRepo.GetDevelopersByCompany(manager.CompanyID)
	if err != nil {
		return nil, err
	}

	return developers, nil
}

func (u *UserUsecase) GetRolesForRegistration() ([]*entities.Role, error) {
	roles, err := u.userRepo.GetRolesForRegistration()
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (u *UserUsecase) GetUserByID(userID uuid.UUID) (*entities.User, error) {
	user, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *UserUsecase) GetRoleByID(roleID uuid.UUID) (*entities.Role, error) {
	role, err := u.userRepo.GetRoleByID(roleID)
	if err != nil {
		return nil, err
	}

	return role, err
}