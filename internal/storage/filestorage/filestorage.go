package filestorage

import (
	"github.com/nabbat/23_kogorta_shotener/internal/liblog"
	"os"
)

type URLDataJSON struct {
	UUID        int    `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type File struct {
	os.File
}

// NewFileStorage Создает или подключает существующий файл
func NewFileStorage(filename string, log liblog.Logger) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Info("Failed to open or create the file: %v", err)
		return nil, err
	}
	return file, nil
}

// AddURL adds a pair of shortened URL -> original URL
func (storage *File) AddURL(shortURL, originalURL string) error {
	// Нужно не просто добавить URL а еще и проставить ему uuid
	//storage.urlMap[shortURL] = originalURL
}

// GetOriginalURL returns the original URL from the shortened URL
func (storage *File) GetOriginalURL(shortURL string) (string, error) {
	//return storage.urlMap[shortURL]
}
