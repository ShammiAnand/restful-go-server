package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	config := &Config{}
	err := config.ReadEnv()
	if err != nil {
		log.Fatal(err)
	}

	portString := config.Port
	if portString == "" {
		log.Fatal("PORT is not found in environment")
	}

	fmt.Println("starting at PORT:", portString)

	r := mainRouter()

	stackOfMiddlewares := CreateStack(Logging)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", portString),
		Handler: stackOfMiddlewares(r),
	}

	serverError := server.ListenAndServe()
	if serverError != nil {
		log.Fatal(serverError)
	}

}

func mainRouter() http.Handler {

	router := http.NewServeMux()

  router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
      http.ServeFile(w, r, "./static/index.html")
  })

	router.HandleFunc("POST /users", CreateUserHandler)
	router.HandleFunc("GET /users", GetUsersHandler)
	router.HandleFunc("GET /users/{id}", GetUserWithIdHandler)
	router.HandleFunc("DELETE /users/{id}", DeleteUserWithIdHandler)

	return router
}
