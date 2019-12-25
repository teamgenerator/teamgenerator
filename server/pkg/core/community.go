package core

import (
	"time"

	"github.com/teamgenerator/teamgenerator/server/pkg/database"

	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

var (
	_ CommunityRepo = (*database.CommunityRepo)(nil)
)

// Community is the internal representation of the Community object
type Community struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Location  string
	Type      string
}

// CommunityCore is the core logic of the community repo
type CommunityCore struct {
	CommunityRepo database.CommunityRepo
}

// CommunityRepo is the interface that the community database should implement
type CommunityRepo interface {
	// GetCommunity is the datbase layer to obtain the first matching community
	GetCommunity(ID string) (*models.Community, error)
}

// GetCommunity is the Core logic to manipulate Community-related objects
func (c *CommunityCore) GetCommunity(ID string) (*Community, error) {
	communities, err := c.CommunityRepo.GetCommunity(ID)
	if err != nil {
		return nil, err
	}
	return castCommunity(*communities), nil
}

func castCommunity(communityModel models.Community) *Community {
	return &Community{
		ID: communityModel.ID,
		CreatedAt: communityModel.CreatedAt,
		UpdatedAt: communityModel.UpdatedAt,
		Name: communityModel.Name,
		Location: communityModel.Location,
		Type: "community",
	}
}
