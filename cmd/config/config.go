package config

import (
	"github.com/nabbat/23_kogorta_shotener/internal/envirements"
	"github.com/nabbat/23_kogorta_shotener/internal/flags"
	"github.com/nabbat/23_kogorta_shotener/internal/liblog"
	"github.com/nabbat/23_kogorta_shotener/internal/storage"
	"github.com/nabbat/23_kogorta_shotener/internal/storage/filestorage"
	urlstorage "github.com/nabbat/23_kogorta_shotener/internal/storage/internalstorage"
)

type Config struct {
	RunAddr   string
	ResultURL string
	FileName  string
}

func SetEnv(log liblog.Logger) (storage.Storage, *Config, error) {
	fl := flags.ParseFlags()
	en := envirements.ParseEnv()
	c := &Config{}

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

	if en.FileName != "" {
		c.FileName = en.FileName
		st, _ := filestorage.NewFileStorage(c.FileName, log, &filestorage.NewFile{})

		return st, c, nil
	}

	if fl.FileName != "" {
		c.FileName = fl.FileName
		st, _ := filestorage.NewFileStorage(c.FileName, log, &filestorage.NewFile{})

		return st, c, nil
	}

	st := urlstorage.NewURLStorage()
	return st, c, nil
}
