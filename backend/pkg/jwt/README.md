# JWT Utility

## Description
This package provides helper functions for handling JWT tokens, including generation and validation.

### Functions
1. **GenerateToken(userName string, expiration time.Duration) (string, error)**
   - Generates a JWT token for the given `userName` with the specified expiration duration.

2. **ValidateToken(tokenStr string) (*CustomClaims, error)**
   - Validates a given JWT token and returns the claims if valid.

### Example Usage
```go
import "your_project/pkg/jwt"

// Generate a token
token, err := jwt.GenerateToken("Soybean", time.Hour*24)
if err != nil {
    log.Fatal(err)
}

// Validate a token
claims, err := jwt.ValidateToken(token)
if err != nil {
    log.Fatal(err)
}
fmt.Println("UserName:", claims.UserName)
