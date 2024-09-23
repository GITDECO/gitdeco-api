package middleware

import (
	"gitdeco-api/internal/exception"
	"gitdeco-api/pkg"

	"github.com/gofiber/fiber/v2"
)

func GlobalMiddleware(c *fiber.Ctx) error {
	return &exception.Error{Key: "NOT_FOUND", Data: c.OriginalURL()}
}

func AccessTokenMiddleware(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")
	username := pkg.ValidateToken(accessToken)
	c.Locals("username", username)
	return c.Next()
}
