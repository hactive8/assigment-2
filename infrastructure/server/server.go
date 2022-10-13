package server

import (
	"hactive/assigment-2/config"
	"hactive/assigment-2/infrastructure/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Start() *fiber.App {
	app := fiber.New()
	conf := config.Config{}

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"code":    200,
			"message": "API Assigment 2",
		})
	})

	routers.RoutesOrder(app, conf)

	return app
}
