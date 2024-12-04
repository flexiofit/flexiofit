package jwt

import "errors"

var (
	ErrTokenInvalid     = errors.New("token is invalid")
	ErrTokenExpired     = errors.New("token has expired")
	ErrTokenNotProvided = errors.New("token not provided")
)
