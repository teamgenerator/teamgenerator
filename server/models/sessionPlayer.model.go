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

// GetSessionPlayers function to return all ratings
func GetSessionPlayers(w http.ResponseWriter, r *http.Request) {
	var sessionPlayers []SessionPlayer
	db.DB.Find(&sessionPlayers)
	json.NewEncoder(w).Encode(&sessionPlayers)
}

// GetSessionPlayer function to get a single SessionPlayer
func GetSessionPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var sessionPlayer SessionPlayer
	db.DB.First(&sessionPlayer, params["id"])
	json.NewEncoder(w).Encode(&sessionPlayer)
}

// CreateSessionPlayer function to create a single SessionPlayer
func CreateSessionPlayer(w http.ResponseWriter, r *http.Request) {
	var sessionPlayer SessionPlayer
	json.NewDecoder(r.Body).Decode(&sessionPlayer)
	db.DB.Create(&sessionPlayer)
	json.NewEncoder(w).Encode(&sessionPlayer)
}

// UpdateSessionPlayer function to update an existing SessionPlayer
func UpdateSessionPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var sessionPlayer SessionPlayer
	db.DB.First(&sessionPlayer, params["id"])
	json.NewDecoder(r.Body).Decode(&sessionPlayer)
	db.DB.Save(&sessionPlayer)
	json.NewEncoder(w).Encode(&sessionPlayer)
}

// DeleteSessionPlayer function to delete a single session-player joint table by ID
func DeleteSessionPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var sessionPlayer SessionPlayer
	db.DB.First(&sessionPlayer, params["id"])
	if sessionPlayer.ID != 0 {
		db.DB.Delete(&sessionPlayer)
	}
	json.NewEncoder(w).Encode(&sessionPlayer)
}
