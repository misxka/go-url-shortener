package controllers

import (
	"net/http"

	"github.com/misxka/go-url-shortener/storage"
)

func ShortUrlRedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.PathValue("url")
	originalUrl := storage.GetOriginalUrl(shortUrl)
	http.Redirect(w, r, originalUrl, http.StatusFound)
}
