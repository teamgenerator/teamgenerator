package core

import "errors"

var (
	ErrInvalidInputParams = errors.New("the input parameters are not valid")
	ErrCommunityNotFound  = errors.New("the desired community is not found")
	ErrSessionsNotFound   = errors.New("the desired session is not found")
)
