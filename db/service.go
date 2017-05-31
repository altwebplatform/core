package db

import (
	"time"
)

type Service struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	Type      string    `db:"type"`
	CreatedAt time.Time `db:"created_at"`
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
