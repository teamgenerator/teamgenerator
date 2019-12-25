package database

import (
	"github.com/teamgenerator/teamgenerator/server/db"
	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

// CommunityRepo is the database layer for interacting with the Community table
type CommunityRepo struct{}

func (r *CommunityRepo) GetCommunity(ID string) (*models.Community, error) {
	var community models.Community
	result := db.DB.First(&community, ID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &community, nil
}
