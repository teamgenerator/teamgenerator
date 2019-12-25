package database

import (
	"github.com/teamgenerator/teamgenerator/server/db"
	"github.com/teamgenerator/teamgenerator/server/pkg/core"
	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

var (
	_ core.UserRepo = (*UserRepo)(nil)
)

// UserRepo is the database layer for interacting with the User table
type UserRepo struct{}

// Get returns the users with respect to the given filter
func (r *UserRepo) Get(filter core.UserFilter) ([]models.User, error) {
	var users []models.User

	if len(filter.ID) == 1 {
		result := db.DB.First(&users, filter.ID[0])
		if result.Error != nil {
			return nil, result.Error
		}
		if len(users) < 1 {
			return nil, nil
		}
		return users, nil
	}

	db.DB.Find(&users)
	return users, nil
}

// Create function to create a single user
func (r *UserRepo) Create(name, username, password string) (*models.User, error) {
	user := models.User{
		Username: username,
		Name:     name,
		Password: password,
	}
	var createdUser models.User
	result := db.DB.Create(&user).Scan(&createdUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &createdUser, nil
}

// Delete a single user given an ID
func (r *UserRepo) Delete(ID string) (*models.User, error) {
	var user models.User
	result := db.DB.First(&user, ID)
	if result.Error != nil {
		return nil, result.Error
	}

	if user.ID != 0 {
		result = db.DB.Delete(&user)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	return &user, nil
}
