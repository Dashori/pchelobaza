package servicesImplementation

import (
	"backend/internal/models"
	repoErrors "backend/internal/pkg/errors/repo_errors"
	serviceErrors "backend/internal/pkg/errors/services_errors"
	"backend/internal/repository"
	"backend/internal/services"
	"fmt"
	"github.com/charmbracelet/log"
)

type FarmImplementation struct {
	FarmRepository  repository.FarmRepository
	UserRepository  repository.UserRepository
	HoneyRepository repository.HoneyRepository
	logger          *log.Logger
}

func NewFarmImplementation(
	FarmRepository repository.FarmRepository,
	UserRepository repository.UserRepository,
	HoneyRepository repository.HoneyRepository,
	logger *log.Logger,
) services.FarmService {
	return &FarmImplementation{
		FarmRepository:  FarmRepository,
		UserRepository:  UserRepository,
		HoneyRepository: HoneyRepository,
		logger:          logger,
	}
}

func (f *FarmImplementation) GetUserByLogin(userLogin string) (*models.User, error) {
	user, err := f.UserRepository.GetUserByLogin(userLogin)

	if err == repoErrors.EntityDoesNotExists {
		f.logger.Warn("REQUEST! Error, user with this login does not exists", "login", userLogin, "error", err)
		return nil, serviceErrors.UserDoesNotExists
	} else if err != nil {
		f.logger.Warn("REQUEST! Error in repository method GetUserByLogin", "login", userLogin, "error", err)
		return nil, serviceErrors.ErrorGetUserByLogin
	}

	return user, nil
}

func (f *FarmImplementation) GetUserById(id uint64) (*models.User, error) {
	f.logger.Debug("USER! Start GetUserById with", "id", id)
	tempUser, err := f.UserRepository.GetUserById(id)

	if err != nil && err == repoErrors.EntityDoesNotExists {
		f.logger.Warn("USER! Error, user with this login does not exists", "id", id, "error", err)
		return nil, serviceErrors.UserDoesNotExists
	} else if err != nil {
		f.logger.Warn("USER! Error in repository method GetUserById", "id", id, "error", err)
		return nil, serviceErrors.ErrorGetUserById
	}

	f.logger.Debug("USER! Successfully GetUserById with", "id", id)

	return tempUser, nil
}

func (f *FarmImplementation) CreateFarm(newFarm *models.Farm) (*models.Farm, error) {
	f.logger.Debug("FARM! Start create farm")
	user, err := f.GetUserById(newFarm.UserId)
	if err != nil {
		return nil, err
	}

	_, err = f.GetFarm(newFarm.Name)
	if err != serviceErrors.FarmDoesNotExists {
		f.logger.Warn("FARM! Farm with this name already exists", "name", newFarm.Name)
		return nil, serviceErrors.FarmAlreadyExists
	}

	newFarm.UserId = user.UserId

	err = f.FarmRepository.CreateFarm(newFarm)
	if err != nil {
		f.logger.Warn("FARM! Error in repository method CreateFarm", "err", err)
		return nil, serviceErrors.ErrorCreateFarm
	}

	farm, err := f.GetFarm(newFarm.Name)
	if err != nil {
		return nil, err
	}
	farm.UserLogin = user.Login

	f.logger.Info("FARM! Successfully create new farm")
	return farm, nil
}

func (f *FarmImplementation) GetFarm(name string) (*models.Farm, error) {
	f.logger.Debug("FARM! Start GetFarm with", "name", name)
	farm, err := f.FarmRepository.GetFarmByName(name)

	if err != nil && err == repoErrors.EntityDoesNotExists {
		f.logger.Warn("FARM! Error, farm with this name does not exists", "name", name, "error", err)
		return nil, serviceErrors.FarmDoesNotExists
	} else if err != nil {
		f.logger.Warn("FARM! Error in repository method GetFarmByName", "name", name, "error", err)
		return nil, serviceErrors.ErrorGetFarmByName
	}

	f.logger.Debug("FARM! Successfully GetFarm with", "name", name)

	return farm, nil
}

func (f *FarmImplementation) GetUsersFarm(login string, limit int, skipped int) ([]models.Farm, error) {
	f.logger.Debug("FARM! Start GetUsersFarm")
	if limit < 0 || skipped < 0 {
		return nil, serviceErrors.ErrorPaginationParams
	}

	user, err := f.GetUserByLogin(login)
	if err != nil {
		return nil, err
	}

	farms, err := f.FarmRepository.GetUsersFarm(user.UserId, limit, skipped)

	if err != nil && err == repoErrors.EntityDoesNotExists {
		f.logger.Warn("FARM! There is no farm with owner", "login", login, "error", err)
		return nil, nil
	} else if err != nil {
		f.logger.Warn("FARM! Error in repository method GetUsersFarm", "login", login, "error", err)
		return nil, serviceErrors.ErrorGetUsersFarm
	}

	f.logger.Info("FARM! Successfully get users farms ", "login", user.Login)

	return farms, nil
}

func (f *FarmImplementation) PatchFarm(newFarm *models.Farm) error {
	f.logger.Debug("FARM! Start PatchFarm")
	user, err := f.GetUserById(newFarm.UserId)
	if err != nil {
		return err
	}

	fmt.Println(newFarm.Name, newFarm.FarmId)
	// проверяем, что такое название еще не занято
	oldFarm, err := f.FarmRepository.GetFarmByName(newFarm.Name)

	if err != nil && err != repoErrors.EntityDoesNotExists {
		f.logger.Warn("FARM! Error in repository method GetFarmByName", "name", newFarm.Name, "error", err)
		return serviceErrors.ErrorGetFarmByName
	} else if err == nil && oldFarm.FarmId != newFarm.FarmId {
		f.logger.Warn("FARM! Farm with this name already exists with", "name", newFarm.Name)
		return serviceErrors.FarmAlreadyExists
	}

	farm, err := f.FarmRepository.GetFarmById(newFarm.FarmId)
	if err != nil && err == repoErrors.EntityDoesNotExists {
		f.logger.Warn("FARM! Error, farm with this name does not exists", "id", newFarm.FarmId, "error", err)
		return serviceErrors.FarmDoesNotExists
	} else if err != nil {
		f.logger.Warn("FARM! Error in repository method GetFarmById", "id", newFarm.FarmId, "error", err)
		return serviceErrors.ErrorGetFarmByName
	}

	if farm.UserId != user.UserId {
		f.logger.Warn("FARM! Error patch farm", "login", user.Login, "farm", farm.Name)
		return serviceErrors.ErrorFarmAccess
	}

	err = f.FarmRepository.PatchFarm(newFarm)
	if err != nil {
		f.logger.Warn("FARM! Error, farm update with login", "login", user.Login, "farm", farm.Name, "error", err)
		return serviceErrors.ErrorFarmUpdate
	}

	return nil
}
