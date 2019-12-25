package database

import (
	"github.com/teamgenerator/teamgenerator/server/db"
	"github.com/teamgenerator/teamgenerator/server/pkg/core"
	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

var (
	_ core.CommunityRepo = (*CommunityRepo)(nil)
)

// CommunityRepo is the database layer for interacting with the Community table
type CommunityRepo struct{}

// Get returns the communities with respect to the given filter
func (r *CommunityRepo) Get(filter core.CommunityFilter) ([]models.Community, error) {
	var communities []models.Community

	if len(filter.ID) == 1 {
		result := db.DB.First(&communities, filter.ID[0])
		if result.Error != nil {
			return nil, result.Error
		}
		if len(communities) < 1 {
			return nil, nil
		}
		return communities, nil
	}

	db.DB.Find(&communities)
	return communities, nil
}

// Create function to create a single community
func (r *CommunityRepo) Create(name, location string) (*models.Community, error) {
	community := models.Community{
		Name:     name,
		Location: location,
	}
	var createdCommunity models.Community
	result := db.DB.Create(&community).Scan(&createdCommunity)
	if result.Error != nil {
		return nil, result.Error
	}
	return &createdCommunity, nil
}

// Delete a single community given an ID
func (r *CommunityRepo) Delete(ID string) (*models.Community, error) {
	var community models.Community
	result := db.DB.First(&community, ID)
	if result.Error != nil {
		return nil, result.Error
	}

	if community.ID != 0 {
		result = db.DB.Delete(&community)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	return &community, nil
}
