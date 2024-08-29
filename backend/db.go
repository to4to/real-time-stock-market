package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBConnection establishes a connection to the database using the provided environment configuration.
// It returns a pointer to a gorm.DB instance.
func DBConnection(env *Env) *gorm.DB {
	// Construct the connection URI using the environment variables
	uri := fmt.Sprintf(
		`
		host=%s user=%s dbname=%s password=%s sslmode=%s port=5432,
		`,
		env.DB_HOST, env.DB_USER, env.DB_NAME, env.DB_SSLMODE,
	)

	// Open a connection to the database using gorm with the constructed URI
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	// Check for any errors that occurred during connection establishment
	if err != nil {
		log.Fatal("Unable to connect to database: %e", err)
	}

	fmt.Println("Connected To Database")

	// Automatically migrate the schema for the Candle model
	if err := db.AutoMigrate(&Candle{}); err != nil {
		log.Fatal("Unable to migrate: %e", err)
	}

	return db
}
