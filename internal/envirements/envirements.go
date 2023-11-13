package envirements

import (
	"os"
	"strings"
)

type EnvConfig struct {
	RunAddr   string
	ResultURL string
	FileName  string
}

// ParseEnv Get system environments
func ParseEnv() *EnvConfig {
	env := &EnvConfig{}
	env.RunAddr = os.Getenv("RUN_ADDR")
	env.ResultURL = os.Getenv("SERVER_ADDRESS")
	env.FileName = os.Getenv("FILE_STORAGE_PATH")
	// парсим переданные серверу аргументы в зарегистрированные переменные
	if !strings.HasPrefix(env.ResultURL, "http://") && env.ResultURL != "" {
		env.ResultURL = "http://" + env.ResultURL
	}
	return env
}
