package shotenermaker

import (
	"crypto/sha1"
	"encoding/hex"
)

// GenerateID Функция для генерации уникального идентификатора
func GenerateID(fullURL []byte) string {
	h := sha1.New()
	h.Write(fullURL)
	hashBytes := h.Sum(nil)
	encodedStr := hex.EncodeToString(hashBytes)

	// Возвращаем первые 6 символов закодированной строки
	if len(encodedStr) > 6 {
		return encodedStr[:6]
	}
	return encodedStr
}
