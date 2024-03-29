package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sf4nu/Anime-Vue-GO-Client/handlers"
	"github.com/sf4nu/Anime-Vue-GO-Client/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// func (h *handlers.Handlers) SetupRoutes(app *fiber.App) {
// 	api := app.Group("/api")
// 	api.Post("/create_books", h.CreateUser)
// }

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	dsn := "postgres://gnbhbxmb:QJsWh1s-Z_sbdvF5gXzpGnyCDgRrokyi@dumbo.db.elephantsql.com/gnbhbxmb"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	h := &handlers.Handlers{DB: db}

	db.AutoMigrate(&models.User{})

	app := fiber.New()

	// h.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/register", h.CreateUser)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Listen("0.0.0.0:" + port)
}
