package config

import (
	"github.com/caarlos0/env"
	// auto loads .env
	_ "github.com/joho/godotenv/autoload"
)

// Config for app
type Config struct {
	Docs                 bool   `env:"DOCS"`
	CORS                 bool   `env:"CORS" envDefault:"true"`
	Port                 string `env:"PORT" envDefault:"4000"`
	DbName               string `env:"DB_NAME"`
	DbPassword           string `env:"DB_PASSWORD" envDefault:"postgres"`
	DbUser               string `env:"DB_USER" envDefault:"postgres"`
	DbHost               string `env:"DB_HOST" envDefault:"localhost"`
	JwtAccessSecret      string `env:"JWT_ACCESS_SECRET" envDefault:"asdf1234"`
	JwtRefreshSecret     string `env:"JWT_REFRESH_SECRET" envDefault:"1234asdf"`
	JwtIssuer            string `env:"JWT_ISSUER" envDefault:"go-docker-api-boilerplate"`
	JwtAccessExpiration  int64  `env:"JWT_ACCESS_EXPIRATION" envDefault:"1440"`    // minutes
	JwtRefreshExpiration int64  `env:"JWT_REFRESH_EXPIRATION" envDefault:"525600"` // minutes
}

// New app config
func New() (Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	return cfg, err
}
