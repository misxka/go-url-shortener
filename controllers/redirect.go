package controllers

import (
	"net/http"

	"github.com/misxka/go-url-shortener/storage"
)

func ShortUrlRedirectHandler(w http.ResponseWriter, r *http.Request, storage *storage.StorageService) {
	shortUrl := r.PathValue("url")
	originalUrl, err := storage.GetOriginalUrl(shortUrl)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalUrl, http.StatusFound)
}
