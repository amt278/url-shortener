package urlshortener

import(
	"crypto/rand"
    // "fmt"
    "url_shortener/db"
    // "log"
)

func generateShortenedURL() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    length := 7
    shortened := make([]byte, length)
    _, err := rand.Read(shortened)
    if err != nil {
    	return "", err
    }

    for i := 0; i < length; i++ {
    	shortened[i] = charset[int(shortened[i])%len(charset)]
    }
    return string(shortened), nil
}

func SaveURLToDB(original string) (string, error) {
	shortened, err := generateShortenedURL()
	if err != nil {
		return "", err
	}

	if err = db.SaveURL(original, shortened); err != nil {
		return "", err
	}

	return shortened, nil
}

func GetURL(shortened string) (string, error) {
	original, err := db.GetURL(shortened)
	if err != nil {
		return "", err
	}

	return original, nil
}