package app

import (
	// "backend/cmd/modes/flags"
	"github.com/spf13/viper"
)

type PostgresFlags struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
}

type Config struct {
	Postgres PostgresFlags `mapstructure:"postgres"`
	Address  string        `mapstructure:"address"`
	Port     string        `mapstructure:"port"`
	LogLevel string        `mapstructure:"loglevel"`
	LogFile  string        `mapstructure:"logfile"`
	Mode     string        `mapstructure:"mode"`
}

func (c *Config) ParseConfig(configFileName, pathToConfig string) error {
	v := viper.New()
	v.SetConfigName(configFileName)
	v.SetConfigType("json")
	v.AddConfigPath(pathToConfig)

	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	err = v.Unmarshal(c) //  в  json
	if err != nil {
		return err
	}

	return nil
}
