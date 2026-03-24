package Controller

import (
	"PizzaApi/Service"
	"github.com/gofiber/fiber/v3"
)

type NewController struct {
	PizzaService *Service.PizzaService
}

// контструктор для DI(dependency injection) чтоб можно было вызывать любой пакет Гошки
func Controller(service *Service.PizzaService) *NewController {
	return &NewController{PizzaService: service}
}

func (pc *NewController) PizzaController(c *fiber.Ctx) error {

}

func (pc *NewController) ChefController(c *fiber.Ctx) {

}
func (pc *NewController) IngriController(c *fiber.Ctx) {

}

func (pc *NewController) ReviewController(c *fiber.Ctx) {

}
func (pc *NewController) RestaraurantController(c *fiber.Ctx) {

}
