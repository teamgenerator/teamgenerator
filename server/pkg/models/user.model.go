package models

import (
	"time"
)

// User object
type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `gorm:"type:varchar(100);unique"`
	Name      string
	Password  string
}
