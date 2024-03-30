package main

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"`         // Standard field for the primary key
	Name      string    `json:"name"`       // A regular string field
	Email     *string   `json:"email"`      // A pointer to a string, allowing for null values
	Age       uint8     `json:"age"`        // An unsigned 8-bit integer
	CreatedAt time.Time `json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedAt time.Time `json:"updated_at"` // Automatically managed by GORM for update time
	JobTitle  string    `json:"job_title"`
}
