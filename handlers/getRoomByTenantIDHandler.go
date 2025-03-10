package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/khemingkapat/fiber_example/auth"
	"github.com/khemingkapat/fiber_example/queries"
	"gorm.io/gorm"
)

func GetRoomByTenantIDHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("jwt")

		// Parse the token from the cookie
		token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (any, error) {
			return []byte(auth.JWTSecretKey), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
		}

		// Extract claims from the token
		claims, ok := token.Claims.(*jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid token claims")
		}

		id, exists := (*claims)["user_id"].(float64)
		log.Println(id)

		if !exists {
			return c.Status(fiber.StatusBadRequest).SendString("No ID")
		}

		room := queries.GetRoomByTenantID(db, uint(id))
		return c.JSON(room)
	}
}
