// pkg/jwt/claims.go
package jwt

import "time"

type StandardClaims struct {
    Audience  string    `json:"aud,omitempty"`
    ExpiresAt time.Time `json:"exp,omitempty"`
    Id        string    `json:"jti,omitempty"`
    IssuedAt  time.Time `json:"iat,omitempty"`
    Issuer    string    `json:"iss,omitempty"`
    NotBefore time.Time `json:"nbf,omitempty"`
    Subject   string    `json:"sub,omitempty"`
}

// ToMap converts StandardClaims to a map for easier JSON handling
func (sc StandardClaims) ToMap() map[string]interface{} {
    claims := make(map[string]interface{})
    
    if sc.Audience != "" {
        claims["aud"] = sc.Audience
    }
    if !sc.ExpiresAt.IsZero() {
        claims["exp"] = sc.ExpiresAt.Unix()
    }
    if sc.Id != "" {
        claims["jti"] = sc.Id
    }
    if !sc.IssuedAt.IsZero() {
        claims["iat"] = sc.IssuedAt.Unix()
    }
    if sc.Issuer != "" {
        claims["iss"] = sc.Issuer
    }
    if !sc.NotBefore.IsZero() {
        claims["nbf"] = sc.NotBefore.Unix()
    }
    if sc.Subject != "" {
        claims["sub"] = sc.Subject
    }
    
    return claims
}
