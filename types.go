package main

type UserParams struct {
	Name     string  `json:"name"`
	Email    *string `json:"email"`
	JobTitle string  `json:"job_title"`
	Age      uint8   `json:"age"`
}

type Message struct {
	Message string `json:"message"`
}

