package interfaces

import (
	"hactive/assigment-2/entity"

	"github.com/gofiber/fiber/v2"
)

type OrderRepository interface {
	CreateOrder(order *entity.Order) error
	GetOrders() ([]entity.Order, error)
	UpdateOrder(id uint, order *entity.Order) error
	DeleteOrder(id uint) error
}

type OrderService interface {
	CreateOrder(order *entity.Order) error
	GetOrders() ([]entity.Order, error)
	UpdateOrder(id uint, order *entity.Order) error
	DeleteOrder(id uint) error
}

type OrderController interface {
	CreateOrder(c *fiber.Ctx) error
	GetOrders(c *fiber.Ctx) error
	UpdateOrder(c *fiber.Ctx) error
	DeleteOrder(c *fiber.Ctx) error
}
