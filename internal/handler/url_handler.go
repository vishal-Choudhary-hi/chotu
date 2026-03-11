package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/vishal-Choudhary-hi/chotu/internal/model"
	"github.com/vishal-Choudhary-hi/chotu/internal/service"
	"github.com/vishal-Choudhary-hi/chotu/pkg"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req model.ShortenRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	baseUrl := os.Getenv("BASE_URL")
	var code string

	for i := 0; i < 5; i++ { // try 5 times
		code = pkg.GenerateCode(6)

		err = service.CreateShortURL(code, req.URL)
		if err == nil {
			break
		}
	}
	if err != nil {
		http.Error(w, "failed to generate unique short url", http.StatusInternalServerError)
		return
	}
	res := model.ShortenResponse{
		ShortURL: baseUrl + "/" + code,
	}

	json.NewEncoder(w).Encode(res)
}
func RedirectURL(w http.ResponseWriter, r *http.Request) {

	code := chi.URLParam(r, "code")
	log.Println("redirecting code:", code)
	url, err := service.GetOriginalURL(code)
	if err != nil {
		http.Error(w, "url not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(url)
}
