package main

import (
	"PizzaApi/Controller"
	"PizzaApi/Service"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	err := app.Listen(":8080")
	if err != nil {
		return
	}
	service := Service.NewPizzaService()
	controller := Controller.Controller(service)
	pizzas := app.Group("/pizzas")

	pizzas.Get("/", controller.PizzaController)

}
