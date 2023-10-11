package shoturlmaker

import "encoding/base64"

// Простая функция для генерации уникального идентификатора
func generateID(fullURL string) string {
	encodedStr := base64.URLEncoding.EncodeToString([]byte(fullURL))
	// Возвращаем первые 6 символов закодированной строки
	if len(encodedStr) > 6 {
		return encodedStr[:6]
	}
	return encodedStr
}
