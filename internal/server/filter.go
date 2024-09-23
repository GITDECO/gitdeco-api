package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewCors(c *fiber.Ctx) error {
	return cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, PATCH, UPDATE, DELETE,",
		AllowHeaders: "Origin, Content-Type, Accept",
	})(c)
}

func NewLogger(c *fiber.Ctx) error {
	return logger.New(logger.Config{
		Format: fmt.Sprintf("%s\n", os.Getenv("LOGGER_FORMAT")),
	})(c)
}
