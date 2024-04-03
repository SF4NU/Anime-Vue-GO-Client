package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, https://sf4nu.github.io/",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
}
