package servicesImplementation

import (
	"backend/internal/models"
	"backend/internal/repository"
	"backend/internal/services"
	"github.com/charmbracelet/log"
)

type HoneyImplementation struct {
	HoneyRepository repository.HoneyRepository
	logger          *log.Logger
}

func NewHoneyImplementation(
	HoneyRepository repository.HoneyRepository,
	logger *log.Logger,
) services.HoneyService {
	return &HoneyImplementation{
		HoneyRepository: HoneyRepository,
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
