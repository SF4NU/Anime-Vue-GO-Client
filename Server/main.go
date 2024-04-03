package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sf4nu/Anime-Vue-GO-Client/handlers"
	"github.com/sf4nu/Anime-Vue-GO-Client/initializers"
	"github.com/sf4nu/Anime-Vue-GO-Client/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.AutoMigrate()
}

func main() {
	var err error
	app := fiber.New()

	middleware.Cors(app)

	app.Post("/register", handlers.CreateUser)
	app.Post("/login", handlers.LoginUser)
	app.Get("/validate", middleware.Auth, handlers.Validate)
	app.Post("/add/anime/:userID", middleware.Auth, handlers.AddAnime)
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
