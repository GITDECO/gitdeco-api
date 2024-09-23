package delivery

import (
	"gitdeco-api/internal/middleware"
	"gitdeco-api/internal/user"

	"github.com/gofiber/fiber/v2"
)

func NewUserRouter(user *fiber.Router, userHandler user.Handler) {
	(*user).Get("/my/info", middleware.AccessTokenMiddleware, userHandler.MyInfo)
	(*user).Get("/:username", userHandler.Read)
	(*user).Get("/", userHandler.ReadAll)
}
