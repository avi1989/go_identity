package services

import (
	"database/sql"
	"errors"
	"github.com/avi1989/Identity/models"
	"github.com/avi1989/Identity/stores"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userStore       stores.IUserStore
	permissionStore stores.PermissionStore
}

var InvalidLogin = errors.New("Invalid Login")

func transformPermissions(permissions []models.Permission) []string {
	vsm := make([]string, len(permissions))
	for i, v := range permissions {
		vsm[i] = v.Code
	}

	return vsm
}

func (service *UserService) GetUser(userId int) (*models.User, error) {
	user, err := service.userStore.GetUser(userId)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		permissions, err := service.permissionStore.GetPermissions(user.RoleCode)
		if err != nil {
			return nil, err
		}

		user.Permissions = transformPermissions(permissions)
		return &user, nil
	default:
		return nil, err
	}
}

func (service *UserService) GetUsers() ([]models.User, error) {
	user, err := service.userStore.GetAllUsers()
	return user, err
}

func (service *UserService) AddUser(user *models.User) (int, error) {
	hashedPassword, err := generateHashedPassword(user.Password)
	if err != nil {
		return 0, err
	}

	user.Password = hashedPassword
	user.IsActive = true
	user.IsEmailVerified = true
	return service.userStore.AddUser(user)
}

func (service *UserService) Login(emailAddress string, password string) (models.User, error) {
	user, err := service.userStore.GetUserByEmailAddress(emailAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, InvalidLogin
		}
		return user, err
	}

	err = comparePassword(password, user.Password)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return user, InvalidLogin
		}

		return user, err
	}

	permissions, err := service.permissionStore.GetPermissions(user.RoleCode)
	if err != nil {
		return user, err
	}

	user.Permissions = transformPermissions(permissions)
	return user, nil
}

func NewUserService(userStore stores.IUserStore, permissionStore stores.PermissionStore) *UserService {
	return &UserService{
		userStore:       userStore,
		permissionStore: permissionStore,
	}
}
