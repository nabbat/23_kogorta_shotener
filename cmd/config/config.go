package config

import (
	"fmt"
	"github.com/nabbat/23_kogorta_shotener/internal/envirements"
	"github.com/nabbat/23_kogorta_shotener/internal/flags"
)

// Config структура для хранения настроек
type Config struct {
	RunAddr   string
	ResultURL string
}

func SetEnv() *Config {
	fl := flags.ParseFlags()
	en := envirements.ParseEnv()
	c := &Config{}

	if en.EnvRunAddr != "" {
		c.RunAddr = en.EnvRunAddr
	} else {
		c.RunAddr = fl.RunAddr
	}

	if en.EnvResultURL != "" && en.EnvRunAddr != "http://" {
		c.ResultURL = en.EnvResultURL
	} else {
		c.ResultURL = fl.ResultURL
	}
	fmt.Println(c)
	return c
}
