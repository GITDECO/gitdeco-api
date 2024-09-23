package svg

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	GetBadge(c *fiber.Ctx) error
	GetTemplate(c *fiber.Ctx) error
	GetTemplate2(c *fiber.Ctx) error
}
