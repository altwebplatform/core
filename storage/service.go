package storage

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Service struct {
	gorm.Model
	ID        uint64    `gorm:"AUTO_INCREMENT"`
	Name      string
	Type      string
	CreatedAt time.Time
}

//db.Create(&db.Service{
//Name:      "test22",
//CreatedAt: time.Now(),
//})
//
//var service db.Service
//db.Find(&service)
//
//fmt.Println(service.Name)

