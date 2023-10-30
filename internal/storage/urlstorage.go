package storage

type URLStorage struct {
	urlMap map[string]string
}

func NewURLStorage() *URLStorage {
	storage := &URLStorage{
		urlMap: make(map[string]string),
	}
	// Adding Test Compliance TEST
	//storage.AddURL("aHR0cH", "https://practicum.yandex.ru/")
	return storage
}

// AddURL adds a pair of shortened URL -> original URL
func (storage *URLStorage) AddURL(shortURL, originalURL string) {
	storage.urlMap[shortURL] = originalURL
}

// GetOriginalURL returns the original URL from the shortened URL
func (storage *URLStorage) GetOriginalURL(shortURL string) string {
	return storage.urlMap[shortURL]
}
