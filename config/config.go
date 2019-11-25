package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config for app
type Config struct {
	Docs bool
	DB   dbConf
}

type dbConf struct {
	Name     string
	Password string
	User     string
	Host     string
}

// New app config
func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("failed to load godotenv")
	}

	docs := getEnvAsBool("DOCS", false)

	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")

	return &Config{
		Docs: docs,
		DB: dbConf{
			Name:     dbName,
			Password: dbPassword,
			User:     dbUser,
			Host:     dbHost,
		},
	}
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := os.Getenv(name)
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultVal
}
