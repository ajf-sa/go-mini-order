package order

import "github.com/gofiber/fiber/v2"

type OHandler interface {
	Create(c *fiber.Ctx) error
	FindOneByID(c *fiber.Ctx) error
	FindAllByUserID(c *fiber.Ctx) error
	FindAllByClientID(c *fiber.Ctx) error
}

type oHandler struct {
	service OService
}

func NewOHandler(s OService) OHandler {
	return &oHandler{s}

}
func (h *oHandler) Create(c *fiber.Ctx) error {
	return c.JSON("Create")

}
func (h *oHandler) FindOneByID(c *fiber.Ctx) error {
	return c.JSON("FindOneByID")

}
func (h *oHandler) FindAllByUserID(c *fiber.Ctx) error {
	return c.JSON("FindAllByUserID")

}
func (h *oHandler) FindAllByClientID(c *fiber.Ctx) error {
	return c.JSON("FindAllByClientID")

}
