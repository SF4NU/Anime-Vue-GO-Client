package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sf4nu/Anime-Vue-GO-Client/initializers"
	"github.com/sf4nu/Anime-Vue-GO-Client/models"
)

func AddManga(c *fiber.Ctx) error {
	var manga models.MangaList

	if err := c.BodyParser(&manga); err != nil {
		c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		return err
	}

	id := c.Params("id")

	newUserId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	var user models.User

	if err := initializers.DB.First(&user, newUserId).Error; err != nil {
		c.Status(fiber.StatusNotFound).SendString("No user found")
		return err
	}

	manga.UserID = newUserId

	var existingManga models.MangaList

	if err := initializers.DB.Where("manga_id = ? AND user_id = ?", existingManga.MangaID, id).First(&existingManga).Error; err == nil {
		c.Status(fiber.StatusConflict).SendString("Manga already in list")
		return err
	}

	if err := initializers.DB.Create(&manga).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("There was a problem with the server")
		return err
	}

	return c.JSON(manga)
}

func UpdateManga(c *fiber.Ctx) error {
	var manga models.MangaList

	id := c.Params("id")

	if err := initializers.DB.First(&manga, id).Error; err != nil {
		c.Status(fiber.StatusNotFound).SendString("Manga not found")
		return err
	}

	var updatedManga models.MangaList

	if err := c.BodyParser(&updatedManga); err != nil {
		c.Status(fiber.StatusBadRequest).SendString("Bad request")
		return err
	}

	if updatedManga.Chapters != manga.Chapters {
		manga.Chapters = updatedManga.Chapters
	}
	if updatedManga.Finished != manga.Finished {
		manga.Finished = updatedManga.Finished
	}
	if updatedManga.Rating != manga.Rating {
		manga.Rating = updatedManga.Rating
	}
	if updatedManga.Comment != manga.Comment {
		manga.Comment = updatedManga.Comment
	}
	if updatedManga.StartingDate != manga.StartingDate {
		manga.StartingDate = updatedManga.StartingDate
	}
	if updatedManga.EndingDate != manga.EndingDate {
		manga.EndingDate = updatedManga.EndingDate
	}
	if updatedManga.PlanToRead != manga.PlanToRead {
		manga.PlanToRead = updatedManga.PlanToRead
	}

	if err := initializers.DB.Save(&manga).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("There was a problem with the server")
		return err
	}

	return c.JSON(&manga)
}

func DeleteManga(c *fiber.Ctx) error {
	id := c.Params("id")

	var manga models.MangaList

	if err := initializers.DB.First(&manga, id).Error; err != nil {
		c.Status(fiber.StatusNotFound).SendString("Manga doesn't exist in list")
		return err
	}

	if err := initializers.DB.Delete(&manga).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("There was a problem with the server")
		return err
	}

	return c.SendStatus(fiber.StatusAccepted)
}

func GetManga(c *fiber.Ctx) error {
	id := c.Params("id") //id dello user a cui appartiene la lista

	var manga []models.MangaList

	if err := initializers.DB.Where("user_id = ?", id).Find(&manga).Error; err != nil {
		c.Status(fiber.StatusNotFound).SendString("No manga found for this user")
		return err
	}

	return c.JSON(manga)
}
