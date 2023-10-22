package postgres

import (
	"backend/internal/models"
	dbErrors "backend/internal/pkg/errors/db_errors"
	repoErrors "backend/internal/pkg/errors/repo_errors"
	"backend/internal/repository"
	"backend/internal/repository/postgres/postgres_models"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type HoneyPostgresRepository struct {
	db *sqlx.DB
}

func CreateHoneyPostgresRepository(db *sql.DB) repository.HoneyRepository {
	dbx := sqlx.NewDb(db, "pgx")

	return &HoneyPostgresRepository{db: dbx}
}

func copyHoney(h postgresModel.HoneyPostgres) models.Honey {
	honey := models.Honey{HoneyId: h.HoneyId,
		Name:        h.Name,
		Description: h.Description,
	}

	return honey
}

func (h *HoneyPostgresRepository) GetAllHoney() ([]models.Honey, error) {
	query := `select * from bee_honey;`

	var honeyPostgres []postgresModel.HoneyPostgres
	err := h.db.Select(&honeyPostgres, query)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, dbErrors.ErrorSelect
	}

	honeyModels := []models.Honey{}

	for _, r := range honeyPostgres {
		honey := copyHoney(r)
		honeyModels = append(honeyModels, honey)
	}

	return honeyModels, nil
}
