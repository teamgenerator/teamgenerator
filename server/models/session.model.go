package models

import (
	"encoding/json"
	"fmt"
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
	IsActive    bool `json:"IsActive"`
	CommunityID int  `json:"CommunityID"`
}

// GetSessions function to return all sessions
func GetSessions(w http.ResponseWriter, r *http.Request) {
	var sessions []Session
	db.DB.Find(&sessions)
	json.NewEncoder(w).Encode(&sessions)
}

// GetSession function to get a single Session
func GetSession(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var session Session
	result := db.DB.First(&session, params["id"])
	if result.Error != nil {
		errMsg := fmt.Sprintf("Session with id %s is not found", params["id"])
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&session)
}

// CreateSession function to create a single Session
func CreateSession(w http.ResponseWriter, r *http.Request) {
	var session Session
	json.NewDecoder(r.Body).Decode(&session)
	result := db.DB.Create(&session)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&session)
}

// UpdateSession function to update an existing session
func UpdateSession(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var session Session
	result := db.DB.First(&session, params["id"])
	if result.Error != nil {
		errMsg := fmt.Sprintf("Session with id %s is not found", params["id"])
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	json.NewDecoder(r.Body).Decode(&session)
	result = db.DB.Save(&session)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&session)
}

// DeleteSession function to delete a single session by ID
func DeleteSession(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var session Session
	result := db.DB.First(&session, params["id"])
	if result.Error != nil {
		errMsg := fmt.Sprintf("Session with id %s is not found", params["id"])
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	if session.ID != 0 {
		result = db.DB.Delete(&session)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusBadRequest)
			return
		}
	}
	json.NewEncoder(w).Encode(&session)
}
