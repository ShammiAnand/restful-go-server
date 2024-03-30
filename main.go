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

	addr := fmt.Sprintf(":%v", portString)
	serverError := http.ListenAndServe(addr, mainRouter())

	if serverError != nil {
		log.Fatal(serverError)
	}
}

func mainRouter() http.Handler {

	router := chi.NewRouter()

	_, err := Database()
	if err != nil {
		log.Fatal("couldn't connect to DB")
	}
	fmt.Println("connected to DB")

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/ping"))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Public Routes
	router.Group(func(r chi.Router) {

		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			respondWithJson(w, 200, Message{
				Message: "Welcome to a simple HTTP server; use GET /info for more info",
			},
			)
		})

		router.Post("/users", CreateUserHandler)
		router.Get("/users", GetUsersHandler)
		router.Get("/users/{id}", GetUserWithIdHandler)
		router.Delete("/users/{id}", DeleteUserWithIdHandler)

	})

	return router
}
