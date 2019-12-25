package database

import (
	"github.com/teamgenerator/teamgenerator/server/db"
	"github.com/teamgenerator/teamgenerator/server/pkg/core"
	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

var (
	_ core.PlayerRepo = (*PlayerRepo)(nil)
)

// PlayerRepo is the database layer for interacting with the Player table
type PlayerRepo struct{}

// Get returns the players with respect to the given filter
func (r *PlayerRepo) Get(filter core.PlayerFilter) ([]models.Player, error) {
	var players []models.Player

	if len(filter.ID) == 1 {
		result := db.DB.First(&players, filter.ID[0])
		if result.Error != nil {
			return nil, result.Error
		}
		return players, nil
	}

	db.DB.Find(&players)
	return players, nil
}

// Create function to create a single player
func (r *PlayerRepo) Create(name string, ratings, form, communityID int) (*models.Player, error) {
	player := models.Player{
		Name:        name,
		Ratings:     ratings,
		Form:        form,
		CommunityID: communityID,
	}
	var createdPlayer models.Player
	result := db.DB.Create(&player).Scan(&createdPlayer)
	if result.Error != nil {
		return nil, result.Error
	}
	return &createdPlayer, nil
}

// Delete a single player given an ID
func (r *PlayerRepo) Delete(ID string) (*models.Player, error) {
	var player models.Player
	result := db.DB.First(&player, ID)
	if result.Error != nil {
		return nil, result.Error
	}

	if player.ID != 0 {
		result = db.DB.Delete(&player)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	return &player, nil
}
