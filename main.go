package main

import (
	"PizzaApi/Controller"
	"PizzaApi/Entity"
	"PizzaApi/Repository"
	"PizzaApi/Service"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	err := app.Listen(":8080")
	if err != nil {
		return
	}
	db := Entity.InitDatabase()
	repository := Repository.NewRepository(db)
	service := Service.NewService(repository)
	controller := Controller.NewController(service)
	app.Post("reviews", controller.PostReviews)

	app.Get("/pizzas", controller.GetPizzas)
	app.Get("/ingredients", controller.GetIngredients)
	app.Get("/restaurants", controller.GetRestaurants)
	app.Get("/chefs", controller.GetChefs)
	app.Get("/orders", controller.GetOrders)
	app.Get("/reviews", controller.GetReviews)

}
