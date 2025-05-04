package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/misxka/go-url-shortener/shortener"
	"github.com/misxka/go-url-shortener/storage"
)

type Payload struct {
	OriginalUrl string `json:"originalUrl"`
	UserId      string `json:"userId"`
}

type Response struct {
	Url string `json:"url"`
}

func CreateUrlHandler(w http.ResponseWriter, r *http.Request) {
	var payload Payload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	shortUrl := shortener.GenerateShortUrl(payload.OriginalUrl, payload.UserId)
	storage.SaveUrlMapping(shortUrl, payload.OriginalUrl, payload.UserId)

	host := "http://localhost:8080/"
	response, err := json.Marshal(Response{Url: host + shortUrl})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
