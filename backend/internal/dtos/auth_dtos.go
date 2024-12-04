// internal/dtos/auth_dtos.go
package dtos

import (
	"backend/pkg/jwt"
)

// Request and Response structures for authentication
type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
    Data struct {
        Token        string `json:"token"`
        RefreshToken string `json:"refreshToken"`
    } `json:"data"`
    Code string `json:"code"`
    Msg  string `json:"msg"`
}

type UserInfoResponse struct {
    Data struct {
        UserID   string   `json:"userId"`
        Username string   `json:"userName"`
        Roles    []string `json:"roles"`
        Buttons  []string `json:"buttons"`
    } `json:"data"`
    Code string `json:"code"`
    Msg  string `json:"msg"`
}

// JWT Claims structure
type JWTClaims struct {
    Data []map[string]interface{} `json:"data"`
    jwt.StandardClaims
}
