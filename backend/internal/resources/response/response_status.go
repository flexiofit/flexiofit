package response

const (
	STATUS_OK                  = 200 // Request succeeded
	STATUS_CREATED             = 201 // Resource created successfully
	STATUS_NO_CONTENT          = 204 // No content to return
	STATUS_BAD_REQUEST         = 400 // Invalid client request
	STATUS_UNAUTHORIZED        = 401 // Authentication required or failed
	STATUS_FORBIDDEN           = 403 // Insufficient permissions
	STATUS_NOT_FOUND           = 404 // Resource not found
	STATUS_CONFLICT            = 409 // Conflict with current state (e.g., duplicate resource)
	STATUS_INTERNAL_SERVER_ERR = 500 // Internal server error
	STATUS_SERVICE_UNAVAILABLE = 503 // Service unavailable
)

