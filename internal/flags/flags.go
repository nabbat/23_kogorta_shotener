package flags

import (
	flag "github.com/spf13/pflag"
	"strings"
)

// Flags структура для хранения настроек
type Flags struct {
	RunAddr   string
	ResultURL string
	FileName  string
}

// ParseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func ParseFlags() *Flags {
	// Create a Config instance
	flg := &Flags{}
	flag.StringVarP(&flg.RunAddr, "a", "a", "localhost:8080", "Адрес запуска HTTP-сервера.")
	flag.StringVarP(&flg.ResultURL, "b", "b", "http://localhost:8080", "Адрес результирующего сокращённого URL.")
	flag.StringVarP(&flg.FileName, "f", "f", "short-url-db.json", "Имя файла для записи на диск. пустое значение отключает функцию записи на диск")
	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()
	if !strings.HasPrefix(flg.ResultURL, "http://") {
		flg.ResultURL = "http://" + flg.ResultURL
	}
	return flg
}
