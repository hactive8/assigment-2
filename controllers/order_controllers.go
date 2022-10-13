package controllers

import (
	"hactive/assigment-2/entity"
	"hactive/assigment-2/interfaces"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Service interfaces.OrderService
}

func NewOrderController(service interfaces.OrderService) interfaces.OrderController {
	return &Controller{
		Service: service,
	}
}

func (h *Controller) CreateOrder(c *fiber.Ctx) error {
	order := entity.Order{}

	err := c.BodyParser(&order)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "failed to parse request body",
		})
	}

	err = h.Service.CreateOrder(&order)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "success create order",
	})
}

func (h *Controller) GetOrders(c *fiber.Ctx) error {
	orders, err := h.Service.GetOrders()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "success get orders",
		"data":    orders,
	})
}

func (h *Controller) UpdateOrder(c *fiber.Ctx) error {
	id := c.Params("orderId")
	orderId, _ := strconv.Atoi(id)

	order := entity.Order{}

	err := c.BodyParser(&order)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "failed to parse request body",
		})
	}

	err = h.Service.UpdateOrder(uint(orderId), &order)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "success update order",
	})
}

func (h *Controller) DeleteOrder(c *fiber.Ctx) error {
	id := c.Params("orderId")
	orderId, _ := strconv.Atoi(id)

	err := h.Service.DeleteOrder(uint(orderId))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "success delete order",
	})
}
