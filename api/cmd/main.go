package main

import (
	"log"

	"github.com/adilsonmenechini/simplesrestapi/api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture !"))
	})
	api := app.Group("/api")
	routes.UserRouter(api)
	log.Fatal(app.Listen(":3000"))
}
