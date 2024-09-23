package delivery

import (
	"gitdeco-api/internal/svg"

	"github.com/gofiber/fiber/v2"
)

func NewSvgRouter(svg *fiber.Router, svgHandler svg.Handler) {
	(*svg).Get("/badge", svgHandler.GetBadge)
	(*svg).Get("/template", svgHandler.GetTemplate)
	(*svg).Get("/template2", svgHandler.GetTemplate2)
}
