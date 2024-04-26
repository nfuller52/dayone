package application

import (
	"dayone/util"
)

type ApplicationConfig struct {
	DatabaseURL string
	Env         string
	Host        string
	LogLevel    string
	Port        string
}

func Start() *ApplicationConfig {
	return &ApplicationConfig{
		DatabaseURL: util.FetchEnv("DATABASE_URL", "postgres://localhost:5432/dayone_api_development"),
		Env:         util.FetchEnv("ENV", "development"),
		Host:        util.FetchEnv("HOST", "localhost"),
		LogLevel:    util.FetchEnv("LOG_LEVEL", "DEBUG"),
		Port:        util.FetchEnv("PORT", "8080"),
	}
}

func (config *ApplicationConfig) IsDev() bool {
	return config.Env == "development"
}

func (config *ApplicationConfig) IsProd() bool {
	return config.Env == "production"
}
