package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var globalDB *gorm.DB

func SetupDB(addr string) {
	globalDB, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Migrate the schema
	globalDB.AutoMigrate(&Service{}, &Plugin{})
}
