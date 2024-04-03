package middleware

import (
	"fmt"

	"log"
	"os"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sf4nu/Anime-Vue-GO-Client/initializers"
	"github.com/sf4nu/Anime-Vue-GO-Client/models"
)

func Auth(c *fiber.Ctx) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cookie := c.Cookies("jwt")
	fmt.Println(cookie)
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized, cookie not found")
	}

	SECRET_KEY := os.Getenv("SECRET_KEY")

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized, token not found")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Locals("user_id", claims["user_id"])
		if time.Now().Unix() > int64(claims["exp"].(float64)) {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized, token expired")
		}

		var user models.User

		if err := initializers.DB.First(&user, claims["user_id"]).Error; err != nil {
			c.Status(fiber.StatusNotFound).SendString("User not found")
			return err
		}

		c.Locals("user", user)

		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}
}
