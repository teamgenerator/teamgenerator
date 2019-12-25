package database

import (
	"github.com/teamgenerator/teamgenerator/server/db"
	"github.com/teamgenerator/teamgenerator/server/pkg/core"
	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

var (
	_ core.SessionPlayerRepo = (*SessionPlayerRepo)(nil)
)

// SessionPlayerRepo is the database layer for interacting with the SessionPlayer table
type SessionPlayerRepo struct{}

// Get returns the sessionPlayers with respect to the given filter
func (r *SessionPlayerRepo) Get(filter core.SessionPlayerFilter) ([]models.SessionPlayer, error) {
	var sessionPlayers []models.SessionPlayer

	if len(filter.ID) == 1 {
		result := db.DB.First(&sessionPlayers, filter.ID[0])
		if result.Error != nil {
			return nil, result.Error
		}
		if len(sessionPlayers) < 1 {
			return nil, nil
		}
		return sessionPlayers, nil
	}

	db.DB.Find(&sessionPlayers)
	return sessionPlayers, nil
}

// Create function to create a single session
func (r *SessionPlayerRepo) Create(playerID, communityID, rating, form, formChange int) (*models.SessionPlayer, error) {
	session := models.SessionPlayer{
		PlayerID:    playerID,
		CommunityID: communityID,
		Rating:      rating,
		Form:        form,
		FormChange:  formChange,
	}
	var createdSessionPlayer models.SessionPlayer
	result := db.DB.Create(&session).Scan(&createdSessionPlayer)
	if result.Error != nil {
		return nil, result.Error
	}
	return &createdSessionPlayer, nil
}

// Delete a single session given an ID
func (r *SessionPlayerRepo) Delete(ID string) (*models.SessionPlayer, error) {
	var session models.SessionPlayer
	result := db.DB.First(&session, ID)
	if result.Error != nil {
		return nil, result.Error
	}

	if session.ID != 0 {
		result = db.DB.Delete(&session)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	return &session, nil
}
