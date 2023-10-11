package envirements

import (
	"os"
	"strings"
)

type EnvConfig struct {
	RunAddr   string
	ResultURL string
}

// ParseEnv Get system environments
func ParseEnv() *EnvConfig {
	env := &EnvConfig{}
	env.RunAddr = os.Getenv("RUN_ADDR")
	env.ResultURL = os.Getenv("SERVER_ADDRESS")
	// парсим переданные серверу аргументы в зарегистрированные переменные
	if !strings.HasPrefix(env.ResultURL, "http://") && env.ResultURL != "" {
		env.ResultURL = "http://" + env.ResultURL
	}
	return env
}
