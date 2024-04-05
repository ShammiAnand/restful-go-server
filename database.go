package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type config struct {
	Port        string `env:"PORT"`
	DB_PASSWORD string `env:"DB_PASSWORD"`
	DB_NAME     string `env:"DB_NAME"`
	DB_HOST     string `env:"DB_HOST"`
	DB_USER     string `env:"DB_USER"`
	DB_PORT     string `env:"DB_PORT"`
}

func Database() (*gorm.DB, error) {

	loadEnvError := godotenv.Load(".env")
	if loadEnvError != nil {
		log.Fatal("couldn't load env")
	}

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("couldn't parse ENV")
	}

	// fmt.Printf("%+v\n", cfg)
	// fmt.Println("ENV parsed")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Mumbai",
			cfg.DB_HOST,
			cfg.DB_USER,
			cfg.DB_PASSWORD,
			cfg.DB_NAME,
			cfg.DB_PORT,
		),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	migrationErr := db.AutoMigrate(&User{})
	if migrationErr != nil {
		log.Fatal("failed MIGRATION")
	}

	return db, err
}
