package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	payload := UserParams{}
	err := decoder.Decode(&payload)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Error Parsing JSON: %v", err))
		return
	}

	newUser := User{
		Name:     payload.Name,
		Age:      payload.Age,
		Email:    payload.Email,
		JobTitle: payload.JobTitle,
	}

	db, err := Database()
	if err != nil {
		log.Fatal("error while connecting to db")
	}
	result := db.Create(&newUser)

	if result.Error != nil {
		respondWithError(w, 403, fmt.Sprintf("failed to create: %v", result.Error))
	}

	respondWithJson(w, 201, newUser)
}

func GetUserWithIdHandler(w http.ResponseWriter, r *http.Request) {

	userId, conversionError := strconv.Atoi(r.PathValue("id"))
	if conversionError != nil {
		respondWithError(w, 403, fmt.Sprintf("id should be a positive integer: %v", conversionError))
		return
	}

	db, err := Database()
	if err != nil {
		log.Fatal("error while connecting to db")
	}

	user := User{ID: uint(userId)}
	result := db.Find(&user)

	if result.Error != nil {
		respondWithError(w, 403, fmt.Sprintf("failed to get users: %v", result.Error))
		return
	}

	respondWithJson(w, 200, user)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	db, err := Database()
	if err != nil {
		log.Fatal("error while connecting to db")
	}

	users := []User{}
	result := db.Find(&users)

	if result.Error != nil {
		respondWithError(w, 403, fmt.Sprintf("failed to get users: %v", result.Error))
	}

	respondWithJson(w, 200, users)

}

func DeleteUserWithIdHandler(w http.ResponseWriter, r *http.Request) {

	userId, conversionError := strconv.Atoi(r.PathValue("id"))
	if conversionError != nil {
		respondWithError(w, 403, fmt.Sprintf("id should be a positive integer: %v", conversionError))
		return
	}

	db, err := Database()
	if err != nil {
		log.Fatal("error while connecting to db")
	}

	user := User{ID: uint(userId)}
	result := db.Delete(&user)

	if result.Error != nil {
		respondWithError(w, 403, fmt.Sprintf("failed to delete user: %v", result.Error))
		return
	}

	respondWithJson(w, 200, struct {
		ID int `json:"id"`
	}{
		ID: userId,
	})
}
