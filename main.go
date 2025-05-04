package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/misxka/go-url-shortener/controllers"
	"github.com/misxka/go-url-shortener/storage"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	storage.InitStorage()

	router := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	router.Handle("POST /create", controllers.WithJSONContentType(http.HandlerFunc(controllers.CreateUrlHandler)))
	router.Handle("GET /{url}", controllers.WithJSONContentType(http.HandlerFunc(controllers.ShortUrlRedirectHandler)))

	fmt.Println("Server starting on port 8080...")
	server.ListenAndServe()
}
