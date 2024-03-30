package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sf4nu/Anime-Vue-GO-Client/handlers"
	"github.com/sf4nu/Anime-Vue-GO-Client/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	var err error

	dsn := "postgres://gnbhbxmb:QJsWh1s-Z_sbdvF5gXzpGnyCDgRrokyi@dumbo.db.elephantsql.com/gnbhbxmb"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	h := &handlers.Handlers{DB: db}

	db.AutoMigrate(&models.User{}, &models.AnimeList{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/register", h.CreateUser)
	app.Post("/login", h.LoginUser)
	app.Post("/add/anime/:userID", h.AddAnime)
	app.Delete("/delete/anime/:id", h.DeleteAnime)
	app.Put("/update/anime/:id", h.UpdateAnime)
	app.Get("/users/:userID/anime", h.GetAnime)
	app.Delete("/users/delete/:id", h.DeleteUser)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Listen("0.0.0.0:" + port)
}
