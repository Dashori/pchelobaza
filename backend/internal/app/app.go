package app

import (
	// "github.com/jmoiron/sqlx"
	"backend/internal/models"
	"backend/internal/pkg/hasher/implementation"
	"backend/internal/repository"
	"backend/internal/repository/postgres"
	"backend/internal/services"
	"backend/internal/services/implementation"
	"database/sql"
	"fmt"
	"os"
	// _ "github.com/lib/pq"
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
	UserService  services.UserService
	HoneyService services.HoneyService
	// DoctorService services.DoctorService
	// PetService    services.PetService
	// RecordService services.RecordService
}

type AppRepositoryFields struct {
	UserRepository  repository.UserRepository
	HoneyRepository repository.HoneyRepository
	// DoctorRepository repository.DoctorRepository
	// PetRepository    repository.PetRepository
	// RecordRepository repository.RecordRepository
}

func (a *App) initRepositories() *AppRepositoryFields {
	f := &AppRepositoryFields{
		UserRepository:  postgres.CreateUserPostgresRepository(a.PostgresDB),
		HoneyRepository: postgres.CreateHoneyPostgresRepository(a.PostgresDB),
		// DoctorRepository: postgres_repo.CreateDoctorPostgresRepository(fields),
		// PetRepository:    postgres_repo.CreatePetPostgresRepository(fields),
		// RecordRepository: postgres_repo.CreateRecordPostgresRepository(fields),
	}

	a.Logger.Info("Success initialization of repositories")

	return f
}

func (a *App) initServices(r *AppRepositoryFields) *AppServiceFields {
	passwordHasher := hasherImplementation.NewBcryptHasher()

	u := &AppServiceFields{
		UserService:  servicesImplementation.NewUserImplementation(r.UserRepository, passwordHasher, a.Logger),
		HoneyService: servicesImplementation.NewHoneyImplementation(r.HoneyRepository, a.Logger),
		// DoctorService: servicesImplementation.NewDoctorServiceImplementation(r.DoctorRepository, passwordHasher, a.Logger),
		// PetService:    servicesImplementation.NewPetServiceImplementation(r.PetRepository, r.ClientRepository, a.Logger),
		// RecordService: servicesImplementation.NewRecordServiceImplementation(r.RecordRepository, r.DoctorRepository,
		// 	r.ClientRepository, r.PetRepository, a.Logger),
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

	Logger.Print("\n")
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

	newUser := models.User{
		Login:           "dashori6",
		Password:        "abcde",
		ConfirmPassword: "abcd",
		Name:            "dasha",
		Surname:         "chepigo",
		Contact:         "daahaaa@icloud.com",
	}

	user, err := a.Services.UserService.Create(&newUser)
	if err != nil {
		fmt.Println("create ", err)
	} else {
		fmt.Println(user.Name)
	}
	user, err = a.Services.UserService.Login("dashori5", "abcd")
	if err != nil {
		fmt.Println("login", err)
	} else {
		fmt.Println(user.Name, user.Surname)
	}
	user2, err := a.Services.UserService.GetUserByLogin("dashori5")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user2.Name, user2.Surname)
	}

	userup := models.UserPatch{
		Login: "dashori3",
		Name:  "arisha",
	}

	err = a.Services.UserService.Update(&userup)
	if err != nil {
		fmt.Println(err)
	} else {
		user2, _ = a.Services.UserService.GetUserByLogin("dashori3")
		fmt.Println(user2.Name, user2.Surname)
	}

	Honey, err := a.Services.HoneyService.GetAllHoney()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(Honey)
	}

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
		fmt.Println("1 error")
		a.Logger.Fatal("POSTGRES! Error in method open")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		a.Logger.Fatal("POSTGRES! Error in method ping")
		fmt.Println("2 error ", err)
		return nil, err
	}

	db.SetMaxOpenConns(10)

	a.Logger.Info("POSTGRES! Successfully init postgreSQL")

	return db, nil
}
