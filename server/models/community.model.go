package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/teamgenerator/teamgenerator/server/db"
)

// Community object
type Community struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"type:varchar(100);unique"`
	Location  string
}

// GetCommunities function to return all communities
func GetCommunities(w http.ResponseWriter, r *http.Request) {
	var communities []Community
	db.DB.Find(&communities)
	json.NewEncoder(w).Encode(&communities)
}

// GetCommunity function to get a single community
func GetCommunity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var community Community
	db.DB.First(&community, params["id"])
	json.NewEncoder(w).Encode(&community)
}

// CreateCommunity function to create a single community
func CreateCommunity(w http.ResponseWriter, r *http.Request) {
	var community Community
	json.NewDecoder(r.Body).Decode(&community)
	db.DB.Create(&community)
	json.NewEncoder(w).Encode(&community)
}

// UpdateCommunity function to update an existing community
func UpdateCommunity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var community Community
	db.DB.First(&community, params["id"])
	json.NewDecoder(r.Body).Decode(&community)
	db.DB.Save(&community)
	json.NewEncoder(w).Encode(&community)
}

// DeleteCommunity function to delete a single communtiy by ID
func DeleteCommunity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var community Community
	db.DB.First(&community, params["id"])
	if community.ID != 0 {
		db.DB.Delete(&community)
	}
	json.NewEncoder(w).Encode(&community)
}
