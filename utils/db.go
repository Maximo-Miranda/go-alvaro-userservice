package utils


import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connect
func Connect() (*gorm.DB, error){
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=elmisticoDB sslmode=disable password=")
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return db, nil
}