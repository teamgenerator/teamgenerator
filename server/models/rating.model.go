package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teamgenerator/teamgenerator/server/db"
)

// Rating object
type Rating struct {
	ID         uint `gorm:"primary_key"`
	RatingGain uint `json:"RatingGain"`
	PlayerID   uint `json:"PlayerID"`
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
	result := db.DB.First(&rating, params["id"])
	if result.Error != nil {
		errMsg := fmt.Sprintf("Rating with id %s is not found", params["id"])
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&rating)
}

// CreateRating function to create a single rating
func CreateRating(w http.ResponseWriter, r *http.Request) {
	var rating Rating
	json.NewDecoder(r.Body).Decode(&rating)
	result := db.DB.Create(&rating)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&rating)
}

// UpdateRating function to update an existing rating
func UpdateRating(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var rating Rating
	result := db.DB.First(&rating, params["id"])
	if result.Error != nil {
		errMsg := fmt.Sprintf("Rating with id %s is not found", params["id"])
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	json.NewDecoder(r.Body).Decode(&rating)

	result = db.DB.Save(&rating)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&rating)
}

// DeleteRating function to delete a single communtiy by ID
func DeleteRating(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var rating Rating
	result := db.DB.First(&rating, params["id"])
	if result.Error != nil {
		errMsg := fmt.Sprintf("Rating with id %s is not found", params["id"])
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	if rating.ID != 0 {
		result = db.DB.Delete(&rating)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusBadRequest)
			return
		}
	}
	json.NewEncoder(w).Encode(&rating)
}
