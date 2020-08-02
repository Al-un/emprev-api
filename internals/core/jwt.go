package core

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecretKey string

const jwtClaimsIssuer = "al-un.fr"

func init() {
	jwtSecretKey = "super-secret-jwt-key"
	if secretVar := os.Getenv("SECRET_JWT_KEY"); secretVar != "" {
		jwtSecretKey = secretVar
	}
}

// GenerateJWT generate a JWT for a specific user with claims basically representing
// the user properties. List of claims is based on https://tools.ietf.org/html/rfc7519
// found through https://auth0.com/docs/tokens/jwt-claims. Tokens are valid 60 days
//
// HMAC is chosen over RSA to protect against manipulation:
// https://security.stackexchange.com/a/220190
//
// Generate Token	: https://godoc.org/github.com/dgrijalva/jwt-go#example-New--Hmac
// Custom claims	: https://godoc.org/github.com/dgrijalva/jwt-go#NewWithClaims
func GenerateJWT(user User) (string, error) {
	tokenExpiration := time.Now().Add(time.Hour * 24 * 60)

	userClaims := JwtClaims{
		IsAdmin: user.IsAdmin,
		UserID:  user.ID.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiration.Unix(),
			Issuer:    jwtClaimsIssuer,
			IssuedAt:  time.Now().Unix(),
			Subject:   user.Username,
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := newToken.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecodeJWT(r *http.Request) (*JwtClaims, int) {
	// Fetch the Authorization header
	authHeaders := r.Header["Authorization"]
	if len(authHeaders) == 0 {
		return nil, http.StatusUnauthorized
	}
	authHeader := authHeaders[0]
	if len(authHeader) == 0 {
		return nil, http.StatusUnauthorized
	}
	if authHeader[:6] != "Bearer" {
		return nil, http.StatusUnauthorized
	}

	// Get the header value and strip "Bearer " out
	tokenString := authHeader[7:]

	// Parse token. Make sure hashing method is the correct one
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("[JWT decode] Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecretKey), nil
	})

	// Decipher claims
	if claims, ok := token.Claims.(*JwtClaims); ok {

		// Check token validity
		if token.Valid {
			return claims, http.StatusOK
		}

		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, http.StatusUnauthorized
			}
		}

		return claims, http.StatusUnauthorized
	}

	return nil, http.StatusUnauthorized
}
