package UserImplementation

import (
	"backend/internal/models"
	// "backend/internal/pkg/errors/repoErrors"
	// "backend/internal/pkg/errors/servicesErrors"
	"backend/internal/pkg/hasher"
	"backend/internal/repository"
	"backend/internal/services"
	// "github.com/charmbracelet/log"
	"fmt"
)

type UserImplementation struct {
	UserRepository repository.UserRepository
	hasher           hasher.Hasher
	// logger           *log.Logger
}

func NewUserImplementation(
	UserRepository repository.UserRepository,
	hasher hasher.Hasher,
	// logger *log.Logger,
) services.UserService {

	fmt.Println("HERE!")
	return &UserImplementation{
		UserRepository: UserRepository,
		hasher:           hasher,
		// logger:           logger,
	}
}

func (c *UserImplementation) Create(user *models.NewUser) (*models.User, error) {
	// c.logger.Debug("CLIENT! Start create client with", "login", client.Login)

	// _, err := c.ClientRepository.GetClientByLogin(client.Login)

	// if err != nil && err != repoErrors.EntityDoesNotExists {
	// 	c.logger.Warn("CLIENT! Error in repository GetClientByLogin", "login", client.Login, "error", err)
	// 	return nil, err
	// } else if err == nil {
	// 	c.logger.Warn("CLIENT! Client already exists", "login", client.Login)
	// 	return nil, serviceErrors.ClientAlreadyExists
	// }

	// passwordHash, err := c.hasher.GetHash(password)
	// if err != nil {
	// 	c.logger.Warn("CLIENT! Error get hash for password", "login", client.Login)
	// 	return nil, serviceErrors.ErrorHash
	// }
	// client.Password = string(passwordHash)

	// err = c.ClientRepository.Create(client)
	// if err != nil {
	// 	c.logger.Warn("CLIENT! Error in repository Create", "login", client.Login, "error", err)
	// 	return nil, err
	// }

	// newClient, err := c.GetClientByLogin(client.Login)
	// if err != nil {
	// 	c.logger.Warn("CLIENT! Error in repository method GetClientByLogin", "login", client.Login, "error", err)
	// 	return nil, err
	// }

	// c.logger.Info("CLIENT! Successfully create client", "login", newClient.Login, "id", newClient.ClientId)
	var newClient *models.User

	return newClient, nil
}

func (c *UserImplementation) Login(login, password string) (*models.User, error) {
	var newClient *models.User

	return newClient, nil
}

func (c *UserImplementation)  GetUserByLogin(login string) (*models.User, error) {
	var newClient *models.User

	return newClient, nil
}