package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/r-rayns/spooky_kingdom/config"
)

// GenerateJWT generates a new JWT token for authentication
func GenerateJWT() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    "spooky_kingdom",
		Subject:   "auth",
		Audience:  "admin",
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
	})
	config := config.Get()

	tokenString, err := token.SignedString(config.JWTSecret)

	if err != nil {
		fmt.Printf("Could not generate JWT: %s", err)
	}

	return tokenString
}

// ParseJWT parses JWT bearer token and extracts claims if valid
func ParseJWT(bearerToken string) (jwt.MapClaims, error) {
	config := config.Get()

	if strings.Count(bearerToken, ".") != 2 {
		return nil, fmt.Errorf("Badly formatted token")
	}

	// remove bearer string if present
	if bearerToken[:len("Bearer")] == "Bearer" {
		bearerToken = bearerToken[len("Bearer")+1:]
	}

	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.JWTSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// VerifyAuthHeader verifies the authorisation header is valid and suitable for the audience
func VerifyAuthHeader(c echo.Context, aud string) error {
	header := c.Request().Header.Get("Authorization")
	claims, err := ParseJWT(header)

	if err != nil || !claims.VerifyAudience(aud, true) {
		return errors.New("Unauthorised")
	}

	return nil
}
