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

func (c *HoneyImplementation) GetAllHoney() ([]models.Honey, error) {
	c.logger.Debug("HONEY! Start get all honey")
	honey, err := c.HoneyRepository.GetAllHoney()
	if err != nil {
		c.logger.Warn("HONEY! Error get all honey", "error", err)
		return nil, err
	}

	c.logger.Info("HONEY! Successfully get all honey")

	return honey, nil
}
