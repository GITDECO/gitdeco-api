package delivery

import (
	"gitdeco-api/internal/auth"

	"github.com/gofiber/fiber/v2"
)

func NewAuthRouter(auth *fiber.Router, authHandler auth.Handler) {
	(*auth).Get("/github/login", authHandler.Login)
	(*auth).Get("/github/callback", authHandler.GithubCallback)
	(*auth).Get("/token/refresh", authHandler.RefreshToken)
}
