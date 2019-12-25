package database

import (
	"github.com/teamgenerator/teamgenerator/server/db"
	"github.com/teamgenerator/teamgenerator/server/pkg/core"
	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

var (
	_ core.SessionRepo = (*SessionRepo)(nil)
)

// SessionRepo is the database layer for interacting with the Session table
type SessionRepo struct{}

// Get returns the sessions with respect to the given filter
func (r *SessionRepo) Get(filter core.SessionFilter) ([]models.Session, error) {
	var sessions []models.Session

	if len(filter.ID) == 1 {
		result := db.DB.First(&sessions, filter.ID[0])
		if result.Error != nil {
			return nil, result.Error
		}
		if len(sessions) < 1 {
			return nil, nil
		}
		return sessions, nil
	}

	db.DB.Find(&sessions)
	return sessions, nil
}

// Create function to create a single session
func (r *SessionRepo) Create(isActive bool, communityID int) (*models.Session, error) {
	session := models.Session{
		IsActive:    isActive,
		CommunityID: communityID,
	}
	var createdSession models.Session
	result := db.DB.Create(&session).Scan(&createdSession)
	if result.Error != nil {
		return nil, result.Error
	}
	return &createdSession, nil
}

// Update changes the status of isActive
func (r *SessionRepo) Update(ID string, isActive bool) (*models.Session, error) {
	var session models.Session
	result := db.DB.First(&session, ID)
	if result.Error != nil {
		return nil, result.Error
	}

	session.IsActive = isActive

	result = db.DB.Save(&session)
	if result.Error != nil {
		return nil, result.Error
	}

	return &session, nil
}

// Delete a single session given an ID
func (r *SessionRepo) Delete(ID string) (*models.Session, error) {
	var session models.Session
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
