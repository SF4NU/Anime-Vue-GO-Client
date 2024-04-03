package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sf4nu/Anime-Vue-GO-Client/handlers"
	"github.com/sf4nu/Anime-Vue-GO-Client/initializers"
	"github.com/sf4nu/Anime-Vue-GO-Client/middleware"
)

func init() {
	initializers.ConnectToDb()
	initializers.AutoMigrate()
}

func main() {
	var err error
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "https://sf4nu.github.io",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Post("/register", handlers.CreateUser)
	app.Post("/login", handlers.LoginUser)
	app.Get("/validate", middleware.Auth, handlers.Validate)
	app.Post("/add/anime/:userID", middleware.Auth, handlers.AddAnime)
	app.Post("logout", handlers.Logout)
	app.Delete("/delete/anime/:id", middleware.Auth, handlers.DeleteAnime)
	app.Put("/update/anime/:id", handlers.UpdateAnime)
	app.Put("/update/user/:id", handlers.UpdateUser)
	app.Get("/users/:userID/anime", middleware.Auth, handlers.GetAnime)
	app.Get("/users/", handlers.GetUsers)
	app.Get("/users/:id", middleware.Auth, handlers.GetUserProfile)
	app.Delete("/users/delete/:id", handlers.DeleteUser)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	err = app.Listen("0.0.0.0:" + port)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
