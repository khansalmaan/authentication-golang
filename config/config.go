package config

import (
	"os"

	"github.com/joho/godotenv"
)

// ConfigStructure holds the configuration settings
type ConfigStructure struct {
	// server
	Port string

	// database
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

// Config is a globally accessible instance of ConfigStructure
var Config *ConfigStructure

// Init initializes the Config with default values
func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	Config = &ConfigStructure{
		// server
		Port: os.Getenv("PORT"),
		// database
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
	}
}
