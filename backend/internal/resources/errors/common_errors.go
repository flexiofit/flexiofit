// backend/internal/resources/errors/common_errors.go
package errors

import "fmt"

// General error messages
const (
	INTERNAL_SERVER_ERROR ErrorMessage = "An internal server error occurred"
	BAD_REQUEST           ErrorMessage = "Invalid request data"
)

// Helper to format error messages with context
func FormatErrorMessage(err ErrorMessage, details ...string) string {
	if len(details) > 0 {
		return fmt.Sprintf("%s: %s", err, details[0])
	}
	return string(err)
}
