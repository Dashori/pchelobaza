package app

import (
	// "fmt"
	// "github.com/joho/godotenv"
	"os"
)

type PostgresFlags struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Config struct {
	Postgres PostgresFlags
	Address  string
	Port     string
	LogLevel string
	LogFile  string
}

func (c *Config) ParseConfig() {
	// при локальном запуске свой .env файл
	// err := godotenv.Load(".env") 
	// if err != nil {
	// 	fmt.Println("error!!!")
	// }
	// c.Address = os.Getenv("BACKEND_HOST")

	c.Port = os.Getenv("BACKEND_PORT")
	c.LogLevel = os.Getenv("LOG_LEVEL")
	c.LogFile = os.Getenv("LOG_FILE")

	c.Postgres.Host = os.Getenv("POSTGRESQL_HOST")
	c.Postgres.Port = os.Getenv("POSTGRESQL_PORT")
	c.Postgres.User = os.Getenv("POSTGRESQL_USERNAME")
	c.Postgres.Password = os.Getenv("POSTGRESQL_PASSWORD")
	c.Postgres.DBName = os.Getenv("POSTGRESQL_DATABASE")

}
