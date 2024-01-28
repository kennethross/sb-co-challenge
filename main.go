package main

import (
	"log"
	"snake-backend/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/new", handler.GetNewHandler)
	app.Post("/validator", handler.PostValidatorHandler)

	log.Fatal(app.Listen(":3000"))
}
