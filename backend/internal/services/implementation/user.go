package servicesImplementation

import (
	"backend/internal/models"
	repoErrors "backend/internal/pkg/errors/repo_errors"
	serviceErrors "backend/internal/pkg/errors/services_errors"
	"backend/internal/pkg/hasher"
	"backend/internal/repository"
	"backend/internal/services"
	"github.com/charmbracelet/log"
	"fmt"
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
	fmt.Println("AAAAAAAA")
	_, err := c.UserRepository.GetUserByLogin(newUser.Login)

	if err != nil && err != repoErrors.EntityDoesNotExists {
		c.logger.Warn("USER! Error in repository GetUserByLogin", "login", newUser.Login, "error", err)
		return nil, err
	} else if err == nil {
		c.logger.Warn("USER! User already exists", "login", newUser.Login)
		return nil, serviceErrors.UserAlreadyExists
	}
	fmt.Println("BBBBBB")

	// passwordHash, err := c.hasher.GetHash(newUser.Password)
	// if err != nil {
	// 	c.logger.Warn("USER! Error get hash for password", "login", newUser.Login)
	// 	return nil, serviceErrors.ErrorHash
	// }

	// newUser.Password = string(passwordHash)
	
	// err = c.UserRepository.Create(newUser)
	// if err != nil {
	// 	c.logger.Warn("USER! Error in repository Create", "login", newUser.Login, "error", err)
	// 	return nil, err
	// }

	// newUser, err = c.GetUserByLogin(newUser.Login)
	// if err != nil {
	// 	c.logger.Warn("USER! Error in repository method GetUserByLogin", "login", newUser.Login, "error", err)
	// 	return nil, err
	// }

	// c.logger.Info("USER! Successfully create newUser", "login", newUser.Login, "id", newUser.UserId)

	return newUser, nil
}

func (c *UserImplementation) Login(login, password string) (*models.User, error) {
	var newUser *models.User

	return newUser, nil
}

func (c *UserImplementation)  GetUserByLogin(login string) (*models.User, error) {
	var newUser *models.User

	return newUser, nil
}