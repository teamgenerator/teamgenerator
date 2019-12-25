package database

import (
	"github.com/teamgenerator/teamgenerator/server/db"
	"github.com/teamgenerator/teamgenerator/server/pkg/core"
	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

var (
	_ core.RatingRepo = (*RatingRepo)(nil)
)

// RatingRepo is the database layer for interacting with the Rating table
type RatingRepo struct{}

// Get returns the ratings with respect to the given filter
func (r *RatingRepo) Get(filter core.RatingFilter) ([]models.Rating, error) {
	var ratings []models.Rating

	if len(filter.ID) == 1 {
		result := db.DB.First(&ratings, filter.ID[0])
		if result.Error != nil {
			return nil, result.Error
		}
		if len(ratings) < 1 {
			return nil, nil
		}
		return ratings, nil
	}

	db.DB.Find(&ratings)
	return ratings, nil
}

// Create function to create a single rating
func (r *RatingRepo) Create(ratingGain, playerID int) (*models.Rating, error) {
	rating := models.Rating{
		RatingGain: ratingGain,
		PlayerID:   playerID,
	}
	var createdRating models.Rating
	result := db.DB.Create(&rating).Scan(&createdRating)
	if result.Error != nil {
		return nil, result.Error
	}
	return &createdRating, nil
}

// Delete a single rating given an ID
func (r *RatingRepo) Delete(ID string) (*models.Rating, error) {
	var rating models.Rating
	result := db.DB.First(&rating, ID)
	if result.Error != nil {
		return nil, result.Error
	}

	if rating.ID != 0 {
		result = db.DB.Delete(&rating)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	return &rating, nil
}
