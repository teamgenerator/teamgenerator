package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/teamgenerator/teamgenerator/server/db"
)

// User object
type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `gorm:"type:varchar(100);unique"`
	Name      string
	Password  string
}

// GetUsers function to return all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

// GetUser function to get a single User
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	db.DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(&user)
}

// CreateUser function to create a single User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.DB.Create(&user)
	json.NewEncoder(w).Encode(&user)
}

// UpdateUser function to update an existing User
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user User
	db.DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	db.DB.Save(&user)
	json.NewEncoder(w).Encode(&user)
}

// DeleteUser function to delete a single user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user User
	db.DB.First(&user, params["id"])
	if user.ID != 0 {
		db.DB.Delete(&user)
	}
	json.NewEncoder(w).Encode(&user)
}
