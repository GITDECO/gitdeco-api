package user

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	MyInfo(c *fiber.Ctx) error
	Read(c *fiber.Ctx) error
	ReadAll(c *fiber.Ctx) error
}
