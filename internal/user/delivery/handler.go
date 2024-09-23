package delivery

import (
	"gitdeco-api/internal/exception"
	"gitdeco-api/internal/response"
	"gitdeco-api/internal/user"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase user.Usecase
}

func NewUserHandler(userUsecase user.Usecase) user.Handler {
	return &UserHandler{userUsecase: userUsecase}
}

func (uh *UserHandler) MyInfo(c *fiber.Ctx) error {
	username := c.Locals("username").(string)
	user, err := uh.userUsecase.Read(username)
	if err != nil {
		panic(&exception.Error{Key: "USECASE_ERROR", Data: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    user,
	})
}

func (uh *UserHandler) Read(c *fiber.Ctx) error {
	param := c.Params("username")
	user, err := uh.userUsecase.Read(param)
	if err != nil {
		panic(&exception.Error{Key: "USECASE_ERROR", Data: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    user,
	})
}

func (uh *UserHandler) ReadAll(c *fiber.Ctx) error {
	users, err := uh.userUsecase.ReadAll()
	if err != nil {
		panic(&exception.Error{Key: "USECASE_ERROR", Data: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    users,
	})
}
