package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nabbat/23_kogorta_shotener/cmd/config"
	"github.com/nabbat/23_kogorta_shotener/internal/envirements"
	"github.com/nabbat/23_kogorta_shotener/internal/handlers"
	"github.com/nabbat/23_kogorta_shotener/internal/shotenermaker"
	urlstorage "github.com/nabbat/23_kogorta_shotener/internal/storage"
	"io"
	"log"
	"net/http"
)

// Словарь для хранения соответствий между сокращёнными и оригинальными URL
// TODO Создать хранилище
var urlMap = map[string]string{}

// Перенаправляем по полной ссылке
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid request type", http.StatusBadRequest)
		return
	}
	// Добавляем тестовое соответствие в словарь
	urlMap["aHR0cH"] = "https://practicum.yandex.ru/"
	// Получаем идентификатор из URL-пути
	id := r.URL.Path[1:]

	// Получаем оригинальный URL из словаря

	if originalURL, found := urlMap[id]; found {
		// Устанавливаем заголовок Location и возвращаем ответ с кодом 307
		w.Header().Set("Location", originalURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	http.Error(w, "Ссылка не найдена", http.StatusBadRequest)

}

func shortenURLHandler(w http.ResponseWriter, r *http.Request, c *envirements.EnvConfig) {
	// Читаем тело запроса (URL)
	urlBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения запроса", http.StatusBadRequest)
		return
	}

	// Преобразуем в строку
	url := string(urlBytes)

	// Генерируем уникальный идентификатор сокращённого URL
	id := shotenermaker.GenerateID([]byte(url))

	// Добавляем соответствие в словарь
	urlMap[id] = url

	// Отправляем ответ с сокращённым URL
	shortenedURL := fmt.Sprintf("%s/%s", c.ResultURL, id)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	if _, err := io.WriteString(w, shortenedURL); err != nil {
		log.Fatal(err)
	}
}

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
