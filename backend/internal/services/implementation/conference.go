package servicesImplementation

import (
	"backend/internal/models"
	repoErrors "backend/internal/pkg/errors/repo_errors"
	serviceErrors "backend/internal/pkg/errors/services_errors"
	"backend/internal/repository"
	"backend/internal/services"
	"github.com/charmbracelet/log"
	"time"
)

type ConferenceImplementation struct {
	ConferenceRepository repository.ConferenceRepository
	UserRepository       repository.UserRepository
	logger               *log.Logger
}

func NewConferenceImplementation(
	ConferenceRepository repository.ConferenceRepository,
	UserRepository repository.UserRepository,
	logger *log.Logger,
) services.ConferenceService {
	return &ConferenceImplementation{
		ConferenceRepository: ConferenceRepository,
		UserRepository:       UserRepository,
		logger:               logger,
	}
}

func (c *ConferenceImplementation) GetUserByLogin(userLogin string) (*models.User, error) {
	user, err := c.UserRepository.GetUserByLogin(userLogin)

	if err == repoErrors.EntityDoesNotExists {
		c.logger.Warn("CONFERENCE! Error, user with this login does not exists", "login", userLogin, "error", err)
		return nil, serviceErrors.UserDoesNotExists
	} else if err != nil {
		c.logger.Warn("CONFERENCE! Error in repository method GetUserByLogin", "login", userLogin, "error", err)
		return nil, serviceErrors.ErrorGetUserByLogin
	}

	return user, nil
}

func (c *ConferenceImplementation) GetAllConferences(limit int, skipped int) ([]models.Conference, error) {
	c.logger.Debug("CONFERENCE! Start get all conference")
	if limit < 0 || skipped < 0 {
		return nil, serviceErrors.ErrorPaginationParams
	}

	conference, err := c.ConferenceRepository.GetAllConferences(limit, skipped)
	if err == repoErrors.EntityDoesNotExists {
		c.logger.Warn("CONFERENCE! No user conferences with this pagination in db")
		return nil, nil
	} else if err != nil {
		c.logger.Warn("CONFERENCE! Error get conferences pagination,", "error", err)
		return nil, serviceErrors.ErrorGetConferencesPagination
	}

	c.logger.Info("CONFERENCE! Successfully get conferences with pagination")

	return conference, nil
}

func (c *ConferenceImplementation) CreateConference(conference *models.Conference) (*models.Conference, error) {
	c.logger.Debug("CONFERENCE! Start create conference")
	user, err := c.GetUserByLogin(conference.UserLogin)
	if err != nil {
		return nil, err
	}
	// проверка что пользователь beemaster
	if user.Role != "beemaster" {
		c.logger.Warn("CONFERENCE! User is not beemaster", "login", user.Login)
		return nil, serviceErrors.ErrorRoleForConference
	}

	today := time.Now()

	if !conference.Date.After(today) {
		c.logger.Warn("CONFERENCE! Bad date for conference", "date", conference.Date)
		return nil, serviceErrors.ErrorDateForConference
	}

	if conference.MaxUsers < 10 {
		c.logger.Warn("CONFERENCE! Bad count of users", "users", conference.MaxUsers)
		return nil, serviceErrors.ErrorUsersForConference
	}

	_, err = c.ConferenceRepository.GetConferenceByName(conference.Name)

	if err != nil && err != repoErrors.EntityDoesNotExists {
		c.logger.Warn("CONFERENCE! Error in repository method GetConferenceByName", "name", conference.Name, "error", err)
		return nil, err
	} else if err == nil {
		c.logger.Warn("CONFERENCE! Conference already exists with", "name", conference.Name)
		return nil, serviceErrors.ErrorNameForConference
	}

	conference.UserId = user.UserId
	conference.CurrentUsers = 0

	err = c.ConferenceRepository.CreateConference(conference)
	if err != nil {
		c.logger.Warn("CONFERENCE! Error in repository CreateConference", "error", err)
		return nil, serviceErrors.ErrorCreateConference
	}

	newConference, err := c.ConferenceRepository.GetConferenceByName(conference.Name)
	if err != nil && err != repoErrors.EntityDoesNotExists {
		c.logger.Warn("CONFERENCE! Error in repository method GetConferenceByName", "name", conference.Name, "error", err)
		return nil, serviceErrors.ErrorGetConference
	}

	c.logger.Info("CONFERENCE! Successfully create conference", "name", conference.Name)

	return newConference, nil
}

