package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname string `gorm:"not null;unique_index"`
	Password string `gorm:"not null"`
}

type Vacation struct {
	gorm.Model
	Name      string    `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
}
