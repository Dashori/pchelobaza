package postgres

import (
	"backend/internal/models"
	repoErrors "backend/internal/pkg/errors/repo_errors"
	"backend/internal/repository"
	"backend/internal/repository/postgres/postgres_models"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type FarmPostgresRepository struct {
	db *sqlx.DB
}

func CreateFarmPostgresRepository(db *sql.DB) repository.FarmRepository {
	dbx := sqlx.NewDb(db, "pgx")

	return &FarmPostgresRepository{db: dbx}
}

func copyFarm(f postgresModel.FarmPostgres) models.Farm {
	farm := models.Farm{
		FarmId:      f.FarmId,
		UserId:      f.UserId,
		UserLogin:   f.UserLogin,
		Name:        f.Name,
		Description: f.Description,
		Address:     f.Address,
	}

	return farm
}

func (f *FarmPostgresRepository) CreateFarm(farm *models.Farm) error {
	query := `insert into bee_farm(id_user, name, description, address) values($1, $2, $3, $4);`

	_, err := f.db.Exec(query, farm.UserId, farm.Name, farm.Description, farm.Address)

	if err != nil {
		return err
	}

	return nil
}

func (f *FarmPostgresRepository) GetFarmByName(name string) (*models.Farm, error) {
	query := `select f.id, f.id_user, f.name, f.description,
	f.address, u.login
	from bee_farm as f
	join bee_user as u on f.id_user = u.id
	where f.name = $1;`
	farmDB := &postgresModel.FarmPostgres{}

	err := f.db.Get(farmDB, query, name)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	farmModel := copyFarm(*farmDB)

	return &farmModel, nil
}

func (f *FarmPostgresRepository) GetUsersFarm(userId uint64, limit int, skipped int) ([]models.Farm, error) {
	query := `select * from bee_farm where id_user = $1
	offset $2
	limit $3;`

	var farmPostgres []postgresModel.FarmPostgres
	err := f.db.Select(&farmPostgres, query, userId, skipped, limit)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	farmModels := []models.Farm{}

	for _, r := range farmPostgres {
		farm := copyFarm(r)
		farmModels = append(farmModels, farm)
	}

	return farmModels, nil
}

func (f *FarmPostgresRepository) PatchFarm(farm *models.Farm) error {
	query := `update bee_farm set description = $1, address = $2 where id = $3;`

	_, err := f.db.Exec(query, farm.Description, farm.Address, farm.FarmId)

	if err != nil {
		return err
	}

	return nil
}
