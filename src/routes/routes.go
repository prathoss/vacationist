package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

//SetRoutes registers handlers for routes
func SetRoutes(app *fiber.App) {
	setAuthRoutes(app)
	setVacationsRoutes(app)
}

func setAuthRoutes(app *fiber.App) {
	authGrp := app.Group("auth")
	authGrp.Post("register", func(c *fiber.Ctx) error {
		return c.SendString("User can register here")
	})
	authGrp.Post("login", func(c *fiber.Ctx) error {
		return c.SendString("User can log in here")
	})
}

func setVacationsRoutes(app *fiber.App) {
	vacationsGrp := app.Group("vacations")
	vacationsGrp.Get("", func(c *fiber.Ctx) error {
		return c.SendString("See all vacations for current year")
	})
	vacationsGrp.Post("", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprint("Add vacation log for id: "))
	})
	vacationsGrp.Delete(":id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString(fmt.Sprint("Deletes vacation log for id: ", id))
	})
}
