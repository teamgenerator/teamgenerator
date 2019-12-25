package database

import (
	"fmt"
	"github.com/teamgenerator/teamgenerator/server/db"
	"github.com/teamgenerator/teamgenerator/server/pkg/core"
	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

var (
	_ core.CommunityRepo = (*CommunityRepo)(nil)
)

// CommunityRepo is the database layer for interacting with the Community table
type CommunityRepo struct{}

func (r *CommunityRepo) GetCommunities(filter core.CommunityFilter) ([]models.Community, error) {
	var communities []models.Community

	if len(filter.ID) == 1 {
		result := db.DB.First(&communities, filter.ID[0])
		if result.Error != nil {
			return nil, result.Error
		}
		return communities, nil
	}

	db.DB.Find(&communities)
	return communities, nil
}
