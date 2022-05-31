package config

import (
	"os"
	"strconv"

	"github.com/brxyxn/ticketing-system-telus/backend/app/constants"
	u "github.com/brxyxn/ticketing-system-telus/backend/app/utils"
	"github.com/joho/godotenv"
)

func Configure() (Config, error) {
	var cf Config

	// Env vars
	env := os.Getenv("ENV")
	if env != "Production" {
		err := godotenv.Load()
		if err != nil {
			u.Log.Error("Error loading .env file.", err)
			return Config{}, err
		}
	}

	cf.Port = os.Getenv("PORT")
	if cf.Port == "" {
		cf.Port = constants.Port // Default port if not set
	}

	sql(&cf)
	cache(&cf)

	return cf, nil
}

func sql(cf *Config) {
	cf.Sql.Host = os.Getenv("PG_HOST")
	cf.Sql.Port = os.Getenv("PG_PORT")
	cf.Sql.User = os.Getenv("PG_USER")
	cf.Sql.Name = os.Getenv("PG_NAME")
	cf.Sql.Password = os.Getenv("PG_PASSWORD")
	cf.Sql.Sslmode = os.Getenv("PG_SSLMODE")

	if cf.Sql.Port == "" ||
		cf.Sql.Host == "" ||
		cf.Sql.User == "" ||
		cf.Sql.Name == "" ||
		cf.Sql.Password == "" ||
		cf.Sql.Sslmode == "" {
		return
	}
}

func cache(cf *Config) {
	var err error
	cf.Cache.Host = os.Getenv("REDIS_HOST")
	cf.Cache.Port = os.Getenv("REDIS_PORT")
	cf.Cache.Password = os.Getenv("REDIS_PASSWORD")
	cf.Cache.Name, err = strconv.Atoi(os.Getenv("REDIS_NAME"))
	if err != nil {
		u.Log.Error("Error converting env RDB_NAME to int.", err)
	}

	if cf.Cache.Host == "" ||
		cf.Sql.Port == "" ||
		cf.Cache.Name < 0 {
		return
	}
}
