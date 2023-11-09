package internalstorage

type InternalURLStorage struct {
	urlMap map[string]string
}

func NewURLStorage() *InternalURLStorage {
	storage := &InternalURLStorage{
		urlMap: make(map[string]string),
	}
	return storage
}

// AddURL adds a pair of shortened URL -> original URL
func (storage *InternalURLStorage) AddURL(shortURL, originalURL string) error {
	storage.urlMap[shortURL] = originalURL
	return nil
}

// GetOriginalURL returns the original URL from the shortened URL
func (storage *InternalURLStorage) GetOriginalURL(shortURL string) (string, error) {
	return storage.urlMap[shortURL], nil
}

// Close returns the original URL from the shortened URL
func (storage *InternalURLStorage) Close() {
	// May be late
}
