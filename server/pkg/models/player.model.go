package models

import (
	"time"
)

// The Player Object
type Player struct {
	ID          int `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Ratings     int
	Form        int
	CommunityID int `json:"CommunityID"`
}
