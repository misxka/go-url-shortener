package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/misxka/go-url-shortener/controllers"
	"github.com/misxka/go-url-shortener/storage"
)

type WithStorageHandler struct {
	storage *storage.StorageService
}

func (h *WithStorageHandler) CreateUrlHandler(w http.ResponseWriter, r *http.Request) {
	controllers.CreateUrlHandler(w, r, h.storage)
}

func (h *WithStorageHandler) ShortUrlRedirectHandler(w http.ResponseWriter, r *http.Request) {
	controllers.ShortUrlRedirectHandler(w, r, h.storage)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	h := &WithStorageHandler{storage: storage.InitStorage()}

	router := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	router.Handle("POST /create", controllers.WithJSONContentType(http.HandlerFunc(h.CreateUrlHandler)))
	router.Handle("GET /{url}", controllers.WithJSONContentType(http.HandlerFunc(h.ShortUrlRedirectHandler)))

	fmt.Println("Server starting on port 8080...")
	server.ListenAndServe()
}
