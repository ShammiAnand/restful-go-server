package main

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Port        string `env:"PORT"`
	DB_PASSWORD string `env:"DB_PASSWORD"`
	DB_NAME     string `env:"DB_NAME"`
	DB_HOST     string `env:"DB_HOST"`
	DB_USER     string `env:"DB_USER"`
	DB_PORT     string `env:"DB_PORT"`
}

func (c *Config) ReadEnv() error {

	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	err := env.Parse(c)
	return err
}
