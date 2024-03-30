package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type GormDB struct {
// 	DB *gorm.DB
// }

// func (d *gorm.DB) Init() {
// 	d.DB, err := Database()
// 	if err != nil {
// 		log.Fatal("failed to initialise db")
// 	}
// }

func Database() (*gorm.DB, error) {

	loadEnvError := godotenv.Load(".env")
	if loadEnvError != nil {
		log.Fatal("couldn't load env")
	}

	var (
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_USER     = os.Getenv("DB_USER")
		DB_NAME     = os.Getenv("DB_NAME")
		DB_PORT     = os.Getenv("DB_PORT")
		DB_HOST     = os.Getenv("DB_HOST")
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Mumbai",
			DB_HOST,
			DB_USER,
			DB_PASSWORD,
			DB_NAME,
			DB_PORT,
		),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	migrationErr := db.AutoMigrate(&User{})
	if migrationErr != nil {
		log.Fatal("failed MIGRATION")
	}

	return db, err
}
