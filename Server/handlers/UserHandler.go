package handlers

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
	"github.com/sf4nu/Anime-Vue-GO-Client/utils"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusBadRequest).SendString("Bad request")
		return err
	}

	if err := initializers.DB.Where("username = ?", user.Username).First(&user).Error; err == nil {
		c.Status(fiber.StatusConflict).SendString("Username already exists")
		return err
	}

	if err := initializers.DB.Where("email = ?", user.Email).First(&user).Error; err == nil {
		c.Status(fiber.StatusUnauthorized).SendString("Email already exists")
		return err
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Password can't be hashed")
	}

	user.Password = hashedPassword

	if err := initializers.DB.Create(&user).Error; err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func LoginUser(c *fiber.Ctx) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var login models.Login

	if err := c.BodyParser(&login); err != nil {
		c.Status(fiber.StatusBadRequest).SendString("Bad request")
		return err
	}

	var user models.User

	username := login.Username

	if err := initializers.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.Status(fiber.StatusForbidden).SendString("Username doesn't exist")
		fmt.Println(login.Username, user.Username)
		return err
	}

	if err := utils.CheckPassword(login.Password, user.Password); err != nil {
		c.Status(fiber.StatusForbidden).SendString("Password doesn't match")
		fmt.Println(login.Password, user.Password)
		return err
	}
	//tutto questo codice serve per creare un token jwt
	//Definizione del payload
	claims := jwt.MapClaims{}
	claims["authorized"] = true                           //l'utente è autorizzato
	claims["user_id"] = user.ID                           //id dell'utente
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //scadenza del token

	//creazione del token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	SECRET_KEY := os.Getenv("SECRET_KEY")
	if SECRET_KEY == "" {
		log.Fatal("SECRET_KEY not found")
	}

	//firma del token
	secretKey := []byte(SECRET_KEY)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    signedToken,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   false, //IN HTTPS DEVE ESSERE TRUE, IN QUESTO CASO è FALSE PERCHE' STIAMO USANDO HTTP
		SameSite: "lax",
	})

	return c.SendStatus(fiber.StatusOK)

}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	if err := initializers.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.Status(fiber.StatusNotFound).SendString("User not found")
		return err
	}

	var animes []models.AnimeList

	if err := initializers.DB.Where("user_ID = ?", id).Find(&animes).Error; err != nil {
		c.Status(fiber.StatusNotFound).SendString("No anime found")
		return err
	}

	for _, anime := range animes {
		if err := initializers.DB.Delete(&anime).Error; err != nil {
			c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
		}
	}

	if err := initializers.DB.Delete(&user).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
		return err
	}

	return c.Status(fiber.StatusAccepted).SendString("User deleted")
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var updatedUser models.User

	if err := c.BodyParser(&updatedUser); err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
		return err
	}

	var user models.User
	if err := initializers.DB.First(&user, id).Error; err != nil {
		c.Status(fiber.StatusNotFound).SendString("User not found")
		return err
	}

	if updatedUser.Username != "" {
		user.Username = updatedUser.Username
	}

	if updatedUser.Password != "" {
		hashedPwd, err := utils.HashPassword(updatedUser.Password)
		if err != nil {
			c.Status(fiber.StatusBadRequest).SendString("Bad request")
		}
		user.Password = hashedPwd

		return c.Status(fiber.StatusCreated).JSON(user)
	}

	// if updatedUser.Email != "" {
	// 	user.Email = updatedUser.Email
	// }

	if err := initializers.DB.Save(&user).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
		return err
	}

	return c.Status(fiber.StatusAccepted).SendString("Account modified successfuly")
}

func GetUsers(c *fiber.Ctx) error {
	var user []models.User

	if err := initializers.DB.Find(&user).Error; err != nil {
		c.Status(fiber.StatusNotFound).SendString("No users found")
	}

	return c.JSON(user)
}

func GetUserProfile(c *fiber.Ctx) error {
	id := c.Params("id")

	authUserID := fmt.Sprintf("%.0f", c.Locals("user_id").(float64))

	if authUserID != id {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	var user models.User

	if err := initializers.DB.First(&user, id).Error; err != nil {
		c.Status(fiber.StatusNotFound).SendString("User Not found")
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"ID":       user.ID,
		"Email":    user.Email,
		"Username": user.Username,
	})
}

func Validate(c *fiber.Ctx) error {
	user := c.Locals("user")

	return c.JSON(user)
}
