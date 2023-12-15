package app

import (
	dbErrors "backend/internal/pkg/errors/db_errors"
	"backend/internal/pkg/hasher/implementation"
	"backend/internal/repository"
	"backend/internal/repository/postgres"
	"backend/internal/services"
	"backend/internal/services/implementation"
	"database/sql"
	"fmt"
	"github.com/charmbracelet/log"
	_ "github.com/jackc/pgx/stdlib"
	"os"
)

type App struct {
	PostgresDB   *sql.DB
	Config       Config
	Repositories *AppRepositoryFields
	Services     *AppServiceFields
	Logger       *log.Logger
}

type AppServiceFields struct {
	UserService       services.UserService
	HoneyService      services.HoneyService
	RequestService    services.RequestService
	FarmService       services.FarmService
	ConferenceService services.ConferenceService
}

type AppRepositoryFields struct {
	UserRepository       repository.UserRepository
	HoneyRepository      repository.HoneyRepository
	RequestRepository    repository.RequestRepository
	FarmRepository       repository.FarmRepository
	ConferenceRepository repository.ConferenceRepository
}

func (a *App) initRepositories() *AppRepositoryFields {
	f := &AppRepositoryFields{
		UserRepository:       postgres.CreateUserPostgresRepository(a.PostgresDB),
		HoneyRepository:      postgres.CreateHoneyPostgresRepository(a.PostgresDB),
		RequestRepository:    postgres.CreateRequestPostgresRepository(a.PostgresDB),
		FarmRepository:       postgres.CreateFarmPostgresRepository(a.PostgresDB),
		ConferenceRepository: postgres.CreateConferencePostgresRepository(a.PostgresDB),
	}

	a.Logger.Info("Success initialization of repositories")

	return f
}

func (a *App) initServices(r *AppRepositoryFields) *AppServiceFields {
	passwordHasher := hasherImplementation.NewBcryptHasher()

	u := &AppServiceFields{
		UserService:       servicesImplementation.NewUserImplementation(r.UserRepository, passwordHasher, a.Logger),
		HoneyService:      servicesImplementation.NewHoneyImplementation(r.HoneyRepository, r.FarmRepository, a.Logger),
		RequestService:    servicesImplementation.NewRequestImplementation(r.RequestRepository, r.UserRepository, a.Logger),
		FarmService:       servicesImplementation.NewFarmImplementation(r.FarmRepository, r.UserRepository, r.HoneyRepository, a.Logger),
		ConferenceService: servicesImplementation.NewConferenceImplementation(r.ConferenceRepository, r.UserRepository, a.Logger),
	}

	a.Logger.Info("Success initialization of services")
	return u
}

func (a *App) initLogger() {
	f, err := os.OpenFile(a.Config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	Logger := log.New(f)

	log.SetFormatter(log.LogfmtFormatter)
	Logger.SetReportTimestamp(true)
	Logger.SetReportCaller(true)

	if a.Config.LogLevel == "debug" {
		Logger.SetLevel(log.DebugLevel)
	} else if a.Config.LogLevel == "info" {
		Logger.SetLevel(log.InfoLevel)
	} else {
		log.Fatal("Error log level")
	}

	Logger.Info("Success initialization of new Logger!")

	a.Logger = Logger
}

func (a *App) Init() error {
	a.Config.ParseConfig()
	a.initLogger()
	var err error
	a.PostgresDB, err = a.InitDB()
	if err != nil {
		return err
	}

	a.Repositories = a.initRepositories()
	a.Services = a.initServices(a.Repositories)
	return nil
}

func (a *App) InitDB() (*sql.DB, error) {
	a.Logger.Debug("POSTGRES! Start init postgreSQL", "user", a.Config.Postgres.User, "DBName", a.Config.Postgres.DBName,
		"host", a.Config.Postgres.Host, "port", a.Config.Postgres.Port)

	dsnPGConn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		a.Config.Postgres.User, a.Config.Postgres.DBName, a.Config.Postgres.Password,
		a.Config.Postgres.Host, a.Config.Postgres.Port)
	fmt.Println(dsnPGConn)

	db, err := sql.Open("pgx", dsnPGConn)
	if err != nil {
		a.Logger.Fatal("POSTGRES! Error in method open", err)
		return nil, dbErrors.ErrorInitDB
	}

	err = db.Ping()
	if err != nil {
		a.Logger.Fatal("POSTGRES! Error in method ping", err)
		return nil, dbErrors.ErrorInitDB
	}

	db.SetMaxOpenConns(10)

	a.Logger.Info("POSTGRES! Successfully init postgreSQL")

	return db, nil
}
