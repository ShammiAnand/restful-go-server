package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
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

	addr := fmt.Sprintf(":%v", portString)
	r := mainRouter()
	serverError := http.ListenAndServe(addr, logging(r))

	if serverError != nil {
		log.Fatal(serverError)
	}
}

func mainRouter() http.Handler {

	router := http.NewServeMux()

	_, err := Database()
	if err != nil {
		log.Fatal("couldn't connect to DB")
	}
	fmt.Println("connected to DB")

	router.HandleFunc("POST /users", CreateUserHandler)
	router.HandleFunc("GET /users", GetUsersHandler)
	router.HandleFunc("GET /users/{id}", GetUserWithIdHandler)
	router.HandleFunc("DELETE /users/{id}", DeleteUserWithIdHandler)

	return router
}
