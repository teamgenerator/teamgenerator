package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teamgenerator/teamgenerator/server/pkg/core"
	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

// CommunityHandler handles the HTTP layer for communities
type CommunityHandler struct {
	CommunityCore core.CommunityCore
}

// GetCommunities function to return all communities
func (h *CommunityHandler) GetCommunities(w http.ResponseWriter, r *http.Request) {
	communities, err := h.CommunityCore.GetCommunities()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&communities)
}

// GetCommunity function to get a single community
func (h *CommunityHandler) GetCommunity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	community, err := h.CommunityCore.GetCommunity(params["id"])
	switch {
	case err == nil:
		json.NewEncoder(w).Encode(&community)
	case errors.Is(err, core.ErrCommunityNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// CreateCommunity function to create a single community
func (h *CommunityHandler) CreateCommunity(w http.ResponseWriter, r *http.Request) {
	var community models.Community
	json.NewDecoder(r.Body).Decode(&community)
	createdCommunity, err := h.CommunityCore.CreateCommunity(community.Name, community.Location)
	switch {
	case err == nil:
		json.NewEncoder(w).Encode(&createdCommunity)
	case errors.Is(err, core.ErrInvalidInputParams):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteCommunity function to delete a single communtiy by ID
func (h *CommunityHandler) DeleteCommunity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	createdCommunity, err := h.CommunityCore.DeleteCommunity(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&createdCommunity)
}
