package main

import (
	"github.com/gorilla/mux"
	"github.com/nabbat/23_kogorta_shotener/cmd/config"
	"github.com/nabbat/23_kogorta_shotener/internal/handlers"
	"github.com/nabbat/23_kogorta_shotener/internal/liblog"
	"github.com/nabbat/23_kogorta_shotener/internal/middlewares"
	urlstorage "github.com/nabbat/23_kogorta_shotener/internal/storage"
	"net/http"
)

func main() {
	// Инициализируем логер
	log := liblog.NewLogger()

	// Получаем переменные если они есть
	c := config.SetEnv()

	// Создаем хранилище
	storage := urlstorage.NewURLStorage()
	// Создаем хэндлеры
	redirectHandler := &handlers.RedirectHandler{}
	shortenURLHandler := &handlers.ShortenURLHandler{}

	r := mux.NewRouter()
	r.Use(middlewares.GzipMiddleware(log))
	// Регистрируем middleware для логирования запросов
	r.Use(middlewares.RequestLoggingMiddleware(log))
	// Регистрируем middleware для логирования ответов
	r.Use(middlewares.ResponseLoggingMiddleware(log))
	r.Use(middlewares.PanicHandler) // Добавляем PanicHandler middleware

	r.HandleFunc("/api/shorten", shortenURLHandler.HandleShortenURLJSON(storage, c, log)).Methods("POST")
	r.HandleFunc("/", shortenURLHandler.HandleShortenURL(storage, c, log)).Methods("POST")
	r.HandleFunc("/{idShortenURL}", redirectHandler.HandleRedirect(storage, log)).Methods("GET")

	log.Info("RunAddr: ", c.RunAddr, " | ", "ResultURL: ", c.ResultURL)
	log.Info("Running server on ", c.RunAddr)

	err := http.ListenAndServe(c.RunAddr, r)
	if err != nil {
		panic(err)
	}
}
