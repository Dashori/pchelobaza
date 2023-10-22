package postgres

import (
	"backend/internal/models"
	repoErrors "backend/internal/pkg/errors/repo_errors"
	"backend/internal/repository"
	"backend/internal/repository/postgres/postgres_models"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type RequestPostgresRepository struct {
	db *sqlx.DB
}

func CreateRequestPostgresRepository(db *sql.DB) repository.RequestRepository {
	dbx := sqlx.NewDb(db, "pgx")

	return &RequestPostgresRepository{db: dbx}
}

func copyRequest(r postgresModel.RequestRostgres) models.Request {
	request := models.Request{
		RequestId:   r.RequestId,
		UserId:      r.UserId,
		UserLogin:   r.UserLogin,
		Description: r.Description,
		Status:      r.Status,
	}

	return request
}

func (r *RequestPostgresRepository) Create(request *models.Request) error {
	query := `insert into bee_request(id_user, description, status) values($1, $2, $3);`

	_, err := r.db.Exec(query, request.UserId, request.Description, request.Status)

	if err != nil {
		return err
	}

	return nil
}

func (r *RequestPostgresRepository) GetAllRequests() ([]models.Request, error) {
	query := `select r.id, r.description, r.status, u.login from bee_request r join bee_user u on r.id_user = u.id
	order by status desc`

	var requestPostgres []postgresModel.RequestRostgres
	err := r.db.Select(&requestPostgres, query)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	requestModels := []models.Request{}

	for _, r := range requestPostgres {
		request := copyRequest(r)
		requestModels = append(requestModels, request)
	}

	return requestModels, nil
}

func (r *RequestPostgresRepository) GetRequestsPagination(limit int, skipped int) ([]models.Request, error) {
	query := `select r.id, r.description, r.status, u.login from bee_request r join bee_user u on r.id_user = u.id
	order by status desc
	offset $1 
	limit $2;`

	var requestPostgres []postgresModel.RequestRostgres
	err := r.db.Select(&requestPostgres, query, skipped, limit)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	requestModels := []models.Request{}

	for _, r := range requestPostgres {
		request := copyRequest(r)
		requestModels = append(requestModels, request)
	}

	return requestModels, nil
}

func (r *RequestPostgresRepository) GetUserRequest(UserLogin string) (*models.Request, error) {
	query := `select r.id, r.description, r.status, u.login from bee_request r join bee_user u on r.id_user = u.id
	where u.login = $1;`

	requestPostgres := &postgresModel.RequestRostgres{}
	err := r.db.Get(requestPostgres, query, UserLogin)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	requestModel := copyRequest(*requestPostgres)

	return &requestModel, nil
}

func (r *RequestPostgresRepository) PatchUserRequest(request *models.Request) error {
	query := `update bee_request
	set status = $1
	from bee_request r join bee_user u on r.id_user = u.id
	where u.login = $2;`

	_, err := r.db.Exec(query, request.Status, request.UserLogin)

	if err != nil {
		return err
	}

	return nil
}
