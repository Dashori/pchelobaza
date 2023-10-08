package app

import (
	// "github.com/jmoiron/sqlx"
	"database/sql"
	"backend/internal/services"
	"backend/internal/repository"
	"fmt"
	"github.com/charmbracelet/log"
	_ "github.com/jackc/pgx/stdlib"
)

type App struct {
	PostgresDB   *sql.DB
	Config       Config
	Repositories *AppRepositoryFields
	Services     *AppServiceFields
	Logger       *log.Logger
}

type AppServiceFields struct {
	UserService services.UserService
	// DoctorService services.DoctorService
	// PetService    services.PetService
	// RecordService services.RecordService
}

type AppRepositoryFields struct {
	UserRepository repository.UserRepository
	// DoctorRepository repository.DoctorRepository
	// PetRepository    repository.PetRepository
	// RecordRepository repository.RecordRepository
}

func (a *App) Init() error {
	var err error
	a.PostgresDB, err = a.InitDB()
	if err != nil {
		return err
	}
	
	return nil
}


func (a *App) InitDB() (*sql.DB, error) {
	// a.logger.Debug("POSTGRES! Start init postgreSQL", "user", Config.Postgres.User, "DBName", Config.Postgres.DBName,
		// "host", Config.Postgres.Host, "port", Config.Postgres.Port)

	dsnPGConn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		a.Config.Postgres.User, a.Config.Postgres.DBName, a.Config.Postgres.Password,
		a.Config.Postgres.Host, a.Config.Postgres.Port)

	db, err := sql.Open("pgx", dsnPGConn)
	if err != nil {
		// logger.Fatal("POSTGRES! Error in method open")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		// logger.Fatal("POSTGRES! Error in method ping")
		return nil, err
	}

	db.SetMaxOpenConns(10)

	// logger.Info("POSTGRES! Successfully init postgreSQL")

	return db, nil
}