func (c *ConferenceImplementation) GetConferenceByName(name string) (*models.Conference, error) {
	c.logger.Debug("CONFERENCE! Start get conference by name")
	conference, err := c.ConferenceRepository.GetConferenceByName(name)
	if err != nil && err != repoErrors.EntityDoesNotExists {
		c.logger.Warn("CONFERENCE! Error in repository method GetConferenceByName", "name", conference.Name, "error", err)
		return nil, serviceErrors.ErrorGetConference
	}

	c.logger.Info("CONFERENCE! Successfully get conference", "name", name)
	return conference, nil
}

func (c *ConferenceImplementation) PatchConference(conference *models.Conference) error {
	// descr, date, max users, address
	c.logger.Debug("CONFERENCE! Start patch conference")
	oldConference, err := c.ConferenceRepository.GetConferenceByName(conference.Name)
	if err != nil && err != repoErrors.EntityDoesNotExists {
		c.logger.Warn("CONFERENCE! Error in repository method GetConferenceByName", "name", conference.Name, "error", err)
		return serviceErrors.ErrorGetConference
	} else if err == repoErrors.EntityDoesNotExists {
		c.logger.Warn("CONFERENCE! There is no conference with", "name", conference.Name)
		return serviceErrors.ErrorNoConference
	}

	user, err := c.GetUserByLogin(conference.UserLogin)
	if err != nil {
		return err
	}

	if oldConference.UserId != user.UserId {
		c.logger.Warn("CONFERENCE! You can not edit this conference", "name", conference.Name)
		return serviceErrors.ErrorNoYourConference
	}

	// проверка что пользователь beemaster
	if user.Role != "beemaster" {
		c.logger.Warn("CONFERENCE! User is not beemaster", "login", user.Login)
		return serviceErrors.ErrorRoleForConference
	}

	today := time.Now()

	if !oldConference.Date.After(today) {
		c.logger.Warn("CONFERENCE! This conference has already passed", "date", oldConference.Date)
		return serviceErrors.ErrorOldConference
	}

	if !conference.Date.After(today) {
		c.logger.Warn("CONFERENCE! Bad date for conference", "date", conference.Date)
		return serviceErrors.ErrorDateForConference
	}

	if conference.MaxUsers < 10 {
		c.logger.Warn("CONFERENCE! Bad count of users", "users", conference.MaxUsers)
		return serviceErrors.ErrorUsersForConference
	}

	conference.ConferenceId = oldConference.ConferenceId

	err = c.ConferenceRepository.PatchConference(conference)
	if err != nil && err != repoErrors.EntityDoesNotExists {
		c.logger.Warn("CONFERENCE! Error in repository method PatchConference", "name", conference.Name, "error", err)
		return serviceErrors.ErrorEditConference
	}

	c.logger.Info("CONFERENCE! Successfully update conference", "name", conference.Name)
	return nil
}

func (c *ConferenceImplementation) GetAllConferenceUsers(name string) ([]models.User, error) {
	c.logger.Debug("CONFERENCE! Start get all conference users")
	users, err := c.ConferenceRepository.GetAllConferenceUsers(name)
	if err != nil && err != repoErrors.EntityDoesNotExists {
		c.logger.Warn("CONFERENCE! Error in repository method GetConferenceByName", "name", name, "error", err)
		return nil, serviceErrors.ErrorGetConferenceUsers
	}

	c.logger.Info("CONFERENCE! Successfully all conference users", "name", name)
	return users, nil
}

