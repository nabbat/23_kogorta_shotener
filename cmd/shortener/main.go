package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Словарь для хранения соответствий между сокращёнными и оригинальными URL
var urlMap map[string]string

// Перенаправляем по полной ссылке
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем идентификатор из URL-пути
	id := r.URL.Path[1:]

	// Получаем оригинальный URL из словаря
	// TODO Создать хранилище

	if originalURL, found := urlMap[id]; found {
		// Устанавливаем заголовок Location и возвращаем ответ с кодом 307
		w.Header().Set("Location", originalURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	http.Error(w, "Ссылка не найдена", http.StatusBadRequest)

}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	// Читаем тело запроса (URL)
	urlBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения запроса", http.StatusBadRequest)
		return
	}

	// Преобразуем в строку
	url := string(urlBytes)

	// Генерируем уникальный идентификатор сокращённого URL
	id := generateID(url)

	// Добавляем соответствие в словарь
	urlMap[id] = url

	// Отправляем ответ с сокращённым URL
	shortenedURL := fmt.Sprintf("http://localhost:8080/%s", id)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	if _, err := io.WriteString(w, shortenedURL); err != nil {
		log.Fatal(err)
	}
}

// Простая функция для генерации уникального идентификатора
func generateID(fullURL string) string {
	encodedStr := base64.URLEncoding.EncodeToString([]byte(fullURL))
	// Возвращаем первые 6 символов закодированной строки
	if len(encodedStr) > 6 {
		return encodedStr[:6]
	}
	return encodedStr
}

func main() {
	mux := http.NewServeMux()
	urlMap = make(map[string]string)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// GET OR POST
		if r.Method != http.MethodPost {
			redirectHandler(w, r)
			return
		}
		shortenURLHandler(w, r)

	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
