package servicesImplementation

import (
	"backend/internal/models"
	repoErrors "backend/internal/pkg/errors/repo_errors"
	serviceErrors "backend/internal/pkg/errors/services_errors"
	"backend/internal/pkg/hasher"
	"backend/internal/repository"
	"backend/internal/services"
	"github.com/charmbracelet/log"
	// "fmt"
)

type UserImplementation struct {
	UserRepository repository.UserRepository
	hasher           hasher.Hasher
	logger           *log.Logger
}

func NewUserImplementation(
	UserRepository repository.UserRepository,
	hasher hasher.Hasher,
	logger *log.Logger,
) services.UserService {
	return &UserImplementation{
		UserRepository: UserRepository,
		hasher:           hasher,
		logger:           logger,
	}
}

func (c *UserImplementation) Create(newUser *models.User) (*models.User, error) {
	c.logger.Debug("USER! Start create user with", "login", newUser.Login)

	_, err := c.UserRepository.GetUserByLogin(newUser.Login)

	if err != nil && err != repoErrors.EntityDoesNotExists {
		c.logger.Warn("USER! Error in repository GetUserByLogin", "login", newUser.Login, "error", err)
		return nil, err
	} else if err == nil {
		c.logger.Warn("USER! User already exists", "login", newUser.Login)
		return nil, serviceErrors.UserAlreadyExists
	}

	passwordHash, err := c.hasher.GetHash(newUser.Password)
	if err != nil {
		c.logger.Warn("USER! Error get hash for password", "login", newUser.Login)
		return nil, serviceErrors.ErrorHash
	}

	newUser.Password = string(passwordHash)

	err = c.UserRepository.Create(newUser)
	if err != nil {
		c.logger.Warn("USER! Error in repository Create", "login", newUser.Login, "error", err)
		return nil, err
	}

	newUser, err = c.UserRepository.GetUserByLogin(newUser.Login)
	if err != nil {
		c.logger.Warn("USER! Error in repository method GetUserByLogin", "login", newUser.Login, "error", err)
		return nil, err
	}

	c.logger.Info("USER! Successfully create newUser", "login", newUser.Login, "id", newUser.UserId)

	return newUser, nil
}

func (c *UserImplementation) Login(login, password string) (*models.User, error) {
	c.logger.Debug("USER! Start login with", "login", login)
	tempUser, err := c.UserRepository.GetUserByLogin(login)

	if err != nil && err == repoErrors.EntityDoesNotExists {
		c.logger.Warn("USER! Error, user with this login does not exists", "login", login, "error", err)
		return nil, serviceErrors.UserDoesNotExists
	} else if err != nil {
		c.logger.Warn("USER! Error in repository method GetUserByLogin", "login", login, "error", err)
		return nil, err
	}

	if !c.hasher.CheckUnhashedValue(tempUser.Password, password) {
		c.logger.Warn("USER! Error user password", "login", login)
		return nil, serviceErrors.InvalidPassword
	}

	c.logger.Info("USER! Success login with", "login", login, "id", tempUser.UserId)

	return tempUser, nil
}

func (c *UserImplementation) GetUserByLogin(login string) (*models.User, error) {
	c.logger.Debug("USER! Start GetUserByLogin with", "login", login)
	tempUser, err := c.UserRepository.GetUserByLogin(login)

	if err != nil && err == repoErrors.EntityDoesNotExists {
		c.logger.Warn("USER! Error, user with this login does not exists", "login", login, "error", err)
		return nil, serviceErrors.UserDoesNotExists
	} else if err != nil {
		c.logger.Warn("USER! Error in repository method GetUserByLogin", "login", login, "error", err)
		return nil, err
	}

	return tempUser, nil
}

func (c *UserImplementation) Update(user *models.UserPatch) (error) {
	_, err := c.UserRepository.GetUserByLogin(user.Login)

	if err != nil && err == repoErrors.EntityDoesNotExists {
		c.logger.Warn("USER! Error, user with this login does not exists", "login", user.Login, "error", err)
		return serviceErrors.UserDoesNotExists
	} else if err != nil {
		c.logger.Warn("USER! Error in repository method GetUserByLogin", "login", user.Login, "error", err)
		return err
	}

	err = c.UserRepository.UpdateUser(user)
	if err != nil {
		c.logger.Warn("USER! Error, user update with login", "login", user.Login, "error", err)
		return serviceErrors.ErrorUserUpdate
	}

	return nil
}