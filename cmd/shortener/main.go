package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nabbat/23_kogorta_shotener/cmd/config"
	"github.com/nabbat/23_kogorta_shotener/internal/handlers"
	urlstorage "github.com/nabbat/23_kogorta_shotener/internal/storage"
	"net/http"
)

func main() {
	// Получаем переменные если они есть
	c := config.SetEnv()

	// Создаем хранилище
	storage := urlstorage.NewURLStorage()
	// Создаем хэндлеры
	redirectHandler := &handlers.RedirectHandler{}
	shortenURLHandler := &handlers.ShortenURLHandler{}

	r := mux.NewRouter()
	r.Use(handlers.PanicHandler) // Добавляем PanicHandler middleware

	r.HandleFunc("/", shortenURLHandler.HandleShortenURL(storage, c)).Methods("POST")
	r.HandleFunc("/{idShortenURL}", redirectHandler.HandleRedirect(storage)).Methods("GET")

	fmt.Println("RunAddr: ResultURL: ", c.RunAddr, c.ResultURL)
	fmt.Println("Running server on", c.RunAddr)
	err := http.ListenAndServe(c.RunAddr, r)
	if err != nil {
		panic(err)
	}
}
