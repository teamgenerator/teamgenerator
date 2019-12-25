package core

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

// Session is the internal representation of the Session object
type Session struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsActive    bool
	CommunityID int
	Type        string
}

// SessionCore is the core logic of the session repo
type SessionCore struct {
	SessionRepo   SessionRepo
	CommunityRepo CommunityRepo
}

// SessionFilter is the filter that the GetSessions DAL use to filter its results
type SessionFilter struct {
	// ID is used to obtain a single session by ID
	ID []string
}

// SessionRepo is the interface that the session database should implement
type SessionRepo interface {
	// GetSessions is the datbase layer to obtain the first matching session
	// This function should accept a SessionFilter and returns the filtered results
	Get(filter SessionFilter) ([]models.Session, error)
	// Creates a session given a name and location
	Create(isActive bool, communityID int) (*models.Session, error)
	// Deletes a session given ID
	Delete(ID string) (*models.Session, error)
	// Updates a session with the desired isActive
	Update(ID string, isActive bool) (*models.Session, error)
}

// GetSession returns a single session given an ID if exists
func (c *SessionCore) GetSession(ID string) (*Session, error) {
	sessionFilter := SessionFilter{
		ID: []string{ID},
	}
	sessions, err := c.SessionRepo.Get(sessionFilter)
	if err != nil {
		return nil, err
	}
	if sessions == nil {
		return nil, ErrSessionNotFound
	}
	parsedCommunities := castSessions(sessions)
	return &(parsedCommunities[0]), nil
}

// GetSessions returns an array of sessions without any filter
func (c *SessionCore) GetSessions() ([]Session, error) {
	sessions, err := c.SessionRepo.Get(SessionFilter{})
	if err != nil {
		return nil, err
	}
	parsedCommunities := castSessions(sessions)
	return parsedCommunities, nil
}

// CreateSession creates a single communtiy given a name and location
func (c *SessionCore) CreateSession(isActive bool, communityID int) (*Session, error) {
	var sb strings.Builder

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

	session, err := c.SessionRepo.Create(isActive, communityID)
	if err != nil {
		return nil, err
	}
	parsedCommunities := castSessions([]models.Session{*session})
	return &(parsedCommunities[0]), nil
}

// DeleteSession deletes a session given an ID
func (c *SessionCore) DeleteSession(ID string) (*Session, error) {
	session, err := c.SessionRepo.Delete(ID)
	if err != nil {
		return nil, err
	}
	parsedCommunities := castSessions([]models.Session{*session})
	return &(parsedCommunities[0]), nil
}

// UpdateSession updates a session given an ID with the intended active state
func (c *SessionCore) UpdateSession(ID string, isActive bool) (*Session, error) {
	session, err := c.SessionRepo.Update(ID, isActive)
	if err != nil {
		return nil, err
	}
	parsedCommunities := castSessions([]models.Session{*session})
	return &(parsedCommunities[0]), nil
}

func castSessions(sessions []models.Session) []Session {
	var parsedCommunities []Session
	for _, v := range sessions {
		newSession := Session{
			ID:          v.ID,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			IsActive:    v.IsActive,
			CommunityID: v.CommunityID,
			Type:        "session",
		}
		parsedCommunities = append(parsedCommunities, newSession)
	}
	return parsedCommunities
}