func (c *ConferenceImplementation) GetConferenceUsers(name string,
	limit int, skipped int) ([]models.User, error) {
	c.logger.Debug("CONFERENCE! Start get all conference users")
	users, err := c.ConferenceRepository.GetConferenceUsers(name, limit, skipped)
	if err != nil && err != repoErrors.EntityDoesNotExists {
		c.logger.Warn("CONFERENCE! Error in repository method GetConferenceByName", "name", name, "error", err)
		return nil, serviceErrors.ErrorGetConferenceUsers
	}

	c.logger.Info("CONFERENCE! Successfully all conference users", "name", name)
	return users, nil
}

func (c *ConferenceImplementation) PatchConferenceUsers(name string, login string) error {
	c.logger.Debug("CONFERENCE! Start patch conference users")
	conference, err := c.GetConferenceByName(name)
	if err != nil {
		return err
	}

	today := time.Now()

	if !conference.Date.After(today) {
		c.logger.Warn("CONFERENCE! This conference has already passed", "date", conference.Date)
		return serviceErrors.ErrorOldConference
	}

	if conference.CurrentUsers+1 > conference.MaxUsers {
		c.logger.Warn("CONFERENCE! Error no free place", "name", name)
		return serviceErrors.ErrorNoPlace
	}

	user, err := c.GetUserByLogin(login)
	if err != nil {
		return err
	}

	// надо проверить что он уже не зареган
	users, err := c.GetAllConferenceUsers(name)
	if err != nil {
		return err
	}

	for _, j := range users {
		if j.Login == user.Login {
			c.logger.Warn("CONFERENCE! User already registrated", "login", user.Login)
			return serviceErrors.ErrorConferenceJoin
		}
	}

	conference.CurrentUsers += 1

	err = c.ConferenceRepository.PatchConferenceUsers(conference, user.UserId)
	if err != nil && err != repoErrors.EntityDoesNotExists {
		c.logger.Warn("CONFERENCE! Error in repository method PatchConferenceUsers", "name", name, "error", err)
		return serviceErrors.ErrorJoinConf
	}

	c.logger.Info("CONFERENCE! Successfully patch conference", "name", name)
	return nil
}

func (c *ConferenceImplementation) GetConferenceReviews(name string,
	limit int, skipped int) ([]models.Review, error) {

	c.logger.Debug("CONFERENCE! Start get all conference reviews")
	_, err := c.GetConferenceByName(name)
	if err != nil {
		return nil, err
	}

	reviews, err := c.ConferenceRepository.GetConferenceReviews(name, limit, skipped)
	if err != nil && err != repoErrors.EntityDoesNotExists {
		c.logger.Warn("CONFERENCE! Error in repository method GetConferenceReviews", "name", name, "error", err)
		return nil, serviceErrors.ErrorGetConferenceReviews
	}

	c.logger.Info("CONFERENCE! Successfully get conference reviews", "name", name)
	return reviews, nil
}

func (c *ConferenceImplementation) CreateReview(review *models.Review) (*models.Review, error) {
	c.logger.Debug("CONFERENCE! Start create review")

	conf, err := c.GetConferenceByName(review.ConferenceName)
	if err != nil {
		return nil, err
	}
	user, err := c.GetUserByLogin(review.Login)
	if err != nil {
		return nil, err
	}

	review.UserId = user.UserId
	review.Date = time.Now()
	review.ConferenceId = conf.ConferenceId

	err = c.ConferenceRepository.CreateReview(review)
	if err != nil && err != repoErrors.EntityDoesNotExists {
		c.logger.Warn("CONFERENCE! Error in repository method CreateReview", "conf name", review.ConferenceName, "error", err)
		return nil, serviceErrors.ErrorGetConferenceReviews
	}

	c.logger.Info("CONFERENCE! Successfully create review", "conf name", review.ConferenceName)
	return review, nil
}
