package delivery

import (
	"gitdeco-api/internal/deco"
	"gitdeco-api/internal/deco/dto"
	"gitdeco-api/internal/exception"
	"gitdeco-api/internal/response"
	"gitdeco-api/tools"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type DecoHandler struct {
	decoUsecase deco.Usecase
}

func NewDecoHandler(decoUsecase deco.Usecase) deco.Handler {
	return &DecoHandler{decoUsecase: decoUsecase}
}

func (dh *DecoHandler) Create(c *fiber.Ctx) error {
	username := c.Locals("username").(string)
	request := new(dto.DecoRequest).ParseX(c)
	if _, err := dh.decoUsecase.Create(username, request); err != nil {
		panic(&exception.Error{Key: "USECASE_ERROR", Data: err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(&response.GeneralResponse{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
	})
}

func (dh *DecoHandler) Read(c *fiber.Ctx) error {
	param := c.Params("id")
	deco, err := dh.decoUsecase.Read(tools.UintParseX(param))
	if err != nil {
		panic(&exception.Error{Key: "USECASE_ERROR", Data: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    deco,
	})
}

func (dh *DecoHandler) ReadAll(c *fiber.Ctx) error {
	decos, err := dh.decoUsecase.ReadAll()
	if err != nil {
		panic(&exception.Error{Key: "USECASE_ERROR", Data: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    decos,
	})
}

func (dh *DecoHandler) Update(c *fiber.Ctx) error {
	param := c.Params("id")
	request := new(dto.DecoRequest).ParseX(c)
	if _, err := dh.decoUsecase.Update(tools.UintParseX(param), request); err != nil {
		panic(&exception.Error{Key: "USECASE_ERROR", Data: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	})
}

func (dh *DecoHandler) Delete(c *fiber.Ctx) error {
	param := c.Params("id")
	if err := dh.decoUsecase.Delete(tools.UintParseX(param)); err != nil {
		panic(&exception.Error{Key: "USECASE_ERROR", Data: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	})
}
