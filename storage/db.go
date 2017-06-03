package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"github.com/altwebplatform/core/config"
)

var sharedDB *gorm.DB

func SharedDB() *gorm.DB {
	if sharedDB == nil {
		Init(config.DB_CONNECT)
	}
	return sharedDB
}

func Init(addr string) {
	var err error
	sharedDB, err = gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal("failed to connect database on: " + addr, err)
	}

	// Migrate the schema
	sharedDB = sharedDB.AutoMigrate(&Service{}, &Plugin{})
}

func Close() {
	if sharedDB != nil {
		sharedDB.Close()
	}
}