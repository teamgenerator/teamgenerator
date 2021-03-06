package models

import (
	"time"
)

// Community object
type Community struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"type:varchar(100);unique"`
	Location  string
}
