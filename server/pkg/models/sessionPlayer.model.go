package models

import (
	"time"
)

// The SessionPlayer object
type SessionPlayer struct {
	ID         int `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	PlayerID   int `json:"PlayerID"`
	SessionID  int `json:"SessionID"`
	Rating     int
	Form       int
	FormChange int `json:"FormChange"`
}
