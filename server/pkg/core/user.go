package core

import (
	"fmt"
	"strings"
	"time"

	"github.com/teamgenerator/teamgenerator/server/pkg/models"
)

// User is the internal representation of the User object
type User struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Name      string
	Password  string
	Type      string
}

// UserCore is the core logic of the user repo
type UserCore struct {
	UserRepo      UserRepo
	CommunityRepo CommunityRepo
}

// UserFilter is the filter that the GetUsers DAL use to filter its results
type UserFilter struct {
	// ID is used to obtain a single user by ID
	ID []string
}

// UserRepo is the interface that the user database should implement
type UserRepo interface {
	// GetUsers is the datbase layer to obtain the first matching user
	// This function should accept a UserFilter and returns the filtered results
	Get(filter UserFilter) ([]models.User, error)
	// Creates a user given a name and location
	Create(name, username, password string) (*models.User, error)
	// Deletes a user given ID
	Delete(ID string) (*models.User, error)
}

// GetUser returns a single user given an ID if exists
func (c *UserCore) GetUser(ID string) (*User, error) {
	userFilter := UserFilter{
		ID: []string{ID},
	}
	users, err := c.UserRepo.Get(userFilter)
	if err != nil {
		return nil, err
	}
	if users == nil {
		return nil, ErrUserNotFound
	}
	parsedCommunities := castUsers(users)
	return &(parsedCommunities[0]), nil
}

// GetUsers returns an array of users without any filter
func (c *UserCore) GetUsers() ([]User, error) {
	users, err := c.UserRepo.Get(UserFilter{})
	if err != nil {
		return nil, err
	}
	parsedCommunities := castUsers(users)
	return parsedCommunities, nil
}

// CreateUser creates a single communtiy given a name and location
func (c *UserCore) CreateUser(name, username, password string) (*User, error) {
	var sb strings.Builder

	if name == "" {
		sb.WriteString("- name must be provided\n")
	}
	if username == "" {
		sb.WriteString("- username must be provided\n")
	}
	if password == "" {
		sb.WriteString("- password must be provided\n")
	}
	if errMsg := sb.String(); errMsg != "" {
		errMsg = fmt.Sprintf("Error validating query params: \n%s", errMsg)
		err := fmt.Errorf("%w, %s", ErrInvalidInputParams, errMsg)
		return nil, err
	}

	user, err := c.UserRepo.Create(name, username, password)
	if err != nil {
		return nil, err
	}
	parsedCommunities := castUsers([]models.User{*user})
	return &(parsedCommunities[0]), nil
}

// DeleteUser deletes a user given an ID
func (c *UserCore) DeleteUser(ID string) (*User, error) {
	user, err := c.UserRepo.Delete(ID)
	if err != nil {
		return nil, err
	}
	parsedCommunities := castUsers([]models.User{*user})
	return &(parsedCommunities[0]), nil
}

func castUsers(users []models.User) []User {
	var parsedCommunities []User
	for _, v := range users {
		newUser := User{
			ID:        v.ID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Username:  v.Username,
			Name:      v.Name,
			Password:  v.Password,
			Type:      "user",
		}
		parsedCommunities = append(parsedCommunities, newUser)
	}
	return parsedCommunities
}
