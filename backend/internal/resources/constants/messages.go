package resources

const (
	SUCCESS = "success"
	FAILED  = "failed"
)
// Auth-related error and success messages
const (
	UNAUTHORIZED_ACCESS        = "Unauthorized access"
	INVALID_CREDENTIALS        = "Invalid credentials"
	SESSION_EXPIRED            = "Session has expired"
	PERMISSION_DENIED          = "Permission denied"
	PASSWORD_RESET_SUCCESSFUL  = "Password reset successful"
	PASSWORD_CHANGE_REQUIRED   = "Password change is required"
	LOGIN_SUCCESSFUL           = "Login successful"
	LOGOUT_SUCCESSFUL          = "Logout successful"
)

// General error and success messages
const (
	INTERNAL_SERVER_ERROR      = "An internal server error occurred"
	BAD_REQUEST                = "Invalid request data"
	RESOURCE_NOT_FOUND         = "Resource not found"
	OPERATION_SUCCESSFUL       = "Operation completed successfully"
	OPERATION_FAILED           = "Operation failed"
	VALIDATION_ERROR           = "Validation error"
	SERVICE_UNAVAILABLE        = "Service is temporarily unavailable"
)

// User-related error and success messages
const (
	USER_NOT_FOUND             = "User not found"
	USER_CREATED               = "User created successfully"
	INVALID_USER_INPUT         = "Invalid user input"
	USER_ALREADY_REGISTERED    = "User is already registered"
	USER_DELETED_SUCCESSFULLY  = "User deleted successfully"
	USER_UPDATE_SUCCESSFUL     = "User updated successfully"
	EMAIL_ALREADY_IN_USE       = "Email address is already in use"
	USERNAME_ALREADY_TAKEN     = "Username is already taken"
	ACCOUNT_LOCKED             = "User account is locked"
)

// File-related error and success messages
const (
	FILE_UPLOAD_SUCCESSFUL     = "File uploaded successfully"
	FILE_UPLOAD_FAILED         = "File upload failed"
	FILE_NOT_FOUND             = "File not found"
	FILE_DELETE_SUCCESSFUL     = "File deleted successfully"
	FILE_FORMAT_NOT_SUPPORTED  = "File format not supported"
)

// Payment-related error and success messages
const (
	PAYMENT_SUCCESSFUL         = "Payment processed successfully"
	PAYMENT_FAILED             = "Payment processing failed"
	INSUFFICIENT_FUNDS         = "Insufficient funds"
	CARD_EXPIRED               = "Card has expired"
	INVALID_PAYMENT_METHOD     = "Invalid payment method"
	REFUND_SUCCESSFUL          = "Refund issued successfully"
	REFUND_FAILED              = "Refund process failed"
)

// Notification-related messages
const (
	NOTIFICATION_SENT          = "Notification sent successfully"
	NOTIFICATION_FAILED        = "Failed to send notification"
	SUBSCRIPTION_SUCCESSFUL    = "Subscription successful"
	SUBSCRIPTION_FAILED        = "Subscription failed"
)
