package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Config struct {
	Env      string `env:"APP_ENV,default=development"`
	Postgres Postgres
	SMTP     SMTP
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST,default=localhost"`
	Port     string `env:"POSTGRES_PORT,default=5432"`
	User     string `env:"POSTGRES_USER,required"`
	Password string `env:"POSTGRES_PASSWORD,required"`
	Name     string `env:"POSTGRES_NAME,required"`
}

// SMTP holds configuration for smtp email.
type SMTP struct {
	Host string `env:"SMTP_HOST,required"`
	Port int    `env:"SMTP_PORT,default=587"`
	User string `env:"SMTP_USER,required"`
	Pass string `env:"SMTP_PASS,required"`
}

// LoadConfig initiate load config either from env file or os env
func LoadConfig(env string) (*Config, error) {
	// just skip loading env files if it is not exists, env files only used in local dev
	_ = godotenv.Load(env)

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding env")
	}

	return &config, nil
}
