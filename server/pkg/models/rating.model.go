package models

// Rating object
type Rating struct {
	ID         int `gorm:"primary_key"`
	RatingGain int `json:"RatingGain"`
	PlayerID   int `json:"PlayerID"`
	SessionID  int `json:"SessionID"`
}
