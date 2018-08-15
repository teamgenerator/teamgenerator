////
// author: Nico Alimin (nicoalimin@hotmail.com)
// date: Tuesday, 14th August 2018 12:35:01 am
// lastModifiedBy: Nico Alimin (nicoalimin@hotmail.com)
// lastModifiedTime: Tuesday, 14th August 2018 12:35:01 am
//
// DESCRIPTION
//
// copyright (c) 2018 Nico Alimin
////

package models

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/teamgenerator/teamgenerator/server/db"
)

// The Player Object
type Player struct {
	gorm.Model
	Name        string
	Ratings     int
	CommunityID int
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

// DeletePlayer function to delete a single communtiy by ID
func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var player Player
	db.DB.First(&player, params["id"])
	if player.ID != 0 {
		db.DB.Delete(&player)
	}
	json.NewEncoder(w).Encode(&player)
}
