package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// IsManager checks if the user is a manager based on JWT claims
func IsManager(c *fiber.Ctx) bool {
	// Get the token from the cookie
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return false
	}

	// Parse the token
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(JWTSecretKey), nil
	})
	if err != nil || !token.Valid {
		return false
	}

	// Extract claims and check the role
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return false
	}

	role, ok := (*claims)["role"].(string)
	return ok && role == "manager"
}
