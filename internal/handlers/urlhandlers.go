package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nabbat/23_kogorta_shotener/internal/envirements"
	"github.com/nabbat/23_kogorta_shotener/internal/liblog"
	"github.com/nabbat/23_kogorta_shotener/internal/shotenermaker"
	urlstorage "github.com/nabbat/23_kogorta_shotener/internal/storage"
	"io"
	"net/http"
	"time"
)

type RedirectHandler struct{}

func (rh *RedirectHandler) HandleRedirect(storage *urlstorage.URLStorage, log liblog.Logger) http.HandlerFunc {
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
		originalURL := storage.GetOriginalURL(shortURL)

		if originalURL == "" {
			http.Error(w, "Ссылка не найдена", http.StatusBadRequest)
			log.Info("Ссылка не найдена")
			return
		}
		// Устанавливаем заголовок Location и возвращаем ответ с кодом 307
		w.Header().Set("Location", originalURL)
		w.WriteHeader(http.StatusTemporaryRedirect)

	}
}

type ShortenURLHandler struct{}

func (sh *ShortenURLHandler) HandleShortenURL(storage *urlstorage.URLStorage, c *envirements.EnvConfig, log liblog.Logger) http.HandlerFunc {
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

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

type ResponseLogger struct {
	ResponseWriter http.ResponseWriter
	StatusCode     int
}

type responseWriterWrapper struct {
	http.ResponseWriter
	status int
	size   int
}

func (rw *responseWriterWrapper) Status() int {
	return rw.status
}

func (rw *responseWriterWrapper) Size() int {
	return rw.size
}

func (rw *responseWriterWrapper) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.size += n
	return n, err
}

func (rw *responseWriterWrapper) WriteHeader(status int) {
	rw.ResponseWriter.WriteHeader(status)
	rw.status = status
}

func ResponseLoggingMiddleware(log liblog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw := &responseWriterWrapper{ResponseWriter: w}
			next.ServeHTTP(rw, r)

			status := rw.Status()
			size := rw.Size()

			log.Info("Status: ", status, " | Size: ", size, " bytes")
		})
	}
}

func RequestLoggingMiddleware(log liblog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			duration := time.Since(start)

			log.Info("Request: ", r.Method, " ", r.RequestURI, " | Time: ", duration)
		})
	}
}
