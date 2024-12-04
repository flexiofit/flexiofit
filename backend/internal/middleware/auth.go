// internal/middleware/auth.go
package middleware

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"backend/internal/resources/response"
)

var jwtSecret = []byte("PRAN1231SINGH") // Replace with a strong, environment-configured secret

type JWTClaims struct {
	Data []map[string]string `json:"data"`
	jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.BadRequestError(c, "authorization header is missing")
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, "Bearer ")
		if len(bearerToken) != 2 {
			response.BadRequestError(c, "invalid token format")
			c.Abort()
			return
		}

		tokenString := bearerToken[1]
		token, err := validateToken(tokenString)
		if err != nil {
			response.BadRequestError(c, err.Error())
			c.Abort()
			return
		}

		// You might want to set user info in the context for later use
		if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
			if len(claims.Data) > 0 {
				c.Set("user", claims.Data[0]["userName"])
				c.Next()
				return
			}
		}

		response.BadRequestError(c, "Invalid token")
		c.Abort()
	}
}

func validateToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return jwtSecret, nil
	})
}

func GenerateTokens(username string) (string, string, error) {
	now := time.Now()

	// Access Token (1 year expiration)
	accessTokenClaims := JWTClaims{
		Data: []map[string]string{{"userName": username}},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24 * 365)),
			Issuer:    "Soybean",
			Subject:   username,
			Audience:  jwt.ClaimStrings{"soybean-admin"},
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	// Refresh Token (2 years expiration)
	refreshTokenClaims := JWTClaims{
		Data: []map[string]string{{"userName": username}},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24 * 365 * 2)),
			Issuer:    "Soybean",
			Subject:   username,
			Audience:  jwt.ClaimStrings{"soybean-admin"},
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

// New method for validating refresh tokens
func ValidateRefreshToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid refresh token")
}


// func (m *AuthMiddleware) ValidateRefreshToken(tokenString string) (*JWTClaims, error) {
// 	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("invalid token signing method")
// 		}
// 		return jwtSecret, nil
// 	})
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
// 		return claims, nil
// 	}
//
// 	return nil, errors.New("invalid refresh token")
// }
// k
