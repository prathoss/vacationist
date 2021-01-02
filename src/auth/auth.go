package auth

import "github.com/gofiber/fiber/v2"

type loginViewModel struct {
	Nickname string `json:nickname`
	Password string `json:password`
}

func Login(c *fiber.Ctx) error {
	var loginModel loginViewModel
	if err := c.BodyParser(&loginModel); err != nil {
		return fiber.ErrBadRequest
	}

}
