package models

import (
	"time"
)

// The Session object
type Session struct {
	ID          uint `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsActive    bool `json:"IsActive"`
	CommunityID int  `json:"CommunityID"`
}
