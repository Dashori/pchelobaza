package app

import (
	// "backend/internal/models"
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
	// "time"
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
		HoneyService:      servicesImplementation.NewHoneyImplementation(r.HoneyRepository, a.Logger),
		RequestService:    servicesImplementation.NewRequestImplementation(r.RequestRepository, r.UserRepository, a.Logger),
		FarmService:       servicesImplementation.NewFarmImplementation(r.FarmRepository, r.UserRepository, a.Logger),
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

	// newUser := models.User{
	// 	Login:           "dashori6",
	// 	Password:        "abcde",
	// 	ConfirmPassword: "abcde",
	// 	Name:            "dasha",
	// 	Surname:         "chepigo",
	// 	Contact:         "daahaaa@icloud.com",
	// 	Role:            "beeman",
	// }

	// user, err := a.Services.UserService.Create(&newUser)
	// if err != nil {
	// 	fmt.Println("create ", err)
	// } else {
	// 	fmt.Println(user.Name)
	// }
	// user, err = a.Services.UserService.Login("dashori5", "abcd")
	// if err != nil {
	// 	fmt.Println("login", err)
	// } else {
	// 	fmt.Println(user.Name, user.Surname)
	// }
	// user2, err := a.Services.UserService.GetUserByLogin("dashori6")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(user2)
	// }

	// userup := models.User{
	// 	Login: "dashori6",
	// 	Name:  "arisha",
	// }

	// err = a.Services.UserService.Update(&userup)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	user2, _ = a.Services.UserService.GetUserByLogin("dashori6")
	// 	fmt.Println(user2.Name, user2.Surname)
	// }

	// Honey, err := a.Services.HoneyService.GetAllHoney()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(Honey)
	// }

	// fmt.Println("\n\n")
	// req, err := a.Services.RequestService.GetRequestsPagination(5, 4)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(req)
	// }

	// fmt.Println("\n\n")
	// req, err = a.Services.RequestService.GetAllRequests()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(req)
	// }

	// fmt.Println("\n\n")
	// req2, err := a.Services.RequestService.GetUserRequest("dashori6")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(req2)
	// }

	// req5 := models.Request{
	// 	UserLogin:   "dashori6",
	// 	Description: "need",
	// }

	// req3, err := a.Services.RequestService.CreateRequest(&req5)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(req3)
	// }

	// requp := models.Request{
	// 	UserLogin: "Wood52",
	// 	Status:    "approved",
	// }

	// err = a.Services.RequestService.PatchUserRequest(requp)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	req2, err := a.Services.RequestService.GetUserRequest("Lindsey69")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(req2)
	// 	}
	// }

	// farm, err := a.Services.FarmService.GetUsersFarm("Pacheco30", 5, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(farm)
	// }

	// confs, err := a.Services.ConferenceService.GetAllConferences(5, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(confs)
	// }

	// conf := models.Conference{
	// 	Name:        "myconf2",
	// 	UserLogin:   "Black99",
	// 	Description: "abcde",
	// 	Date:        time.Date(2023, 11, 11, 19, 00, 00, 00, time.UTC),
	// 	Address:     "aaaa",
	// 	MaxUsers:    12,
	// }
	// confPatch := models.Conference{
	// 	Name:        "myconf2",
	// 	UserLogin:   "Black99",
	// 	Description: "abcdefghijklmnop",
	// 	Date:        time.Date(2023, 11, 11, 19, 00, 00, 00, time.UTC),
	// 	Address:     "aaaa",
	// 	MaxUsers:    12,
	// }

	// newConf, err := a.Services.ConferenceService.CreateConference(&conf)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(newConf)
	// }
	// err = a.Services.ConferenceService.PatchConference(&confPatch)
	// fmt.Println("!!!", err)

	// newConf, err = a.Services.ConferenceService.GetConferenceByName("myconf2")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(newConf)
	// }

	// usersss, err := a.Services.ConferenceService.GetConferenceUsers("Conference1", 5, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(usersss)
	// }

	// reviews, err := a.Services.ConferenceService.GetConferenceReviews("Conference1", 5, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(reviews)
	// }
	// err = a.Services.ConferenceService.PatchConferenceUsers("myconf2", "Black99")
	// fmt.Println("***", err)
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
