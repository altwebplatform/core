package storage

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Plugin struct {
	gorm.Model
	ID            uint64    `gorm:"AUTO_INCREMENT"`
	Name          string
	PodDefinition string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}
