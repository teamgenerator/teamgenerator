package core

import "errors"

var (
	ErrInvalidInputParams = errors.New("the input parameters are not valid")
	ErrCommunityNotFound  = errors.New("the desired community is not found")
	ErrSessionNotFound    = errors.New("the desired session is not found")
	ErrPlayerNotFound     = errors.New("the desired player is not found")
)
