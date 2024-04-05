package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

type Middleware func(http.Handler) http.Handler

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapped, r)
		log.Println(wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
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

	stackOfMiddlewares := CreateStack(logging)

	server := http.Server{
		Addr:    addr,
		Handler: stackOfMiddlewares(r),
	}

	serverError := server.ListenAndServe()
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
