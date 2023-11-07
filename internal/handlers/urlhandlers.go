package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nabbat/23_kogorta_shotener/internal/envirements"
	"github.com/nabbat/23_kogorta_shotener/internal/liblog"
	"github.com/nabbat/23_kogorta_shotener/internal/shotenermaker"
	urlstorage "github.com/nabbat/23_kogorta_shotener/internal/storage"
	"io"
	"net/http"
)

type RedirectHandler struct{}

func (rh *RedirectHandler) HandleRedirect(storage urlstorage.Storage, log liblog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "invalid request type", http.StatusBadRequest)
			log.Info("invalid request type")
			return
		}

		// Получаем идентификатор из URL-пути
		vars := mux.Vars(r)
		shortURL := vars["idShortenURL"]

		// Получаем оригинальный URL
		originalURL, err := storage.GetOriginalURL(shortURL)
		if err != nil {
			log.Info(err)
		}

		if originalURL == "" {
			http.Error(w, "Ссылка не найдена", http.StatusBadRequest)
			log.Info("Ссылка не найдена")
			return
		}
		// Устанавливаем заголовок Location и возвращаем ответ с кодом 307
		w.Header().Set("Location", originalURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
		log.Info("Location set:" + originalURL)

	}
}

type ShortenURLHandler struct{}

func (sh *ShortenURLHandler) HandleShortenURL(storage urlstorage.Storage, c *envirements.EnvConfig, log liblog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Читаем тело запроса (URL)
		defer r.Body.Close()
		urlBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения запроса", http.StatusBadRequest)
			return
		}

		// Генерируем уникальный идентификатор сокращённого URL
		shortURL := shotenermaker.GenerateID(urlBytes)

		// Добавляем соответствие в словарь
		storage.AddURL(shortURL, string(urlBytes))

		// Отправляем ответ с сокращённым URL
		shortenedURL := fmt.Sprintf("%s/%s", c.ResultURL, shortURL)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		if _, err := io.WriteString(w, shortenedURL); err != nil {
			log.Info("Ошибка записи ответа", err)
		}
	}
}

func (sh *ShortenURLHandler) HandleShortenURLJSON(storage urlstorage.Storage, c *envirements.EnvConfig, log liblog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Читаем JSON из тела запроса
		type URLJSONRequest struct {
			URL string `json:"url"`
		}
		var urlJSONRequest URLJSONRequest
		var buf bytes.Buffer

		// читаем тело запроса
		_, err := buf.ReadFrom(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// десериализуем JSON
		if err = json.Unmarshal(buf.Bytes(), &urlJSONRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Генерируем уникальный идентификатор сокращенного URL
		shortURL := shotenermaker.GenerateID([]byte(urlJSONRequest.URL))

		// Добавляем соответствие в словарь
		storage.AddURL(shortURL, urlJSONRequest.URL)

		// Формируем JSON-ответ с сокращенным URL
		response := map[string]string{"result": c.ResultURL + "/" + shortURL}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Error("Ошибка записи JSON-ответа:", err)
			http.Error(w, "Ошибка записи JSON-ответа", http.StatusInternalServerError)
		}
	}
}
