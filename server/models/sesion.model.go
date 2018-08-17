////
// author: Nico Alimin (nicoalimin@hotmail.com)
// date: Friday, 17th August 2018 12:48:23 am
// lastModifiedBy: Nico Alimin (nicoalimin@hotmail.com)
// lastModifiedTime: Friday, 17th August 2018 12:48:24 am
//
// DESCRIPTION
//
// copyright (c) 2018 Nico Alimin
////

package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/teamgenerator/teamgenerator/server/db"
)

// The Session object
type Session struct {
	ID          uint `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsActive    bool
	CommunityID int
}

// GetSession function to get a single Session
func GetSession(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var session Session
	db.DB.First(&session, params["id"])
	json.NewEncoder(w).Encode(&session)
}

// CreateSession function to create a single Session
func CreateSession(w http.ResponseWriter, r *http.Request) {
	var session Session
	json.NewDecoder(r.Body).Decode(&session)
	db.DB.Create(&session)
	json.NewEncoder(w).Encode(&session)
}

// UpdateSession function to update an existing session
func UpdateSession(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var session Session
	db.DB.First(&session, params["id"])
	json.NewDecoder(r.Body).Decode(&session)
	db.DB.Save(&session)
	json.NewEncoder(w).Encode(&session)
}

// DeleteSession function to delete a single session by ID
func DeleteSession(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var session Session
	db.DB.First(&session, params["id"])
	if session.ID != 0 {
		db.DB.Delete(&session)
	}
	json.NewEncoder(w).Encode(&session)
}
