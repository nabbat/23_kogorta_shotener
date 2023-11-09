package storage

type Storage interface {
	AddURL(shortURL, originalURL string) error
	GetOriginalURL(shortURL string) (string, error)
	Close()
}
