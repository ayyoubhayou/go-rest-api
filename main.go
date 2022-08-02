package main

import (
	"github.com/go-rest-api/configs"
	"github.com/go-rest-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectToMongoDb()

	routes.UserRoute((app))

	app.Listen(":6000")
}
