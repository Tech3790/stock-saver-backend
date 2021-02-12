package functions

import (
	"fmt"

	"github.com/Tech3790/stock-saver-backend/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=abdullah dbname=stock-saver port=5432 sslmode=disable TimeZone=US/Central"
var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

// IntitMigration starts the database
func IntitMigration() {
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	db.AutoMigrate(&models.User{})
}
