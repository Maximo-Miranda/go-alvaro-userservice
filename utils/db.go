package utils


import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

// Connect DB
func Connect() (*gorm.DB, error){

	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := gorm.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return db, nil
}