package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teamgenerator/teamgenerator/server/db"
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
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&community)
}

// CreateCommunity function to create a single community
func (h *CommunityHandler) CreateCommunity(w http.ResponseWriter, r *http.Request) {
	var community models.Community
	json.NewDecoder(r.Body).Decode(&community)
	result := db.DB.Create(&community)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&community)
}

// UpdateCommunity function to update an existing community
func (h *CommunityHandler) UpdateCommunity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var community models.Community

	result := db.DB.First(&community, params["id"])
	if result.Error != nil {
		errMsg := fmt.Sprintf("Community with id %s is not found", params["id"])
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	json.NewDecoder(r.Body).Decode(&community)
	result = db.DB.Save(&community)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&community)
}

// DeleteCommunity function to delete a single communtiy by ID
func (h *CommunityHandler) DeleteCommunity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var community models.Community
	result := db.DB.First(&community, params["id"])
	if result.Error != nil {
		errMsg := fmt.Sprintf("Community with id %s is not found", params["id"])
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	if community.ID != 0 {
		result = db.DB.Delete(&community)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusBadRequest)
			return
		}
	}
	json.NewEncoder(w).Encode(&community)
}
