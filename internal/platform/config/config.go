package config

import "os"

const HTTP_ADDR = ":8080"
const DATABASE_URL = "postgres://postgres:postgres@localhost:5432/saas?sslmode=disable"

type Config struct {
	HTTPAddr    string
	DatabaseURL string
}

func Load() Config {
	addr := LoadFromEnvOrDefault("ALPHA_HTTP_ADDR", HTTP_ADDR)
	databaseUrl := LoadFromEnvOrDefault("ALPHA_DATABASE_URL", DATABASE_URL)

	return Config{
		HTTPAddr:    addr,
		DatabaseURL: databaseUrl,
	}
}

func LoadFromEnvOrDefault(envName string, defaultValue string) string {
	envValue := os.Getenv(envName)

	if envValue == "" {
		envValue = defaultValue
	}

	return envValue
}
