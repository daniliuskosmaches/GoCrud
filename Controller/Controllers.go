package Controller

import (
	"PizzaApi/Service"
	"github.com/gofiber/fiber/v3"
)

type Controller struct {
	PizzaService *Service.PizzaService
}

func NewPizzaController(controller *Service.PizzaService) *Controller {
	return &Controller{PizzaService: controller}
}

func (pc *NewPizzaController) PizzaController(c *fiber.Ctx) {
	result := pc.PizzaService.PizzaService()

}
