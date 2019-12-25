package core

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

// SessionPlayer is the internal representation of the SessionPlayer object
type SessionPlayer struct {
	ID          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PlayerID    int
	CommunityID int
	Rating      int
	Form        int
	FormChange  int
	Type        string
}

// SessionPlayerCore is the core logic of the sessionPlayers repo
type SessionPlayerCore struct {
	SessionPlayerRepo SessionPlayerRepo
	CommunityRepo     CommunityRepo
	PlayerRepo        PlayerRepo
}

// SessionPlayerFilter is the filter that the GetSessionPlayers DAL use to filter its results
type SessionPlayerFilter struct {
	// ID is used to obtain a single sessionPlayers by ID
	ID []string
}

// SessionPlayerRepo is the interface that the sessionPlayers database should implement
type SessionPlayerRepo interface {
	// GetSessionPlayers is the datbase layer to obtain the first matching sessionPlayers
	// This function should accept a SessionPlayerFilter and returns the filtered results
	Get(filter SessionPlayerFilter) ([]models.SessionPlayer, error)
	// Creates a sessionPlayers given a name and location
	Create(playerID, communityID, rating, form, formChange int) (*models.SessionPlayer, error)
	// Deletes a sessionPlayers given ID
	Delete(ID string) (*models.SessionPlayer, error)
}

// GetSessionPlayer returns a single sessionPlayers given an ID if exists
func (c *SessionPlayerCore) GetSessionPlayer(ID string) (*SessionPlayer, error) {
	sessionPlayersFilter := SessionPlayerFilter{
		ID: []string{ID},
	}
	sessionPlayerss, err := c.SessionPlayerRepo.Get(sessionPlayersFilter)
	if err != nil {
		return nil, err
	}
	if sessionPlayerss == nil {
		return nil, ErrSessionPlayerNotFound
	}
	parsedCommunities := castSessionPlayers(sessionPlayerss)
	return &(parsedCommunities[0]), nil
}

// GetSessionPlayers returns an array of sessionPlayerss without any filter
func (c *SessionPlayerCore) GetSessionPlayers() ([]SessionPlayer, error) {
	sessionPlayerss, err := c.SessionPlayerRepo.Get(SessionPlayerFilter{})
	if err != nil {
		return nil, err
	}
	parsedCommunities := castSessionPlayers(sessionPlayerss)
	return parsedCommunities, nil
}

// CreateSessionPlayer creates a single communtiy given a name and location
func (c *SessionPlayerCore) CreateSessionPlayer(playerID, communityID, rating, form, formChange int) (*SessionPlayer, error) {
	var sb strings.Builder

	if playerID == 0 {
		sb.WriteString("- playerID must be provided\n")
	}
	if communityID == 0 {
		sb.WriteString("- communityID must be provided\n")
	}
	if errMsg := sb.String(); errMsg != "" {
		errMsg = fmt.Sprintf("Error validating query params: \n%s", errMsg)
		err := fmt.Errorf("%w, %s", ErrInvalidInputParams, errMsg)
		return nil, err
	}

	playerFilter := PlayerFilter{
		ID: []string{strconv.Itoa(playerID)},
	}
	player, err := c.PlayerRepo.Get(playerFilter)
	if err != nil {
		return nil, err
	}
	if player == nil {
		return nil, ErrPlayerNotFound
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

	sessionPlayers, err := c.SessionPlayerRepo.Create(playerID, communityID, rating, form, formChange)
	if err != nil {
		return nil, err
	}
	parsedCommunities := castSessionPlayers([]models.SessionPlayer{*sessionPlayers})
	return &(parsedCommunities[0]), nil
}

// DeleteSessionPlayer deletes a sessionPlayers given an ID
func (c *SessionPlayerCore) DeleteSessionPlayer(ID string) (*SessionPlayer, error) {
	sessionPlayers, err := c.SessionPlayerRepo.Delete(ID)
	if err != nil {
		return nil, err
	}
	parsedCommunities := castSessionPlayers([]models.SessionPlayer{*sessionPlayers})
	return &(parsedCommunities[0]), nil
}

func castSessionPlayers(sessionPlayerss []models.SessionPlayer) []SessionPlayer {
	var parsedCommunities []SessionPlayer
	for _, v := range sessionPlayerss {
		newSessionPlayer := SessionPlayer{
			ID:          v.ID,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			PlayerID:    v.PlayerID,
			CommunityID: v.CommunityID,
			Rating:      v.Rating,
			Form:        v.Form,
			FormChange:  v.FormChange,
			Type:        "sessionPlayer",
		}
		parsedCommunities = append(parsedCommunities, newSessionPlayer)
	}
	return parsedCommunities
}
