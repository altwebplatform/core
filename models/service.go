package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

func SetupDB(addr string) *gorm.DB {
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Service{}, &Plugin{})

	return db
}

type Service struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	Type      string    `db:"type"`
	CreatedAt time.Time `db:"created_at"`
}
