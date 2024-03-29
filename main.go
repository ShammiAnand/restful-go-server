package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type hello struct {
	Message string `json:"message"`
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("cannot load env")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in environment")
	}

	fmt.Println("starting at PORT:", portString)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		respondWithJson(w, 200, hello{
			Message: "Welcome to Shammi's HTTP server; use GET /info for more info",
		},
		)
	})

	serverError := http.ListenAndServe(":"+portString, router)
	if serverError != nil {
		log.Fatal(serverError)
	}

}
