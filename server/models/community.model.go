package models

import (
	"encoding/json"
	"fmt"
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
	result := db.DB.First(&community, params["id"])
	if result.Error != nil {
		errMsg := fmt.Sprintf("Community with id %s is not found", params["id"])
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&community)
}

// CreateCommunity function to create a single community
func CreateCommunity(w http.ResponseWriter, r *http.Request) {
	var community Community
	json.NewDecoder(r.Body).Decode(&community)
	result := db.DB.Create(&community)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&community)
}

// UpdateCommunity function to update an existing community
func UpdateCommunity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var community Community

	result := db.DB.First(&community, params["id"])
	if result.Error != nil {
		errMsg := fmt.Sprintf("Community with id %s is not found", params["id"])
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	json.NewDecoder(r.Body).Decode(&community)
	result = db.DB.Save(&community)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&community)
}

// DeleteCommunity function to delete a single communtiy by ID
func DeleteCommunity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var community Community
	result := db.DB.First(&community, params["id"])
	if result.Error != nil {
		errMsg := fmt.Sprintf("Community with id %s is not found", params["id"])
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	if community.ID != 0 {
		result = db.DB.Delete(&community)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusBadRequest)
			return
		}
	}
	json.NewEncoder(w).Encode(&community)
}
