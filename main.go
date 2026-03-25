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

	restraunts := app.Group("/restaurants")
	reviews := app.Group("reviews")

	app.Get("/pizzas", controller.GetPizzas)
	app.Get("/ingredients", controller.GetIngredients)
	app.Get("/chefs", controller.GetChefs)

	restraunts.Get("/", controller.GetRestaraunt)
	restraunts.Get("id:/menu", controller.GetRestrauntsMenu)
	reviews.Get("/", controller.GetReviews)
	reviews.Post("/", controller.PostReviews)

}
