package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	HTTPPort         string
}

func Load() Config {
	c := Config{}

	c.HTTPPort = cast.ToString(Look("HTTP_PORT","8888"))

	c.PostgresHost = cast.ToString(Look("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(Look("POSTGRES_HOST", 5432))
	c.PostgresDB = cast.ToString(Look("POSTGRES_DATABASE", "mybook"))
	c.PostgresUser = cast.ToString(Look("POSTGRES_USER", "najmiddin"))
	c.PostgresPassword = cast.ToString(Look("POSTGRES_PASSWORD", "1234"))

	return c
}

func Look(key string, defVal interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defVal
}
