package servicesImplementation

import (
	"backend/internal/models"
	repoErrors "backend/internal/pkg/errors/repo_errors"
	serviceErrors "backend/internal/pkg/errors/services_errors"
	"backend/internal/repository"
	"backend/internal/services"
	"github.com/charmbracelet/log"
)

type RequestImplementation struct {
	RequestRepository repository.RequestRepository
	UserRepository    repository.UserRepository
	logger            *log.Logger
}

func NewRequestImplementation(
	RequestRepository repository.RequestRepository,
	UserRepository repository.UserRepository,
	logger *log.Logger,
) services.RequestService {
	return &RequestImplementation{
		RequestRepository: RequestRepository,
		UserRepository:    UserRepository,
		logger:            logger,
	}
}

func (r *RequestImplementation) GetAllRequests() ([]models.Request, error) {
	r.logger.Debug("REQUEST! Start get all requests")
	requests, err := r.RequestRepository.GetAllRequests()

	if err == repoErrors.EntityDoesNotExists {
		r.logger.Warn("REQUEST! No user requests in db")
		return nil, nil
	} else if err != nil {
		r.logger.Warn("REQUEST! Error get all requests,", "error", err)
		return nil, serviceErrors.ErrorGetAllRequests
	}

	r.logger.Info("USER! Successfully get all requests")

	return requests, nil
}

func (r *RequestImplementation) GetRequestsPagination(limit int, skipped int) ([]models.Request, error) {
	r.logger.Debug("REQUEST! Start get all requests with pagination")
	if limit < 0 || skipped < 0 {
		return nil, serviceErrors.ErrorPaginationParams
	}

	requests, err := r.RequestRepository.GetRequestsPagination(limit, skipped)
	if err == repoErrors.EntityDoesNotExists {
		r.logger.Warn("REQUEST! No user requests with this pagination in db")
		return nil, nil
	} else if err != nil {
		r.logger.Warn("REQUEST! Error get requests pagination,", "error", err)
		return nil, serviceErrors.ErrorGetRequestsPagination
	}

	r.logger.Info("USER! Successfully get requests with pagination")

	return requests, nil
}

func (r *RequestImplementation) GetUserRequest(UserLogin string) (*models.Request, error) {
	r.logger.Debug("REQUEST! Start get user request")
	_, err := r.UserRepository.GetUserByLogin(UserLogin)

	if err == repoErrors.EntityDoesNotExists {
		r.logger.Warn("REQUEST! Error, user with this login does not exists", "login", UserLogin, "error", err)
		return nil, serviceErrors.UserDoesNotExists
	} else if err != nil {
		r.logger.Warn("REQUEST! Error in repository method GetUserByLogin", "login", UserLogin, "error", err)
		return nil, serviceErrors.ErrorGetUserByLogin
	}

	request, err := r.RequestRepository.GetUserRequest(UserLogin)
	if err == repoErrors.EntityDoesNotExists {
		r.logger.Warn("REQUEST! Error, request with this login does not exists", "login", UserLogin, "error", err)
		return nil, serviceErrors.RequestDoesNotExists
	} else if err != nil {
		r.logger.Warn("REQUEST! Error get user requests,", "error", err)
		return nil, serviceErrors.ErrorGetUserRequest
	}

	r.logger.Info("USER! Successfully get user request")

	return request, nil
}

func (r *RequestImplementation) PatchUserRequest(request models.Request) error {
	r.logger.Debug("REQUEST! Start patch user request")
	oldRequest, err := r.GetUserRequest(request.UserLogin)
	if err != nil {
		return err
	}

	if oldRequest.Status != "waiting" {
		r.logger.Warn("REQUEST! Can't patch request without status waiting")
		return serviceErrors.ErrorRequestStatus // нельзя его редактировать
	}

	err = r.RequestRepository.PatchUserRequest(&request)
	if err != nil {
		r.logger.Warn("REQUEST! Error patch user request", "error", err)
		return serviceErrors.ErrorRequestPatch
	}

	r.logger.Info("USER! Successfully patch user request")

	return nil
}

func (r *RequestImplementation) CreateRequest(newRequest *models.Request) (*models.Request, error) {
	r.logger.Debug("REQUEST! Start create user request")
	UserLogin := newRequest.UserLogin

	// проверка что такое пользователь существует
	user, err := r.UserRepository.GetUserByLogin(UserLogin)
	if err == repoErrors.EntityDoesNotExists {
		r.logger.Warn("REQUEST! Error, user with this login does not exists", "login", UserLogin, "error", err)
		return nil, serviceErrors.UserDoesNotExists
	} else if err != nil {
		r.logger.Warn("REQUEST! Error in repository method GetUserByLogin", "login", UserLogin, "error", err)
		return nil, serviceErrors.ErrorGetUserByLogin
	}

	// проверка что пользователь не beemaster
	if user.Role == "beemaster" {
		r.logger.Warn("REQUEST! Error, user already beemaster", "login", UserLogin)
		return nil, serviceErrors.UserAlreadyBeemaster
	}

	// проверка что заявки еще нет
	request, err := r.RequestRepository.GetUserRequest(UserLogin)
	if err == nil && err == repoErrors.EntityDoesNotExists {
		r.logger.Warn("REQUEST! Request for this user already exists", "login", newRequest.UserLogin)
		return nil, serviceErrors.RequestAlreadyExists
	}

	newRequest.UserId = user.UserId
	newRequest.Status = "waiting"
	err = r.RequestRepository.Create(newRequest)
	if err != nil {
		r.logger.Warn("REQUEST! Error create user request", "error", err)
		return nil, serviceErrors.ErrorCreateRequest
	}

	request, err = r.GetUserRequest(newRequest.UserLogin)
	if err != nil {
		return nil, err
	}

	r.logger.Info("USER! Successfully create user request")

	return request, nil
}
