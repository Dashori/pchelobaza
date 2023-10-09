package app

import (
	// "github.com/jmoiron/sqlx"
	"database/sql"
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/services/implementation"
	"backend/internal/pkg/hasher/implementation"
	"backend/internal/repository/postgres"
	"backend/internal/repository"
	"fmt"
	"github.com/charmbracelet/log"
	// _ "github.com/lib/pq"
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

func (a *App) initRepositories() *AppRepositoryFields {
	f := &AppRepositoryFields{
		UserRepository: postgres.CreateUserPostgresRepository(a.PostgresDB),
		// DoctorRepository: postgres_repo.CreateDoctorPostgresRepository(fields),
		// PetRepository:    postgres_repo.CreatePetPostgresRepository(fields),
		// RecordRepository: postgres_repo.CreateRecordPostgresRepository(fields),
	}

	// a.Logger.Info("Success initialization of repositories")

	return f
}

func (a *App) initServices(r *AppRepositoryFields) *AppServiceFields {
	passwordHasher := hasherImplementation.NewBcryptHasher()

	u := &AppServiceFields{
		UserService: servicesImplementation.NewUserImplementation(r.UserRepository, passwordHasher),
		// DoctorService: servicesImplementation.NewDoctorServiceImplementation(r.DoctorRepository, passwordHasher, a.Logger),
		// PetService:    servicesImplementation.NewPetServiceImplementation(r.PetRepository, r.ClientRepository, a.Logger),
		// RecordService: servicesImplementation.NewRecordServiceImplementation(r.RecordRepository, r.DoctorRepository,
		// 	r.ClientRepository, r.PetRepository, a.Logger),
	}

	// a.Logger.Info("Success initialization of services")
	return u
}


func (a *App) Init() error {
	a.Config.ParseConfig()
	fmt.Println("port ", a.Config.Port)
	fmt.Println("port 2 ", a.Config.Postgres.Port)


	var err error
	a.PostgresDB, err = a.InitDB()
	if err != nil {
		return err
	}
	a.Repositories = a.initRepositories()
	newUser2 := models.User {
		Login    :       "dashori",
		Password  :      "aaaa",
		// ConfirmPassword: "aaaa",
		Name           : "dasha",
		Surname        : "chepigo",
		Contacts       : "daahaaa@icloud.com",
	}

	a.Repositories.UserRepository.Create(&newUser2)
	a.Services = a.initServices(a.Repositories)

	newUser := models.NewUser {
		Login    :       "dashori",
		Password  :      "aaaa",
		ConfirmPassword: "aaaa",
		Name           : "dasha",
		Surname        : "chepigo",
		Contacts       : "daahaaa@icloud.com",
	}

	a.Services.UserService.Create(&newUser)
	return nil
}


func (a *App) InitDB() (*sql.DB, error) {
	// a.logger.Debug("POSTGRES! Start init postgreSQL", "user", Config.Postgres.User, "DBName", Config.Postgres.DBName,
		// "host", Config.Postgres.Host, "port", Config.Postgres.Port)

	dsnPGConn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		a.Config.Postgres.User, a.Config.Postgres.DBName, a.Config.Postgres.Password,
		a.Config.Postgres.Host, a.Config.Postgres.Port)
	fmt.Println(dsnPGConn)

	db, err := sql.Open("pgx", dsnPGConn)
	if err != nil {
		fmt.Println("1 error")
		// logger.Fatal("POSTGRES! Error in method open")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		// logger.Fatal("POSTGRES! Error in method ping")
		fmt.Println("2 error ", err)
		return nil, err
	}

	db.SetMaxOpenConns(10)
	fmt.Println("ALL IS OK")

	// logger.Info("POSTGRES! Successfully init postgreSQL")

	return db, nil
}
