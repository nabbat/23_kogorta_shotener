package config

import (
	"fmt"
	"github.com/nabbat/23_kogorta_shotener/internal/envirements"
	"github.com/nabbat/23_kogorta_shotener/internal/flags"
)

func SetEnv() *envirements.EnvConfig {
	fl := flags.ParseFlags()
	en := envirements.ParseEnv()
	c := &envirements.EnvConfig{}

	if en.RunAddr != "" {
		c.RunAddr = en.RunAddr
	} else {
		c.RunAddr = fl.RunAddr
	}

	if en.ResultURL != "" && en.RunAddr != "http://" {
		c.ResultURL = en.ResultURL
	} else {
		c.ResultURL = fl.ResultURL
	}
	fmt.Println(c)
	return c
}
