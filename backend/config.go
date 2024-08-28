package main

import (
	"log"

	"github.com/joho/godotenv"
)

type Env struct {
	ServerPort string `env:"SERVER_PORT,required"`
	API_KEY    string `env:"API_KEY,required"`

	//postgreSQL Connection
	DB_HOST     string `env:"DB_HOST,required"`
	DB_NAME     string `env:"DB_NAMEt,required"`
	DB_USER     string `env:"DB_USER,required"`
	DB_PASSWORD string `env:"DB_PASSWORD,required"`
	DB_SSLMODE  string `env:"DB_SSLMODE,required"`
}

func EnvConfig() *Env {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load .env file: %e", err)
	}


	config:=&Env{}


	

}
