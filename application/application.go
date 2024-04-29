package application

import (
	"dayone/application/database"
	"dayone/cli"
	"dayone/src/util"
	"os"
)

type ApplicationConfig struct {
	DatabaseURL string
	Env         string
	Host        string
	LogLevel    string
	Port        string
}

func Start() *ApplicationConfig {
	config := &ApplicationConfig{
		DatabaseURL: util.FetchEnv("DATABASE_URL", "postgres://localhost:5432/dayone_development"),
		Env:         util.FetchEnv("ENV", "development"),
		Host:        util.FetchEnv("HOST", "localhost"),
		LogLevel:    util.FetchEnv("LOG_LEVEL", "DEBUG"),
		Port:        util.FetchEnv("PORT", "8080"),
	}

	database.ConnectToDB(config.DatabaseURL)

	if len(os.Args) > 1 {
		cli.RunCLI(os.Args)

		os.Exit(0)
	}

	return config
}

func (config *ApplicationConfig) ServerUrl() string {
	return config.Host + ":" + config.Port
}

func (config *ApplicationConfig) IsDev() bool {
	return config.Env == "development"
}

func (config *ApplicationConfig) IsProd() bool {
	return config.Env == "production"
}
