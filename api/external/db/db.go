package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-rest/config"
)

var Database *gorm.DB

func Init() {
	var err error
	fmt.Println("Connecting to the database...")
	host := config.Config.DBHost
	username := config.Config.DBUser
	password := config.Config.DBPassword
	databaseName := config.Config.DBName
	port := config.Config.DBPort

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databaseName, port)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}
