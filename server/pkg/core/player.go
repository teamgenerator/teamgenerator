package core

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

// Player is the internal representation of the Player object
type Player struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Ratings     int
	Form        int
	CommunityID int
	Type        string
}

// PlayerCore is the core logic of the player repo
type PlayerCore struct {
	PlayerRepo    PlayerRepo
	CommunityRepo CommunityRepo
}

// PlayerFilter is the filter that the GetPlayers DAL use to filter its results
type PlayerFilter struct {
	// ID is used to obtain a single player by ID
	ID []string
}

// PlayerRepo is the interface that the player database should implement
type PlayerRepo interface {
	// GetPlayers is the datbase layer to obtain the first matching player
	// This function should accept a PlayerFilter and returns the filtered results
	Get(filter PlayerFilter) ([]models.Player, error)
	// Creates a player given a name and location
	Create(name string, ratings, form, communityID int) (*models.Player, error)
	// Deletes a player given ID
	Delete(ID string) (*models.Player, error)
}

// GetPlayer returns a single player given an ID if exists
func (c *PlayerCore) GetPlayer(ID string) (*Player, error) {
	playerFilter := PlayerFilter{
		ID: []string{ID},
	}
	players, err := c.PlayerRepo.Get(playerFilter)
	if err != nil {
		return nil, err
	}
	parsedCommunities := castPlayers(players)
	return &(parsedCommunities[0]), nil
}

// GetPlayers returns an array of players without any filter
func (c *PlayerCore) GetPlayers() ([]Player, error) {
	players, err := c.PlayerRepo.Get(PlayerFilter{})
	if err != nil {
		return nil, err
	}
	parsedCommunities := castPlayers(players)
	return parsedCommunities, nil
}

// CreatePlayer creates a single communtiy given a name and location
func (c *PlayerCore) CreatePlayer(name string, ratings, form, communityID int) (*Player, error) {
	var sb strings.Builder

	if name == "" {
		sb.WriteString("- name must be provided\n")
	}
	if communityID == 0 {
		sb.WriteString("- communityID must be provided\n")
	}
	if errMsg := sb.String(); errMsg != "" {
		errMsg = fmt.Sprintf("Error validating query params: \n%s", errMsg)
		err := fmt.Errorf("%w, %s", ErrInvalidInputParams, errMsg)
		return nil, err
	}

	communityFilter := CommunityFilter{
		ID: []string{strconv.Itoa(communityID)},
	}
	community, err := c.CommunityRepo.Get(communityFilter)
	if err != nil {
		return nil, err
	}
	if community == nil || len(community) < 1 {
		return nil, ErrCommunityNotFound
	}

	player, err := c.PlayerRepo.Create(name, ratings, form, communityID)
	if err != nil {
		return nil, err
	}
	parsedCommunities := castPlayers([]models.Player{*player})
	return &(parsedCommunities[0]), nil
}

// DeletePlayer deletes a player given an ID
func (c *PlayerCore) DeletePlayer(ID string) (*Player, error) {
	player, err := c.PlayerRepo.Delete(ID)
	if err != nil {
		return nil, err
	}
	parsedCommunities := castPlayers([]models.Player{*player})
	return &(parsedCommunities[0]), nil
}

func castPlayers(players []models.Player) []Player {
	var parsedCommunities []Player
	for _, v := range players {
		newPlayer := Player{
			ID:          v.ID,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			Name:        v.Name,
			Ratings:     v.Ratings,
			Form:        v.Form,
			CommunityID: v.CommunityID,
			Type:        "player",
		}
		parsedCommunities = append(parsedCommunities, newPlayer)
	}
	return parsedCommunities
}
