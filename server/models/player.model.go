package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/teamgenerator/teamgenerator/server/db"
)

// The Player Object
type Player struct {
	ID          uint `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Ratings     int
	Form        int
	CommunityID int `json:"community_id"`
}

// GetPlayers function to return all players
func GetPlayers(w http.ResponseWriter, r *http.Request) {
	var players []Player
	db.DB.Find(&players)
	json.NewEncoder(w).Encode(&players)
}

// GetPlayer function to get a single Player
func GetPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var player Player
	db.DB.First(&player, params["id"])
	json.NewEncoder(w).Encode(&player)
}

// CreatePlayer function to create a single Player
func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var player Player
	json.NewDecoder(r.Body).Decode(&player)
	db.DB.Create(&player)
	json.NewEncoder(w).Encode(&player)
}

// UpdatePlayer function to update an existing players
func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var player Player
	db.DB.First(&player, params["id"])
	json.NewDecoder(r.Body).Decode(&player)
	db.DB.Save(&player)
	json.NewEncoder(w).Encode(&player)
}

// DeletePlayer function to delete a single player by ID
func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var player Player
	db.DB.First(&player, params["id"])
	if player.ID != 0 {
		db.DB.Delete(&player)
	}
	json.NewEncoder(w).Encode(&player)
}
