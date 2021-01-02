package main

import (
	"github.com/Prathoss/vacationist/src/db"
	"github.com/Prathoss/vacationist/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	db.Connect()

	server := fiber.New()
	server.Use(recover.New())
	server.Use(logger.New())

	routes.SetRoutes(server)
	server.Listen(":8000")
}
