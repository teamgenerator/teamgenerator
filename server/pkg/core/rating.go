package core

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

// Rating is the internal representation of the Rating object
type Rating struct {
	ID         int
	RatingGain int
	PlayerID   int
	Type       string
}

// RatingCore is the core logic of the rating repo
type RatingCore struct {
	RatingRepo RatingRepo
	PlayerRepo PlayerRepo
}

// RatingFilter is the filter that the GetRatings DAL use to filter its results
type RatingFilter struct {
	// ID is used to obtain a single rating by ID
	ID []string
}

// RatingRepo is the interface that the rating database should implement
type RatingRepo interface {
	// GetRatings is the datbase layer to obtain the first matching rating
	// This function should accept a RatingFilter and returns the filtered results
	Get(filter RatingFilter) ([]models.Rating, error)
	// Creates a rating given a name and location
	Create(ratingGain, playerID int) (*models.Rating, error)
	// Deletes a rating given ID
	Delete(ID string) (*models.Rating, error)
}

// GetRating returns a single rating given an ID if exists
func (c *RatingCore) GetRating(ID string) (*Rating, error) {
	ratingFilter := RatingFilter{
		ID: []string{ID},
	}
	ratings, err := c.RatingRepo.Get(ratingFilter)
	if err != nil {
		return nil, err
	}
	if ratings == nil {
		return nil, ErrRatingNotFound
	}
	parsedCommunities := castRatings(ratings)
	return &(parsedCommunities[0]), nil
}

// GetRatings returns an array of ratings without any filter
func (c *RatingCore) GetRatings() ([]Rating, error) {
	ratings, err := c.RatingRepo.Get(RatingFilter{})
	if err != nil {
		return nil, err
	}
	parsedCommunities := castRatings(ratings)
	return parsedCommunities, nil
}

// CreateRating creates a single communtiy given a name and location
func (c *RatingCore) CreateRating(ratingGain, playerID int) (*Rating, error) {
	var sb strings.Builder

	if playerID == 0 {
		sb.WriteString("- playerID must be provided\n")
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

	rating, err := c.RatingRepo.Create(ratingGain, playerID)
	if err != nil {
		return nil, err
	}
	parsedCommunities := castRatings([]models.Rating{*rating})
	return &(parsedCommunities[0]), nil
}

// DeleteRating deletes a rating given an ID
func (c *RatingCore) DeleteRating(ID string) (*Rating, error) {
	rating, err := c.RatingRepo.Delete(ID)
	if err != nil {
		return nil, err
	}
	parsedCommunities := castRatings([]models.Rating{*rating})
	return &(parsedCommunities[0]), nil
}

func castRatings(ratings []models.Rating) []Rating {
	var parsedCommunities []Rating
	for _, v := range ratings {
		newRating := Rating{
			ID:         v.ID,
			RatingGain: v.RatingGain,
			PlayerID:   v.PlayerID,
			Type:       "rating",
		}
		parsedCommunities = append(parsedCommunities, newRating)
	}
	return parsedCommunities
}
