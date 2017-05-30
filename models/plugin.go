package models

import "time"

type Plugin struct {
	ID            uint64    `db:"id"`
	Name          string    `db:"name"`
	PodDefinition string    `db:"type"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
	DeletedAt     time.Time `db:"deleted_at"`
}
