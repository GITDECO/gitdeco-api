package main

import (
	"flag"
	"fmt"
	"gitdeco-api/internal/middleware"
	"gitdeco-api/internal/server"
	"gitdeco-api/pkg"
	"gitdeco-api/tools"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var port string

func init() {
	if err := tools.EnvFileLoad(); err != nil {
		panic(err)
	}

	flogPort := flag.String("p", os.Getenv("SERVER_PORT"), "Enter the port")
	flag.Parse()
	port = fmt.Sprintf(":%s", *flogPort)
}

func main() {
	app := fiber.New(fiber.Config{
		AppName:      os.Getenv("SERVER_NAME"),
		Prefork:      true,
		ErrorHandler: server.ExceptionHandler,
	})

	app.Use(recover.New())
	app.Use(helmet.New())
	app.Use(server.NewCors)
	app.Use(server.NewLogger)

	DB, err := pkg.NewDatabase()
	if err != nil {
		panic(err)
	}

	server.NewRouter(app, DB)

	app.Use(middleware.GlobalMiddleware)

	log.Fatal(app.Listen(port))
}
