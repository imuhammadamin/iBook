package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error while loading configurations")
	}

	cfg := Config{}

	cfg.PostgresHost = cast.ToString(getValueFromEnv("POSTGRES_HOST", "localhost"))
	cfg.PostgresPort = cast.ToInt(getValueFromEnv("POSTGRES_PORT", 5432))
	cfg.PostgresDatabase = cast.ToString(getValueFromEnv("POSTGRES_DATABASE", "postgres"))
	cfg.PostgresUser = cast.ToString(getValueFromEnv("POSTGRES_USER", "postgres"))
	cfg.PostgresPassword = cast.ToString(getValueFromEnv("POSTGRES_PASSWORD", "789"))

	return cfg
}

func getValueFromEnv(key string, defaultValue interface{}) interface{} {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}
