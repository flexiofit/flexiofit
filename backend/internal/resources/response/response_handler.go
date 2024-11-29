// internal/resources/response.go
package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Response represents a standard API response structure
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{}      `json:"error,omitempty"`
}

// SendSuccessResponse sends a success response with optional data
func SendSuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	})
}

// SendErrorResponse sends a custom error response
func SendErrorResponse(c *gin.Context, statusCode int, message string, error interface{}) {
	if error == nil {
		error = nil
	}
	c.JSON(statusCode, Response{
		Status:  statusCode,
		Message: message,
		Error:   error,
	})
}

// InternalServerError sends a 500 internal server error response
func InternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, Response{
		Status:  http.StatusInternalServerError,
		Message: "Internal Server Error",
		Error:   err.Error(),
	})
}

// BadRequestError sends a 400 bad request response
func BadRequestError(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Response{
		Status:  http.StatusBadRequest,
		Message: message,
		Error:   message,
	})
}

// NotFoundError sends a 404 not found response
func NotFoundError(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{
		Status:  http.StatusNotFound,
		Message: message,
		Error:   message,
	})
}

// UnauthorizedError sends a 401 unauthorized response
func UnauthorizedError(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized",
		Error:   "Unauthorized access",
	})
}
