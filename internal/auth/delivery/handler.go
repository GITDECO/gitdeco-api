package delivery

import (
	"fmt"
	"gitdeco-api/internal/auth"
	"gitdeco-api/internal/exception"
	"gitdeco-api/internal/response"
	"gitdeco-api/pkg"
	"gitdeco-api/tools"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

var gitHubState string

type AuthHandler struct {
	authUsecase auth.Usecase
}

func NewAuthHandler(authUsecase auth.Usecase) auth.Handler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (ah *AuthHandler) Login(c *fiber.Ctx) error {
	state, err := tools.GenerateRandomString()
	if err != nil {
		panic(&exception.Error{Key: "OAUTH_ERROR", Data: err.Error()})
	}
	gitHubState = state
	url := pkg.GithubLogin(c, gitHubState)
	return c.Redirect(url, http.StatusTemporaryRedirect)
}

func (ah *AuthHandler) GithubCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != gitHubState {
		panic(&exception.Error{Key: "OAUTH_ERROR", Data: fmt.Sprintf("'state' parameters are inconsistent with '%s'.", gitHubState)})
	}
	user := pkg.GithubCallback(c)
	token, err := ah.authUsecase.Auth(user)
	if err != nil {
		panic(&exception.Error{Key: "USECASE_ERROR", Data: err.Error()})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token.AccessToken,
		Expires:  time.Now().Add(pkg.AccessTokenExpirationTime),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "strict",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    token.RefreshToken,
		Expires:  time.Now().Add(pkg.RefreshTokenExpirationTime),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "strict",
	})
	c.Set("Content-Security-Policy", "default-src 'self'")
	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	})
}

func (ah AuthHandler) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	token := pkg.RefreshToken(refreshToken)
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token.AccessToken,
		Expires:  time.Now().Add(pkg.AccessTokenExpirationTime),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "strict",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    token.RefreshToken,
		Expires:  time.Now().Add(pkg.RefreshTokenExpirationTime),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "strict",
	})
	c.Set("Content-Security-Policy", "default-src 'self'")
	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	})
}
