package service

import (
	"errors"

	"github.com/vishal-Choudhary-hi/chotu/internal/repository"
)

func CreateShortURL(code string, originalURL string) error {

	err := repository.SaveURL(code, originalURL)
	if err != nil {
		return err
	}

	repository.SaveURLCache(code, originalURL)

	return nil
}

func GetOriginalURL(code string) (string, error) {

	url, err := repository.GetURLCache(code)
	if err == nil {
		return url, nil
	}

	url, err = repository.GetURL(code)
	if err != nil {
		return "", errors.New("url not found")
	}

	repository.SaveURLCache(code, url)

	return url, nil
}
