package filestorage

import (
	"bufio"
	"encoding/json"
	"errors"
	"github.com/nabbat/23_kogorta_shotener/internal/liblog"
	"os"
)

type URLDataJSON struct {
	UUID        int    `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type NewFile struct {
	os.File
	i int
}

// NewFileStorage Создает или подключает существующий файл
func NewFileStorage(filename string, log liblog.Logger, storage *NewFile) (*NewFile, error) {
	// TRYING TO OPEN A FILE
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	// IF THE FILE COULD NOT BE OPENED CREATE IT
	if err != nil {
		_, err := os.OpenFile(filename, os.O_CREATE, 0666)
		if err != nil {
			log.Info("Failed to open or create the file: %v", err)
			return nil, err
		}
		storage.i = 0
		return storage, nil
	}
	// Прочитать последнее значение UUID
	lastUUID, err := readLastUUID(file)
	if err != nil {
		log.Info("Failed to read the last UUID: %v", err)
		storage.i = 0
		return nil, err
	}

	storage.i = lastUUID
	return storage, nil
}

// readLastUUID FUNCTION TO READ THE LATEST UUID FROM A FILE
func readLastUUID(file *os.File) (int, error) {
	var lastUUID int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var urlData URLDataJSON
		err := json.Unmarshal([]byte(line), &urlData)
		if err != nil {
			return 0, err
		}
		lastUUID = urlData.UUID
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lastUUID, nil
}

// AddURL adds a pair of shortened URL -> original URL
func (storage *NewFile) AddURL(shortURL, originalURL string) error {
	storage.i++
	u := URLDataJSON{
		UUID:        storage.i,
		ShortURL:    shortURL,
		OriginalURL: originalURL,
	}

	d, err := json.Marshal(u)
	if err != nil {
		return err
	}

	// Добавляем разделитель (например, перевод строки) между записями
	d = append(d, '\n')

	_, err = storage.File.Write(d)
	if err != nil {
		return err
	}

	return nil
}

// GetOriginalURL returns the original URL from the shortened URL
func (storage *NewFile) GetOriginalURL(shortURL string) (string, error) {
	s := bufio.NewScanner(&storage.File)

	for s.Scan() {
		buffer := s.Bytes()
		u := URLDataJSON{}
		err := json.Unmarshal(buffer, &u)
		if err != nil {
		}

		if u.ShortURL == shortURL {
			return u.OriginalURL, nil
		}
	}

	return "", errors.New("Short url not found")
}
