// backend/internal/resources/errors/user_errors.go
package errors

type ErrorMessage string

// User-related error messages
const (
	USER_NOT_FOUND          ErrorMessage = "User not found"
	INVALID_USER_INPUT      ErrorMessage = "Invalid user input"
	USER_ALREADY_REGISTERED ErrorMessage = "User is already registered"
)
