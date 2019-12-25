package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teamgenerator/teamgenerator/server/pkg/core"
	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

// SessionPlayerHandler handles the HTTP layer for sessions
type SessionPlayerHandler struct {
	SessionPlayerCore core.SessionPlayerCore
}

// GetSessionPlayers function to return all sessions
func (h *SessionPlayerHandler) GetSessionPlayers(w http.ResponseWriter, r *http.Request) {
	sessions, err := h.SessionPlayerCore.GetSessionPlayers()
	switch {
	case err == nil:
		json.NewEncoder(w).Encode(&sessions)
	case errors.Is(err, core.ErrSessionPlayerNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetSessionPlayer function to get a single session
func (h *SessionPlayerHandler) GetSessionPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	session, err := h.SessionPlayerCore.GetSessionPlayer(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&session)
}

// CreateSessionPlayer function to create a single session
func (h *SessionPlayerHandler) CreateSessionPlayer(w http.ResponseWriter, r *http.Request) {
	var session models.SessionPlayer
	json.NewDecoder(r.Body).Decode(&session)
	createdSessionPlayer, err := h.SessionPlayerCore.CreateSessionPlayer(session.PlayerID, session.CommunityID, session.Rating, session.Form, session.FormChange)
	switch {
	case err == nil:
		json.NewEncoder(w).Encode(&createdSessionPlayer)
	case errors.Is(err, core.ErrInvalidInputParams), errors.Is(err, core.ErrCommunityNotFound), errors.Is(err, core.ErrPlayerNotFound):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteSessionPlayer function to delete a single communtiy by ID
func (h *SessionPlayerHandler) DeleteSessionPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	createdSessionPlayer, err := h.SessionPlayerCore.DeleteSessionPlayer(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&createdSessionPlayer)
}
