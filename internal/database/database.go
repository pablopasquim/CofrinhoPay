package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pablopasquim/CofrinhoPay/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // var global

func Connect() { // db config
	godotenv.Load()

	connection := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"),
	)

	database, err := gorm.Open(postgres.Open(connection)) // GORM open postgres connection
	if err != nil {
		log.Fatal("❌ Error connection in database", err)
	}

	DB = database

	err = DB.AutoMigrate(
		&models.Account{},
		&models.Transaction{},
		&models.Category{},
		&models.User{},
	)
	if err != nil {
		log.Fatal("❌ Error auto-migrate in database", err)
	}

	log.Println("✅ Connected in the database")

}
