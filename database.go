package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbDriver struct {
	db *gorm.DB
}

func (d *DbDriver) Database() error {

	cfg := &Config{}
	envError := cfg.ReadEnv()
	if envError != nil {
		log.Fatal(envError)
	}

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

	if err != nil {
		return err
	}

	migrationErr := db.AutoMigrate(&User{})
	if migrationErr != nil {
		return err
	}

	d.db = db
	return nil
}
