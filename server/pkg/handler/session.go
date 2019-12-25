package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teamgenerator/teamgenerator/server/pkg/core"
	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

// SessionHandler handles the HTTP layer for sessions
type SessionHandler struct {
	SessionCore core.SessionCore
}

// GetSessions function to return all sessions
func (h *SessionHandler) GetSessions(w http.ResponseWriter, r *http.Request) {
	sessions, err := h.SessionCore.GetSessions()
	switch {
	case err == nil:
		json.NewEncoder(w).Encode(&sessions)
	case errors.Is(err, core.ErrSessionsNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetSession function to get a single session
func (h *SessionHandler) GetSession(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	session, err := h.SessionCore.GetSession(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&session)
}

// CreateSession function to create a single session
func (h *SessionHandler) CreateSession(w http.ResponseWriter, r *http.Request) {
	var session models.Session
	json.NewDecoder(r.Body).Decode(&session)
	createdSession, err := h.SessionCore.CreateSession(session.IsActive, session.CommunityID)
	switch {
	case err == nil:
		json.NewEncoder(w).Encode(&createdSession)
	case errors.Is(err, core.ErrInvalidInputParams), errors.Is(err, core.ErrCommunityNotFound):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateSession changes the isActive state of a session
func (h *SessionHandler) UpdateSession(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var session models.Session
	json.NewDecoder(r.Body).Decode(&session)
	createdSession, err := h.SessionCore.UpdateSession(params["id"], session.IsActive)
	switch {
	case err == nil:
		json.NewEncoder(w).Encode(&createdSession)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteSession function to delete a single communtiy by ID
func (h *SessionHandler) DeleteSession(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	createdSession, err := h.SessionCore.DeleteSession(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&createdSession)
}
