package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// EnvMode - Environment mode
type EnvMode string

const (
	EnvModeDev     EnvMode = "development"
	EnvModeProd    EnvMode = "production"
	EnvModeTesting EnvMode = "testing"
)

// Config - env
type Config struct {
	Port            string
	DbHost          string
	DbPort          string
	DbUser          string
	DbPassword      string
	DbName          string
	ShutdownTimeout time.Duration
}

// NewConfig - Config factory
func NewConfig() (*Config, error) {

	err := godotenv.Load("./config/.env")

	if err != nil {
		log.Print(".env file was not loaded\n")
	}

	port := getEnvOrDefault("PORT", ":3000")

	cfg := &Config{
		port,
		getEnv("DB_HOST"),
		getEnv("DB_PORT"),
		getEnv("DB_USER"),
		getEnv("DB_PASSWORD"),
		getEnv("DB_NAME"),
		time.Second * 5,
	}

	return cfg, nil
}

// PgConnURI - returns connection string
func (c Config) PgConnURI() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", c.DbHost, c.DbPort, c.DbUser, c.DbName, c.DbPassword)
}

func getEnv(key string) string {
	return getEnvOrDefault(key, "")
}

func getEnvOrDefault(key, defaultValue string) string {
	v := os.Getenv(key)

	if v == "" {
		return defaultValue
	}

	return v
}
