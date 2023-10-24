package servicesImplementation

import (
	"backend/internal/models"
	repoErrors "backend/internal/pkg/errors/repo_errors"
	serviceErrors "backend/internal/pkg/errors/services_errors"
	"backend/internal/pkg/hasher"
	"backend/internal/repository"
	"backend/internal/services"
	"github.com/charmbracelet/log"
	"time"
)

type UserImplementation struct {
	UserRepository repository.UserRepository
	hasher         hasher.Hasher
	logger         *log.Logger
}

func NewUserImplementation(
	UserRepository repository.UserRepository,
	hasher hasher.Hasher,
	logger *log.Logger,
) services.UserService {
	return &UserImplementation{
		UserRepository: UserRepository,
		hasher:         hasher,
		logger:         logger,
	}
}

func (u *UserImplementation) GetUserByLogin(login string) (*models.User, error) {
	u.logger.Debug("USER! Start GetUserByLogin with", "login", login)
	tempUser, err := u.UserRepository.GetUserByLogin(login)

	if err != nil && err == repoErrors.EntityDoesNotExists {
		u.logger.Warn("USER! Error, user with this login does not exists", "login", login, "error", err)
		return nil, serviceErrors.UserDoesNotExists
	} else if err != nil {
		u.logger.Warn("USER! Error in repository method GetUserByLogin", "login", login, "error", err)
		return nil, serviceErrors.ErrorGetUserByLogin
	}

	u.logger.Debug("USER! Successfully GetUserByLogin with", "login", login)

	return tempUser, nil
}

func (u *UserImplementation) Create(newUser *models.User) (*models.User, error) {
	u.logger.Debug("USER! Start create user with", "login", newUser.Login)

	// проверка совпадают ли пароли при регистрации
	if newUser.Password != newUser.ConfirmPassword {
		return nil, serviceErrors.ErrorConfirmPassword
	}

	// проверка что такого юзера нет
	_, err := u.UserRepository.GetUserByLogin(newUser.Login)

	if err != nil && err != repoErrors.EntityDoesNotExists {
		u.logger.Warn("USER! Error in repository method GetUserByLogin", "login", newUser.Login, "error", err)
		return nil, serviceErrors.ErrorUserCreate
	} else if err == nil {
		u.logger.Warn("USER! User already exists with", "login", newUser.Login)
		return nil, serviceErrors.UserAlreadyExists
	}

	// хэшируем пароль
	passwordHash, err := u.hasher.GetHash(newUser.Password)
	if err != nil {
		u.logger.Warn("USER! Error get hash for password", "login", newUser.Login)
		return nil, serviceErrors.ErrorHash
	}
	newUser.Role = "beeman"
	newUser.Password = string(passwordHash)
	newUser.RegisteredAt = time.Now()

	err = u.UserRepository.Create(newUser)
	if err != nil {
		u.logger.Warn("USER! Error in repository Create", "login", newUser.Login, "error", err)
		return nil, serviceErrors.ErrorUserCreate
	}

	newUser, err = u.GetUserByLogin(newUser.Login)
	if err != nil {
		return nil, err
	}

	u.logger.Info("USER! Successfully create newUser", "login", newUser.Login, "id", newUser.UserId)

	return newUser, nil
}

func (u *UserImplementation) Login(login, password string) (*models.User, error) {
	u.logger.Debug("USER! Start login with", "login", login)
	tempUser, err := u.GetUserByLogin(login)
	if err != nil {
		return nil, err
	}

	if !u.hasher.CheckUnhashedValue(tempUser.Password, password) {
		u.logger.Warn("USER! Error user password", "login", login)
		return nil, serviceErrors.InvalidPassword
	}

	u.logger.Info("USER! Success login with", "login", login, "id", tempUser.UserId)

	return tempUser, nil
}

func (u *UserImplementation) Update(user *models.User) error {
	u.logger.Debug("USER! Start update user with", "login", user.Login)
	_, err := u.GetUserByLogin(user.Login)
	if err != nil {
		return err
	}

	passwordHash, err := u.hasher.GetHash(user.Password)
	if err != nil {
		u.logger.Warn("USER! Error get hash for password", "login", user.Login)
		return serviceErrors.ErrorHash
	}

	user.Password = string(passwordHash)

	err = u.UserRepository.UpdateUser(user)
	if err != nil {
		u.logger.Warn("USER! Error, user update with login", "login", user.Login, "error", err)
		return serviceErrors.ErrorUserUpdate
	}

	u.logger.Info("USER! Successfully update user with", "login", user.Login)

	return nil
}
