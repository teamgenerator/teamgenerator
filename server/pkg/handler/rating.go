package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teamgenerator/teamgenerator/server/pkg/core"
	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

// RatingHandler handles the HTTP layer for rating
type RatingHandler struct {
	RatingCore core.RatingCore
}

// GetRatings function to return all rating
func (h *RatingHandler) GetRatings(w http.ResponseWriter, r *http.Request) {
	rating, err := h.RatingCore.GetRatings()
	switch {
	case err == nil:
		json.NewEncoder(w).Encode(&rating)
	case errors.Is(err, core.ErrRatingNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetRating function to get a single rating
func (h *RatingHandler) GetRating(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rating, err := h.RatingCore.GetRating(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&rating)
}

// CreateRating function to create a single rating
func (h *RatingHandler) CreateRating(w http.ResponseWriter, r *http.Request) {
	var rating models.Rating
	json.NewDecoder(r.Body).Decode(&rating)
	createdRating, err := h.RatingCore.CreateRating(rating.RatingGain, rating.PlayerID)
	switch {
	case err == nil:
		json.NewEncoder(w).Encode(&createdRating)
	case errors.Is(err, core.ErrInvalidInputParams), errors.Is(err, core.ErrPlayerNotFound):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteRating function to delete a single communtiy by ID
func (h *RatingHandler) DeleteRating(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	createdRating, err := h.RatingCore.DeleteRating(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&createdRating)
}
