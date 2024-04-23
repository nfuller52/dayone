package config

import (
	"dayone/util"
)

type ApplicationConfig struct {
	DatabaseURL string
	Env         string
	LogLevel    string
	Port        string
}

func LoadApplicationConfig() *ApplicationConfig {
	config := &ApplicationConfig{
		DatabaseURL: util.FetchEnv("DATABASE_URL", "postgres://localhost:5432/dayone_api_development"),
		Env:         util.FetchEnv("ENV", "development"),
		LogLevel:    util.FetchEnv("LOG_LEVEL", "DEBUG"),
		Port:        util.FetchEnv("PORT", "8080"),
	}

	switch config.Env {
	case "production":
		LoadProductionConfig(config)
	default:
		LoadDevelopmentConfig(config)
	}

	return config
}

func LoadDevelopmentConfig(config *ApplicationConfig) {
	// TODO: Dev specific settings
}

func LoadProductionConfig(config *ApplicationConfig) {
	// TODO: Production specific settings
}

func (config *ApplicationConfig) IsDevelopment() bool {
	return config.Env == "development"
}

func (config *ApplicationConfig) IsProduction() bool {
	return config.Env == "production"
}
