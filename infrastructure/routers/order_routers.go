package routers

import (
	"hactive/assigment-2/config"
	"hactive/assigment-2/controllers"
	"hactive/assigment-2/repository"
	"hactive/assigment-2/service"

	"github.com/gofiber/fiber/v2"
)

func RoutesOrder(fiber *fiber.App, conf config.Config) {
	db := config.InitDatabase(conf)

	repo := repository.NewOrderRepository(db)
	service := service.NewOrderService(conf, repo)
	controller := controllers.NewOrderController(service)

	fiber.Post("/orders", controller.CreateOrder)
	fiber.Get("/orders", controller.GetOrders)
	fiber.Put("/orders/:orderId", controller.UpdateOrder)
	fiber.Delete("/orders/:orderId", controller.DeleteOrder)
}
