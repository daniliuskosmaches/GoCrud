package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	err := app.Listen(":8080")
	if err != nil {
		return
	}

}
