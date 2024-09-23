package server

import (
	"gitdeco-api/internal/response"
	"gitdeco-api/pkg"
	"gitdeco-api/tools"

	"github.com/gofiber/fiber/v2"
)

func ExceptionHandler(c *fiber.Ctx, err error) error {
	key, data := tools.ErrorParse(err)
	exception := pkg.GetException(key)
	if data != "" {
		return c.Status(exception.Code).JSON(&response.GeneralResponse{
			Code:    exception.Code,
			Message: exception.Message,
			Data:    tools.FormatErrorData(exception.Data.(string), data),
		})
	}
	return c.Status(exception.Code).JSON(&response.GeneralResponse{
		Code:    exception.Code,
		Message: exception.Message,
		Data:    exception.Data,
	})
}
