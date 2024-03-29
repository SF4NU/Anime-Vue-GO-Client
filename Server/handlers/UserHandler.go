package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sf4nu/Anime-Vue-GO-Client/models"
	"github.com/sf4nu/Anime-Vue-GO-Client/utils"
)

func (h *Handlers) CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		c.Status(400).SendString("Bad request")
		return err
	}

	if err := h.DB.Where("username = ?", user.Username).First(&user).Error; err == nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Username already exists")
	}

	if err := h.DB.Where("email = ?", user.Email).First(&user).Error; err == nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Email already exists")
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
