package core

import (
	"fmt"
	"strings"
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
	Get(filter CommunityFilter) ([]models.Community, error)
	// Creates a community given a name and location
	Create(name, location string) (*models.Community, error)
	// Deletes a community given ID
	Delete(ID string) (*models.Community, error)
}

// GetCommunity returns a single community given an ID if exists
func (c *CommunityCore) GetCommunity(ID string) (*Community, error) {
	communityFilter := CommunityFilter{
		ID: []string{ID},
	}
	communities, err := c.CommunityRepo.Get(communityFilter)
	if err != nil {
		return nil, err
	}
	if communities == nil {
		return nil, ErrCommunityNotFound
	}
	parsedCommunities := castCommunities(communities)
	return &(parsedCommunities[0]), nil
}

// GetCommunities returns an array of communities without any filter
func (c *CommunityCore) GetCommunities() ([]Community, error) {
	communities, err := c.CommunityRepo.Get(CommunityFilter{})
	if err != nil {
		return nil, err
	}
	parsedCommunities := castCommunities(communities)
	return parsedCommunities, nil
}

// CreateCommunity creates a single communtiy given a name and location
func (c *CommunityCore) CreateCommunity(name, location string) (*Community, error) {
	var sb strings.Builder

	if name == "" {
		sb.WriteString("- name must be provided\n")
	}
	if location == "" {
		sb.WriteString("- location must be provided\n")
	}
	if errMsg := sb.String(); errMsg != "" {
		errMsg = fmt.Sprintf("Error validating query params: \n%s", errMsg)
		err := fmt.Errorf("%w, %s", ErrInvalidInputParams, errMsg)
		return nil, err
	}

	community, err := c.CommunityRepo.Create(name, location)
	if err != nil {
		return nil, err
	}
	parsedCommunities := castCommunities([]models.Community{*community})
	return &(parsedCommunities[0]), nil
}

// DeleteCommunity deletes a community given an ID
func (c *CommunityCore) DeleteCommunity(ID string) (*Community, error) {
	community, err := c.CommunityRepo.Delete(ID)
	if err != nil {
		return nil, err
	}
	parsedCommunities := castCommunities([]models.Community{*community})
	return &(parsedCommunities[0]), nil
}

func castCommunities(communities []models.Community) []Community {
	var parsedCommunities []Community
	for _, v := range communities {
		newCommunity := Community{
			ID:        v.ID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Name:      v.Name,
			Location:  v.Location,
			Type:      "community",
		}
		parsedCommunities = append(parsedCommunities, newCommunity)
	}
	return parsedCommunities
}
