package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type userParams struct {
	Name     string  `json:"name"`
	Email    *string `json:"email"`
	JobTitle string  `json:"job_title"`
	Age      uint8   `json:"age"`
}

type message struct {
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
	serverError := http.ListenAndServe(":"+portString, mainRouter())
	if serverError != nil {
		log.Fatal(serverError)
	}
}

func mainRouter() http.Handler {

	router := chi.NewRouter()

	db, err := database()
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

	// NOTE: Public Routes
	router.Group(func(r chi.Router) {
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			respondWithJson(w, 200, message{
				Message: "Welcome to Shammi's HTTP server; use GET /info for more info",
			},
			)
		})
		router.Post("/users", func(w http.ResponseWriter, r *http.Request) {

			decoder := json.NewDecoder(r.Body)
			payload := userParams{}
			err := decoder.Decode(&payload)
			if err != nil {
				respondWithError(w, 403, fmt.Sprintf("Error Parsing JSON: %v", err))
			}

			newUser := User{
				Name:     payload.Name,
				Age:      payload.Age,
				Email:    payload.Email,
				JobTitle: payload.JobTitle,
			}

			result := db.Create(&newUser)

			if result.Error != nil {
				respondWithError(w, 403, fmt.Sprintf("failed to create: %v", result.Error))
			}

			respondWithJson(w, 201, newUser)

		})
	})

	// NOTE: Private Routes
	// Require Authentication
	// r.Group(func(r chi.Router) {
	// 	r.Use(AuthMiddleware)
	// 	r.Post("/manage", CreateAsset)
	// })

	// NOTE: v1APIRouter
	// v1APIRouter := chi.NewRouter()
	// router.Mount("/v1/api", v1APIRouter)

	return router
}
