package middlewares

import "github.com/gofiber/fiber/v2"

func Authenticator(c *fiber.Ctx) error {
	return c.Next()
}
