package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/config"
	"github.com/retry19/challenge-hacktiv8/12-final-project/pkg/hasher"
	"github.com/retry19/challenge-hacktiv8/12-final-project/pkg/utility"
)

func NewAuthJwt() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(utility.Response{
				Status:  false,
				Message: "Unauthorized",
			})
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(utility.Response{
				Status:  false,
				Message: "Invalid format token",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := hasher.VerifyJwt(tokenString, config.JwtSecret)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(utility.Response{
				Status:  false,
				Message: "Invalid token",
			})
		}

		c.Locals("UserId", token.Claims.(jwt.MapClaims)["userId"])

		return c.Next()
	}
}
