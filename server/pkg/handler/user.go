package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teamgenerator/teamgenerator/server/pkg/core"
	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

// UserHandler handles the HTTP layer for user
type UserHandler struct {
	UserCore core.UserCore
}

// GetUsers function to return all user
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	user, err := h.UserCore.GetUsers()
	switch {
	case err == nil:
		json.NewEncoder(w).Encode(&user)
	case errors.Is(err, core.ErrUserNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetUser function to get a single user
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := h.UserCore.GetUser(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&user)
}

// CreateUser function to create a single user
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	createdUser, err := h.UserCore.CreateUser(user.Name, user.Username, user.Password)
	switch {
	case err == nil:
		json.NewEncoder(w).Encode(&createdUser)
	case errors.Is(err, core.ErrInvalidInputParams), errors.Is(err, core.ErrCommunityNotFound):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteUser function to delete a single communtiy by ID
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	createdUser, err := h.UserCore.DeleteUser(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&createdUser)
}
