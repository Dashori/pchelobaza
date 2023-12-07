package postgres

import (
	"backend/internal/models"
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
	honey := models.Honey{
		HoneyId:     h.HoneyId,
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
		return nil, err
	}

	honeyModels := []models.Honey{}

	for _, r := range honeyPostgres {
		honey := copyHoney(r)
		honeyModels = append(honeyModels, honey)
	}

	return honeyModels, nil
}

func (h *HoneyPostgresRepository) GetHoneyId(name string) (uint64, error) {
	query := `select id from bee_honey where name = $1;`
	var honeyId uint64

	err := h.db.Get(&honeyId, query, name)

	if err == sql.ErrNoRows {
		return 0, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return 0, err
	}

	return honeyId, nil
}

func (h *HoneyPostgresRepository) GetFarmHoney(name string) ([]models.Honey, error) {
	query := `select h.id, h.name, h.description from bee_farm as f
	join bee_farm_honey as fm on f.id = fm.id_farm
	join bee_honey as h on h.id = fm.id_honey
	where f.name = $1;`

	var honeyPostgres []postgresModel.HoneyPostgres
	err := h.db.Select(&honeyPostgres, query, name)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	honeyModels := []models.Honey{}

	for _, r := range honeyPostgres {
		honey := copyHoney(r)
		honeyModels = append(honeyModels, honey)
	}

	return honeyModels, nil
}
