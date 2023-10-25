package servicesImplementation

import (
	"backend/internal/models"
	repoErrors "backend/internal/pkg/errors/repo_errors"
	serviceErrors "backend/internal/pkg/errors/services_errors"
	"backend/internal/repository"
	"backend/internal/services"
	"github.com/charmbracelet/log"
)

type HoneyImplementation struct {
	HoneyRepository repository.HoneyRepository
	FarmRepository  repository.FarmRepository
	logger          *log.Logger
}

func NewHoneyImplementation(
	HoneyRepository repository.HoneyRepository,
	FarmRepository repository.FarmRepository,
	logger *log.Logger,
) services.HoneyService {
	return &HoneyImplementation{
		HoneyRepository: HoneyRepository,
		FarmRepository:  FarmRepository,
		logger:          logger,
	}
}

func (h *HoneyImplementation) GetAllHoney() ([]models.Honey, error) {
	h.logger.Debug("HONEY! Start get all honey")
	honey, err := h.HoneyRepository.GetAllHoney()
	if err != nil {
		h.logger.Warn("HONEY! Error get all honey", "error", err)
		return nil, err
	}

	h.logger.Info("HONEY! Successfully get all honey")

	return honey, nil
}

func (h *HoneyImplementation) GetFarmHoney(name string) ([]models.Honey, error) {
	h.logger.Debug("HONEY! Start get farm honey")
	_, err := h.FarmRepository.GetFarmByName(name)

	if err != nil && err == repoErrors.EntityDoesNotExists {
		h.logger.Warn("HONEY! Error, farm with this name does not exists", "name", name, "error", err)
		return nil, serviceErrors.FarmDoesNotExists
	} else if err != nil {
		h.logger.Warn("HONEY! Error in repository method GetFarmByName", "name", name, "error", err)
		return nil, serviceErrors.ErrorGetFarmByName
	}

	honey, err := h.HoneyRepository.GetFarmHoney(name)
	if err != nil {
		h.logger.Warn("HONEY! Error get farm honey", "error", err)
		return nil, err
	}

	h.logger.Info("HONEY! Successfully get farm honey")

	return honey, nil
}
