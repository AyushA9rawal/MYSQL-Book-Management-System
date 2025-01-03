package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	// Simplified DSN connection string
	dsn := "root:Ayush123@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"

	// Open database connection
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	} else {
		log.Println("Database connection successful!")
	}
}

func GetDB() *gorm.DB {
	return db
}
