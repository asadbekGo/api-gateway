package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Enviroment string // develop, staging, production

	TodoServiceHost string
	TodoServicePort int

	// context timeout in seconds
	CtxTimeout int

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Enviroment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))
	c.TodoServiceHost = cast.ToString(getOrReturnDefault("TODO_SERVICE_HOST", "localhost"))
	c.TodoServicePort = cast.ToInt(getOrReturnDefault("TODO_SERVICE_PORT", 50051))
	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
