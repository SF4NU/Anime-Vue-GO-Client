package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sf4nu/Anime-Vue-GO-Client/models"
	"github.com/sf4nu/Anime-Vue-GO-Client/utils"
)

func (h *Handlers) CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusBadRequest).SendString("Bad request")
		return err
	}

	if err := h.DB.Where("username = ?", user.Username).First(&user).Error; err == nil {
		c.Status(fiber.StatusConflict).SendString("Username already exists")
		return err
	}

	if err := h.DB.Where("email = ?", user.Email).First(&user).Error; err == nil {
		c.Status(fiber.StatusUnauthorized).SendString("Email already exists")
		return err
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Password can't be hashed")
	}

	user.Password = hashedPassword

	if err := h.DB.Create(&user).Error; err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *Handlers) LoginUser(c *fiber.Ctx) error {
	var login models.Login

	if err := c.BodyParser(&login); err != nil {
		c.Status(fiber.StatusBadRequest).SendString("Bad request")
		return err
	}

	var user models.User

	username := login.Username

	if err := h.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.Status(fiber.StatusForbidden).SendString("Username doesn't exist")
		fmt.Println(login.Username, user.Username)
		return err
	}

	if err := utils.CheckPassword(login.Password, user.Password); err != nil {
		c.Status(fiber.StatusForbidden).SendString("Password doesn't match")
		fmt.Println(login.Password, user.Password)
		return err
	}

	return c.Status(fiber.StatusAccepted).SendString("Login successful")

}

func (h *Handlers) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	if err := h.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.Status(fiber.StatusNotFound).SendString("User not found")
		return err
	}

	var animes []models.AnimeList

	if err := h.DB.Where("user_ID = ?", id).Find(&animes).Error; err != nil {
		c.Status(fiber.StatusNotFound).SendString("No anime found")
		return err
	}

	for _, anime := range animes {
		if err := h.DB.Delete(&anime).Error; err != nil {
			c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
		}
	}

	if err := h.DB.Delete(&user).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
		return err
	}

	return c.Status(fiber.StatusAccepted).SendString("User deleted")
}
