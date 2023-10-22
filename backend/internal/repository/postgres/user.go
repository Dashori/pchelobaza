package postgres

import (
	"backend/internal/models"
	dbErrors "backend/internal/pkg/errors/db_errors"
	repoErrors "backend/internal/pkg/errors/repo_errors"
	"backend/internal/repository"
	"backend/internal/repository/postgres/postgres_models"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type UserPostgresRepository struct {
	db *sqlx.DB
}

func CreateUserPostgresRepository(db *sql.DB) repository.UserRepository {
	dbx := sqlx.NewDb(db, "pgx")

	return &UserPostgresRepository{db: dbx}
}

func copyUser(u postgresModel.UserPostgres) models.User {
	user := models.User{UserId: u.UserId,
		Login:    u.Login,
		Password: u.Password,
		Name:     u.Name,
		Surname:  u.Surname,
		Contact:  u.Contact,
		Role:     u.Role,

		RegisteredAt: time.Date(
			u.RegisteredAt.Year(),
			u.RegisteredAt.Month(),
			u.RegisteredAt.Day(),
			u.RegisteredAt.Hour(),
			u.RegisteredAt.Minute(),
			u.RegisteredAt.Second(),
			u.RegisteredAt.Nanosecond(),
			time.UTC),
	}

	return user
}

func (u *UserPostgresRepository) Create(user *models.User) error {
	query := `insert into bee_user(login, password, name, surname, contact, registered_at, role) values($1, $2, $3, $4, $5, $6, $7);`

	_, err := u.db.Exec(query, user.Login, user.Password, user.Name, user.Surname, user.Contact, user.RegisteredAt, user.Role)

	if err != nil {
		return dbErrors.ErrorInsert
	}

	return nil
}

func (u *UserPostgresRepository) GetUserByLogin(login string) (*models.User, error) {
	query := `select * from bee_user where login = $1;`
	userDB := &postgresModel.UserPostgres{}

	err := u.db.Get(userDB, query, login)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, dbErrors.ErrorSelect
	}

	if err != nil {
		fmt.Println(err)
	}

	userModel := copyUser(*userDB)

	return &userModel, nil
}

func (u *UserPostgresRepository) UpdateUser(user *models.UserPatch) error {
	query := `update bee_user set password = $1, name = $2, surname = $3, contact = $4 where login = $5;`

	_, err := u.db.Exec(query, user.Password, user.Name, user.Surname, user.Contact, user.Login)

	if err != nil {
		return dbErrors.ErrorUpdate
	}

	return nil
}
