package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/teamgenerator/teamgenerator/server/db"
)

// The SessionPlayer object
type SessionPlayer struct {
	ID          uint `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PlayerID    int
	CommunityID int
	Rating      int
	Form        int
	FormChange  int
}

// GetSessionPlayer function to get a single SessionPlayer
func GetSessionPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var sessionPlayer SessionPlayer
	db.DB.First(&sessionPlayer, params["id"])
	json.NewEncoder(w).Encode(&sessionPlayer)
}
