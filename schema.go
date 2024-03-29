package main

import (
	"time"
)

type User struct {
	ID        uint      // Standard field for the primary key
	Name      string    // A regular string field
	Email     *string   // A pointer to a string, allowing for null values
	Age       uint8     // An unsigned 8-bit integer
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time // Automatically managed by GORM for update time
	JobTitle  string
}
