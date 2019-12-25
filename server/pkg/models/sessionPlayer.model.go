package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/teamgenerator/teamgenerator/server/db"
)

// The SessionPlayer object
type SessionPlayer struct {
	ID          int `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PlayerID    int `json:"PlayerID"`
	CommunityID int `json:"CommunityID"`
	Rating      int
	Form        int
	FormChange  int `json:"FormChange"`
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
	result := db.DB.First(&sessionPlayer, params["id"])
	if result.Error != nil {
		errMsg := fmt.Sprintf("SessionPlayer with id %s is not found", params["id"])
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&sessionPlayer)
}

// CreateSessionPlayer function to create a single SessionPlayer
func CreateSessionPlayer(w http.ResponseWriter, r *http.Request) {
	var sessionPlayer SessionPlayer
	json.NewDecoder(r.Body).Decode(&sessionPlayer)
	result := db.DB.Create(&sessionPlayer)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&sessionPlayer)
}

// UpdateSessionPlayer function to update an existing SessionPlayer
func UpdateSessionPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var sessionPlayer SessionPlayer
	result := db.DB.First(&sessionPlayer, params["id"])
	if result.Error != nil {
		errMsg := fmt.Sprintf("SessionPlayer with id %s is not found", params["id"])
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	json.NewDecoder(r.Body).Decode(&sessionPlayer)
	result = db.DB.Save(&sessionPlayer)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&sessionPlayer)
}

// DeleteSessionPlayer function to delete a single session-player joint table by ID
func DeleteSessionPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var sessionPlayer SessionPlayer
	result := db.DB.First(&sessionPlayer, params["id"])
	if result.Error != nil {
		errMsg := fmt.Sprintf("SessionPlayer with id %s is not found", params["id"])
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	if sessionPlayer.ID != 0 {
		result = db.DB.Delete(&sessionPlayer)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusBadRequest)
			return
		}
	}
	json.NewEncoder(w).Encode(&sessionPlayer)
}
