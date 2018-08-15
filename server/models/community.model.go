////
// author: Nico Alimin (nicoalimin@gmail.com)
// date: Sunday, 5th August 2018 1:24:59 am
// lastModifiedBy: Nico Alimin (nicoalimin@gmail.com)
// lastModifiedTime: Sunday, 5th August 2018 1:30:28 am
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

// Community object
type Community struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);unique"`
	Location string
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
