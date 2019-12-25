package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teamgenerator/teamgenerator/server/pkg/core"
	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

// PlayerHandler handles the HTTP layer for players
type PlayerHandler struct {
	PlayerCore core.PlayerCore
}

// GetPlayers function to return all players
func (h *PlayerHandler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := h.PlayerCore.GetPlayers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&players)
}

// GetPlayer function to get a single player
func (h *PlayerHandler) GetPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	player, err := h.PlayerCore.GetPlayer(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&player)
}

// CreatePlayer function to create a single player
func (h *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	json.NewDecoder(r.Body).Decode(&player)
	createdPlayer, err := h.PlayerCore.CreatePlayer(player.Name, player.Ratings, player.Form, player.CommunityID)
	switch {
	case err == nil:
		json.NewEncoder(w).Encode(&createdPlayer)
	case errors.Is(err, core.ErrInvalidInputParams), errors.Is(err, core.ErrCommunityNotFound):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeletePlayer function to delete a single communtiy by ID
func (h *PlayerHandler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	createdPlayer, err := h.PlayerCore.DeletePlayer(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&createdPlayer)
}
