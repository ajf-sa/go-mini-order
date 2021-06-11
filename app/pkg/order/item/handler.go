package item

import "github.com/gofiber/fiber/v2"

type IHandler interface {
	Create(c *fiber.Ctx) error
	FindAllByOrderID(c *fiber.Ctx) error
}

type iHandler struct {
	service IService
}

func NewIHandler(s IService) IHandler {
	return &iHandler{s}
}
func (h *iHandler) Create(c *fiber.Ctx) error {
	return c.JSON("Create")

}
func (h *iHandler) FindAllByOrderID(c *fiber.Ctx) error {
	return c.JSON("FindAllByOrderID")

}
