package postgres

import (
	"backend/internal/models"
	// "backend/internal/pkg/errors/dbErrors"
	// "backend/internal/pkg/errors/repoErrors"
	// "backend/internal/repository/postgres/postgres_models"
	"database/sql"
	"backend/internal/repository"
	// "github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type UserPostgresRepository struct {
	db *sqlx.DB
}

func CreateUserPostgresRepository(db *sql.DB) repository.UserRepository {
	dbx := sqlx.NewDb(db, "pgx")

	return &UserPostgresRepository{db: dbx}
}


func (c *UserPostgresRepository) Create(user *models.User) error {
	// fmt.Println("HERE")
	query := `insert into bee_user(login, password, name, contact, registered_at, role) values($1, $2, $3, $4, $5, $6);`

	_, err := c.db.Exec(query, user.Login, user.Password, user.Name, user.Surname, user.Contacts, user.RegistrationDate, user.Role)

	if err != nil {
		// return dbErrors.ErrorInsert
		return fmt.Errorf("aaaaa")
		// return err
	}

	return nil
}

// func (c *UserPostgresRepository) GetUserByLogin(login string) (*models.User, error) {
// 	query := `select * from bee_user where login = $1;`
// 	userDB := &postgresModel.UserPostgres{}

// 	err := c.db.Get(userDB, query, login)

// 	if err == sql.ErrNoRows {
// 		return nil, repoErrors.EntityDoesNotExists
// 	} else if err != nil {
// 		return nil, dbErrors.ErrorSelect
// 	}

// 	userModels := &models.User{}
// 	err = copier.Copy(userModels, userDB)

// 	if err != nil {
// 		return nil, dbErrors.ErrorCopy
// 	}

// 	return userModels, nil
// }