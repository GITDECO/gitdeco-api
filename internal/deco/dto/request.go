package dto

import (
	"gitdeco-api/internal/exception"
	"gitdeco-api/internal/models"
	"gitdeco-api/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
)

type DecoRequest struct {
	Title    string `json:"title" validate:"required"`
	Markdown string `json:"markdown" validate:"required"`
}

func (request *DecoRequest) ToEntity(username string) *models.Deco {
	deco := new(models.Deco)
	deco.Title = request.Title
	deco.Markdown = request.Markdown
	deco.Username = username
	deco.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	deco.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	return deco
}

func (request *DecoRequest) ParseX(c *fiber.Ctx) *DecoRequest {
	if err := c.BodyParser(request); err != nil {
		panic(&exception.Error{Key: "DONT_PARSE", Data: ""})
	}

	if err := pkg.Validator(request); err != nil {
		panic(&exception.Error{Key: "BAD_REQUEST", Data: ""})
	}

	return request
}
