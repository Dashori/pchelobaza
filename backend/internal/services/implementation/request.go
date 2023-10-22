package servicesImplementation

import (
	"backend/internal/models"
	"backend/internal/repository"
	"backend/internal/services"
	"github.com/charmbracelet/log"
	repoErrors "backend/internal/pkg/errors/repo_errors"
	serviceErrors "backend/internal/pkg/errors/services_errors"
)

type RequestImplementation struct {
	RequestRepository repository.RequestRepository
	UserRepository repository.UserRepository
	logger          *log.Logger
}

func NewRequestImplementation(
	RequestRepository repository.RequestRepository,
	UserRepository repository.UserRepository,
	logger *log.Logger,
) services.RequestService {
	return &RequestImplementation{
		RequestRepository: RequestRepository,
		UserRepository: UserRepository,
		logger:          logger,
	}
}

func (r *RequestImplementation) GetAllRequests() ([]models.Request, error) {
	r.logger.Debug("REQUEST! Start get all requests")
	requests, err := r.RequestRepository.GetAllRequests()
	if err != nil {
		r.logger.Warn("REQUEST! Error get all requests,", "error", err)
		return nil, err
	}

	return requests, nil
}

func (r *RequestImplementation) GetRequestsPagination(limit int, skipped int) ([]models.Request, error) {
	r.logger.Debug("REQUEST! Start get all requests")
	requests, err := r.RequestRepository.GetRequestsPagination(limit, skipped)
	if err != nil {
		r.logger.Warn("REQUEST! Error get all requests,", "error", err)
		return nil, err
	}

	return requests, nil
}

func (r *RequestImplementation) GetUserRequest(UserLogin string) (*models.Request, error) {
	r.logger.Debug("REQUEST! Start get user request")
	_, err := r.UserRepository.GetUserByLogin(UserLogin)

	if err != nil && err == repoErrors.EntityDoesNotExists {
		r.logger.Warn("REQUEST! Error, user with this login does not exists", "login", UserLogin, "error", err)
		return nil, serviceErrors.UserDoesNotExists
	} else if err != nil {
		r.logger.Warn("REQUEST! Error in repository method GetUserByLogin", "login", UserLogin, "error", err)
		return nil, err
	}

	request, err := r.RequestRepository.GetUserRequest(UserLogin)
	if err != nil {
		r.logger.Warn("REQUEST! Error get user requests,", "error", err)
		return nil, err
	}

	return request, nil
}

func (r *RequestImplementation) PatchUserRequest(request models.Request) (error) {
	r.logger.Debug("REQUEST! Start patch user request")
	oldRequest, err := r.GetUserRequest(request.UserLogin)
	if err != nil {
		return err
	}
	if oldRequest.Status != "waiting" {
		return nil // нельзя его редактировать
	}
		
	err = r.RequestRepository.PatchUserRequest(&request)
	if err != nil {
		r.logger.Warn("REQUEST! Error patch user request", "error", err)
		return err
	}
	r.logger.Debug("REQUEST! PATCH IS OK")

	return nil
}