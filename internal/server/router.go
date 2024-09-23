package server

import (
	ad "gitdeco-api/internal/auth/delivery"
	dd "gitdeco-api/internal/deco/delivery"
	sd "gitdeco-api/internal/svg/delivery"
	ud "gitdeco-api/internal/user/delivery"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouter(app *fiber.App, DB *gorm.DB) {
	v1 := app.Group(os.Getenv("SERVER_VERSION"))
	api := v1.Group("api")

	auth := api.Group("auth")
	authHandler := AuthDI(DB)
	ad.NewAuthRouter(&auth, authHandler)

	user := api.Group("user")
	userHandler := UserDI(DB)
	ud.NewUserRouter(&user, userHandler)

	svg := api.Group("component")
	svgHandler := SvgDI(DB)
	sd.NewSvgRouter(&svg, svgHandler)

	deco := api.Group("deco")
	decoHandler := DecoDI(DB)
	dd.NewDecoRouter(&deco, decoHandler)
}
