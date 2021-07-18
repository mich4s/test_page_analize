package repositories

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var connection *gorm.DB

func Init() {
	var err error

	connection, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	err = connection.AutoMigrate(&Page{})

	if err != nil {
		panic("failed to migrate database")
	}
}
