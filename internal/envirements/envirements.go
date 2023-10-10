package envirements

import (
	"os"
	"strings"
)

type EnvConfig struct {
	EnvRunAddr   string
	EnvResultURL string
}

// ParseEnv Get system environments
func ParseEnv() *EnvConfig {
	env := &EnvConfig{}
	env.EnvRunAddr = os.Getenv("RUN_ADDR")
	env.EnvResultURL = os.Getenv("SERVER_ADDRESS")
	// парсим переданные серверу аргументы в зарегистрированные переменные
	if !strings.HasPrefix(env.EnvResultURL, "http://") {
		env.EnvResultURL = "http://" + env.EnvResultURL
	}
	return env
}
