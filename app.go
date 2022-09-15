package main

import (
	// "os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/DucGiDay/go-fiber-restapi-firebase/config"
	"github.com/DucGiDay/go-fiber-restapi-firebase/route"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	config.ConnectDatabase()

	setupRoutes(app)
	defer config.FI.Client.Close()
	// app.Listen(":" + os.Getenv("PORT"))
	app.Listen(":4000")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")

	route.UserRoute(api.Group("/users"))
}
