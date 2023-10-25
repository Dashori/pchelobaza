package postgres

import (
	"backend/internal/models"
	repoErrors "backend/internal/pkg/errors/repo_errors"
	"backend/internal/repository"
	"backend/internal/repository/postgres/postgres_models"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"time"
)

type ConferencePostgresRepository struct {
	db *sqlx.DB
}

func CreateConferencePostgresRepository(db *sql.DB) repository.ConferenceRepository {
	dbx := sqlx.NewDb(db, "pgx")

	return &ConferencePostgresRepository{db: dbx}
}

func copyConference(c postgresModel.ConferencePostgres) models.Conference {
	conference := models.Conference{
		ConferenceId: c.ConferenceId,
		UserId:       c.UserId,
		Name:         c.Name,
		Description:  c.Description,
		Address:      c.Address,
		MaxUsers:     c.MaxUsers,
		CurrentUsers: c.CurrentUsers,
		Date: time.Date(
			c.Date.Year(),
			c.Date.Month(),
			c.Date.Day(),
			c.Date.Hour(),
			c.Date.Minute(),
			c.Date.Second(),
			c.Date.Nanosecond(),
			time.UTC),
	}

	return conference
}

func copyReview(r postgresModel.ReviewPostgres) models.Review {
	review := models.Review{
		ConferenceId: r.ConferenceId,
		UserId:       r.UserId,
		Description:  r.Description,
		Login:        r.Login,
		Name:         r.Name,
		Surname:      r.Surname,
		Date: time.Date(
			r.Date.Year(),
			r.Date.Month(),
			r.Date.Day(),
			r.Date.Hour(),
			r.Date.Minute(),
			r.Date.Second(),
			r.Date.Nanosecond(),
			time.UTC),
	}

	return review
}

func (c *ConferencePostgresRepository) GetAllConferences(limit int, skipped int) ([]models.Conference, error) {
	query := `select * from bee_conference
	order by date desc
	offset $1
	limit $2;`

	var conferencePostgres []postgresModel.ConferencePostgres
	err := c.db.Select(&conferencePostgres, query, skipped, limit)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	conferenceModels := []models.Conference{}

	for _, r := range conferencePostgres {
		conference := copyConference(r)
		conferenceModels = append(conferenceModels, conference)
	}

	return conferenceModels, nil
}

func (c *ConferencePostgresRepository) CreateConference(conference *models.Conference) error {
	query := `insert into bee_conference(id_user, name, description, date, address, 
		maximum_users, current_users) values($1, $2, $3, $4, $5, $6, $7);`

	_, err := c.db.Exec(query, conference.UserId, conference.Name, conference.Description,
		conference.Date, conference.Address, conference.MaxUsers, conference.CurrentUsers)

	if err != nil {
		return err
	}

	return nil
}

func (c *ConferencePostgresRepository) GetConferenceByName(name string) (*models.Conference, error) {
	query := `select * from bee_conference where name = $1;`
	conferenceDB := &postgresModel.ConferencePostgres{}

	err := c.db.Get(conferenceDB, query, name)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	conferenceModel := copyConference(*conferenceDB)

	return &conferenceModel, nil
}

func (c *ConferencePostgresRepository) PatchConference(conference *models.Conference) error {
	query := `update bee_conference set description = $1, date = $2, address = $3,
	 maximum_users = $4  where id = $5;`

	_, err := c.db.Exec(query, conference.Description, conference.Date, conference.Address,
		conference.MaxUsers, conference.ConferenceId)

	if err != nil {
		return err
	}

	return nil
}

func (c *ConferencePostgresRepository) GetAllConferenceUsers(name string) ([]models.User, error) {
	query := `select u.login, u.name, u.surname, cn.name 
	from bee_user_conference as c
	join bee_user as u on c.id_user = u.id
	join bee_conference as cn on c.id_conference = cn.id
	where cn.name = $1;`

	var usersPostgres []postgresModel.UserPostgres
	err := c.db.Select(&usersPostgres, query, name)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	conferenceUsers := []models.User{}

	for _, r := range usersPostgres {
		user := copyUser(r)
		conferenceUsers = append(conferenceUsers, user)
	}

	return conferenceUsers, nil
}

func (c *ConferencePostgresRepository) GetConferenceUsers(name string, limit int,
	skipped int) ([]models.User, error) {
	query := `select u.login, u.name, u.surname, cn.name 
	from bee_user_conference as c
	join bee_user as u on c.id_user = u.id
	join bee_conference as cn on c.id_conference = cn.id
	where cn.name = $1
	offset $2
	limit $3;`

	var usersPostgres []postgresModel.UserPostgres
	err := c.db.Select(&usersPostgres, query, name, skipped, limit)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	conferenceUsers := []models.User{}

	for _, r := range usersPostgres {
		user := copyUser(r)
		conferenceUsers = append(conferenceUsers, user)
	}

	return conferenceUsers, nil
}

func (c *ConferencePostgresRepository) PatchConferenceUsers(conference *models.Conference, UserId uint64) error {
	query := `update bee_conference set current_users = $1 where id = $2;`

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, conference.CurrentUsers, conference.ConferenceId)

	if err != nil {
		tx.Rollback()
		return err
	}

	query = `insert into bee_user_conference(id_user, id_conference) values($1, $2);`

	_, err = tx.Exec(query, UserId, conference.ConferenceId)

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (c *ConferencePostgresRepository) GetConferenceReviews(name string, limit int,
	skipped int) ([]models.Review, error) {
	query := `select r.description, u.login, u.name, u.surname
	from bee_review as r 
	join bee_user as u on u.id= r.id_user
	join bee_conference as cn on r.id_conference = cn.id
	where cn.name = $1
	offset $2
	limit $3;`

	var reviewsPostgres []postgresModel.ReviewPostgres
	err := c.db.Select(&reviewsPostgres, query, name, skipped, limit)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	reviewsModels := []models.Review{}

	for _, r := range reviewsPostgres {
		review := copyReview(r)
		reviewsModels = append(reviewsModels, review)
	}

	return reviewsModels, nil
}

func (c *ConferencePostgresRepository) CreateReview(review *models.Review) error {
	query := `insert into bee_review(id_conference, id_user, date, description) values($1, $2, $3, $4);`

	_, err := c.db.Exec(query, review.ConferenceId, review.UserId, review.Date, review.Description)

	if err != nil {
		return err
	}

	return nil
}
