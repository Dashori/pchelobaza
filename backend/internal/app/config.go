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
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	fmt.Println("error!!!")
	// 	// log.Fatalf("Some error occured. Err: %s", err)
	// }

	c.Address = os.Getenv("BACKEND_HOST")
	c.Port = os.Getenv("BACKEND_PORT")
	c.LogLevel = os.Getenv("LOG_LEVEL")
	c.LogFile = os.Getenv("LOG_FILE")

	c.Postgres.Host = os.Getenv("POSTGRES_HOST")
	c.Postgres.Port = os.Getenv("POSTGRES_PORT")
	c.Postgres.User = os.Getenv("POSTGRES_USER")
	c.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	c.Postgres.DBName = os.Getenv("POSTGRES_DB")

}
