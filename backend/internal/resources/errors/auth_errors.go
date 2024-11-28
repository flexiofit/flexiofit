// backend/internal/resources/errors/auth_errors.go
package errors

// Auth-related error messages
const (
	UNAUTHORIZED_ACCESS ErrorMessage = "Unauthorized access"
	INVALID_CREDENTIALS ErrorMessage = "Invalid credentials"
	SESSION_EXPIRED     ErrorMessage = "Session has expired"
	PERMISSION_DENIED   ErrorMessage = "Permission denied"
)
