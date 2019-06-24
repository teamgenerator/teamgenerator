package models

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teamgenerator/teamgenerator/server/db"
)

// Rating object
type Rating struct {
	ID         uint `gorm:"primary_key"`
	RatingGain uint
	PlayerID   uint
}

// GetRatings function to return all ratings
func GetRatings(w http.ResponseWriter, r *http.Request) {
	var ratings []Rating
	db.DB.Find(&ratings)
	json.NewEncoder(w).Encode(&ratings)
}

// GetRating function to get a single rating
func GetRating(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var rating Rating
	db.DB.First(&rating, params["id"])
	json.NewEncoder(w).Encode(&rating)
}

// CreateRating function to create a single rating
func CreateRating(w http.ResponseWriter, r *http.Request) {
	var rating Rating
	json.NewDecoder(r.Body).Decode(&rating)
	db.DB.Create(&rating)
	json.NewEncoder(w).Encode(&rating)
}

// UpdateRating function to update an existing rating
func UpdateRating(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var rating Rating
	db.DB.First(&rating, params["id"])
	json.NewDecoder(r.Body).Decode(&rating)
	db.DB.Save(&rating)
	json.NewEncoder(w).Encode(&rating)
}

// DeleteRating function to delete a single communtiy by ID
func DeleteRating(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var rating Rating
	db.DB.First(&rating, params["id"])
	if rating.ID != 0 {
		db.DB.Delete(&rating)
	}
	json.NewEncoder(w).Encode(&rating)
}
