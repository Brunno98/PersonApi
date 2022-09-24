package model

import (
	"time"
)

type Person struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	Name      string
	Birthdate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
