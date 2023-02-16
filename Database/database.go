package database

import (
	"BulindingGoRestAPI/config"
	"BulindingGoRestAPI/models"
	"fmt"
	"log"
	"os"

	// "gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
var err error

type DbInstance struct {
	Db *gorm.DB
	// FirstName string `json:"firstname"`
	// LastName  string  `json:"lastname"`
	// Email     string  `json:"email"`
}

var Database DbInstance

func ConnectDb() {
	config, err := config.LoadConfig("/home/tito/Desktop/java/config")
	if err != nil {
			log.Fatal("cannot load config:", err)
	}

	// db , err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Africa/Nairobi", config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to open database")
		os.Exit(2)
	}

	if err != nil {
		log.Panic("Failed to open database")
		os.Exit(2)
	}
	log.Println("Database Connected succefully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	// add migrations
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.CreditCard{})
	Database = DbInstance{Db: db}
}
