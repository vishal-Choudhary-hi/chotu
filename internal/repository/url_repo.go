package repository

import "time"

func SaveURL(code string, originalURL string) error {
	query := "INSERT INTO urls (short_code, original_url) VALUES (?, ?)"
	_, err := DB.Exec(query, code, originalURL)
	return err
}

func GetURL(code string) (string, error) {
	var url string

	query := "SELECT original_url FROM urls WHERE short_code=?"
	err := DB.QueryRow(query, code).Scan(&url)

	if err != nil {
		return "", err
	}

	return url, nil
}

func SaveURLCache(code string, url string) error {
	return RedisClient.Set(Ctx, "url:"+code, url, 24*time.Hour).Err()
}

func GetURLCache(code string) (string, error) {
	return RedisClient.Get(Ctx, "url:"+code).Result()
}
