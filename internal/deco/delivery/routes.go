package delivery

import (
	"gitdeco-api/internal/deco"
	"gitdeco-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewDecoRouter(deco *fiber.Router, decoHandler deco.Handler) {
	(*deco).Post("/", middleware.AccessTokenMiddleware, decoHandler.Create)
	(*deco).Get("/:id", decoHandler.Read)
	(*deco).Get("/", decoHandler.ReadAll)
	(*deco).Put("/:id", middleware.AccessTokenMiddleware, decoHandler.Update)
	(*deco).Delete("/:id", middleware.AccessTokenMiddleware, decoHandler.Delete)
}
