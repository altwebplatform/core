package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var sharedDB *gorm.DB

func Init(addr string) {
	globalDB, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal("failed to connect database on: " + addr, err)
	}

	// Migrate the schema
	globalDB.AutoMigrate(&Service{}, &Plugin{})
}
