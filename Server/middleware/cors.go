package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Cors(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", c.Get("Origin"))
		c.Set("Access-Control-Allow-Credentials", "true")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Method() == "OPTIONS" {
			c.Status(fiber.StatusOK)
			return nil
		}
		return c.Next()
	})
}
