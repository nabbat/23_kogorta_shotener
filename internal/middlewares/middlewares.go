package middlewares

import (
	"compress/gzip"
	"github.com/gorilla/mux"
	"github.com/nabbat/23_kogorta_shotener/internal/liblog"
	"io"
	"net/http"
	"strings"
	"time"
)

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

// compressWriter реализует интерфейс http.ResponseWriter и позволяет прозрачно для сервера
// сжимать передаваемые данные и выставлять правильные HTTP-заголовки
type compressWriter struct {
	w  http.ResponseWriter
	zw *gzip.Writer
}

func newCompressWriter(w http.ResponseWriter) *compressWriter {
	return &compressWriter{
		w:  w,
		zw: gzip.NewWriter(w),
	}
}

func (c *compressWriter) Header() http.Header {
	return c.w.Header()
}

func (c *compressWriter) Write(p []byte) (int, error) {
	return c.zw.Write(p)
}

func (c *compressWriter) WriteHeader(statusCode int) {
	if statusCode < 300 {
		c.w.Header().Set("Content-Encoding", "gzip")
	}
	c.w.WriteHeader(statusCode)
}

// Close закрывает gzip.Writer и досылает все данные из буфера.
func (c *compressWriter) Close() error {
	return c.zw.Close()
}

// compressReader реализует интерфейс io.ReadCloser и позволяет прозрачно для сервера
// декомпрессировать получаемые от клиента данные
type compressReader struct {
	r  io.ReadCloser
	zr *gzip.Reader
}

func newCompressReader(r io.ReadCloser) (*compressReader, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	return &compressReader{
		r:  r,
		zr: zr,
	}, nil
}

func (c compressReader) Read(p []byte) (n int, err error) {
	return c.zr.Read(p)
}

func (c *compressReader) Close() error {
	if err := c.r.Close(); err != nil {
		return err
	}
	return c.zr.Close()
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

// gzipMiddleware Делает сжатие и распаковку если это возможно
func GzipMiddleware(log liblog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// по умолчанию устанавливаем оригинальный http.ResponseWriter как тот,
			// который будем передавать следующей функции
			ow := w

			// проверяем, что клиент умеет получать от сервера сжатые данные в формате gzip
			acceptEncoding := r.Header.Get("Accept-Encoding")
			supportsGzip := strings.Contains(acceptEncoding, "gzip")
			if supportsGzip {
				// оборачиваем оригинальный http.ResponseWriter новым с поддержкой сжатия
				cw := newCompressWriter(w)
				// меняем оригинальный http.ResponseWriter на новый
				ow = cw
				// не забываем отправить клиенту все сжатые данные после завершения middleware
				defer cw.Close()
				log.Info("COMPRESSION SUPPORTED")
			} else {
				log.Info("COMPRESSION UNSUPPORTED")
			}
			// проверяем, что клиент отправил серверу сжатые данные в формате gzip
			contentEncoding := r.Header.Get("Content-Encoding")
			sendsGzip := strings.Contains(contentEncoding, "gzip")
			if sendsGzip {
				// оборачиваем тело запроса в io.Reader с поддержкой декомпрессии
				cr, err := newCompressReader(r.Body)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				// меняем тело запроса на новое
				r.Body = cr
				defer cr.Close()
			}
			// передаём управление хендлеру
			next.ServeHTTP(ow, r)
		})
	}
}
