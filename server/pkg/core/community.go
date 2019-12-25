package core

import (
	"time"

	"github.com/teamgenerator/teamgenerator/server/pkg/models"
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
	CommunityRepo CommunityRepo
}

// CommunityFilter is the filter that the GetCommunities DAL use to filter its results
type CommunityFilter struct {
	// ID is used to obtain a single community by ID
	ID []string
}

// CommunityRepo is the interface that the community database should implement
type CommunityRepo interface {
	// GetCommunities is the datbase layer to obtain the first matching community
	// This function should accept a CommunityFilter and returns the filtered results
	GetCommunities(filter CommunityFilter) ([]models.Community, error)
}

// GetCommunity returns a single community given an ID if exists
func (c *CommunityCore) GetCommunity(ID string) (*Community, error) {
	communityFilter := CommunityFilter{
		ID: []string{ID},
	}
	communities, err := c.CommunityRepo.GetCommunities(communityFilter)
	if err != nil {
		return nil, err
	}
	parsedCommunities := castCommunities(communities)
	return &parsedCommunities[0], nil
}

// GetCommunities returns an array of communities without any filter
func (c *CommunityCore) GetCommunities() ([]Community, error) {
	communities, err := c.CommunityRepo.GetCommunities(CommunityFilter{})
	if err != nil {
		return nil, err
	}
	parsedCommunities := castCommunities(communities)
	return parsedCommunities, nil
}

func castCommunities(communities []models.Community) []Community {
	var parsedCommunities []Community
	for _, v := range communities {
		newCommunity := Community{
			ID: v.ID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Name: v.Name,
			Location: v.Location,
			Type: "community",
		}
		parsedCommunities = append(parsedCommunities, newCommunity)
	}
	return parsedCommunities
}
