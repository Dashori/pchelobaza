package postgres

import (
	"backend/internal/models"
	// "backend/internal/pkg/errors/dbErrors"
	// "backend/internal/pkg/errors/repoErrors"
	// "backend/internal/repository"
	"database/sql"
	// "github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type UserPostgresRepository struct {
	db *sqlx.DB
}

// func NewUserPostgresRepository(db *sqlx.DB) repository.UserRepository {
// 	return &UserPostgresRepository{db: db}
// }

// type PostgresRepositoryFields struct {
// 	DB     *sql.DB
// 	// Config config.Config
// }

func CreateUserPostgresRepository(db *sql.DB) UserPostgresRepository {
	dbx := sqlx.NewDb(db, "pgx")

	return UserPostgresRepository{db: dbx}
}


func (c *UserPostgresRepository) Create(user *models.User) error {
	query := `insert into bee_user(login, password) values($1, $2, $3, $4, $5, $6);`

	_, err := c.db.Exec(query, user.Login, user.Password, user.Name, user.Surname, user.Contacts, user.RegistrationDate, user.Role)

	if err != nil {
		// return dbErrors.ErrorInsert
		return fmt.Errorf("aaaaa")
	}

	return nil
}