package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sf4nu/Anime-Vue-GO-Client/initializers"
	"github.com/sf4nu/Anime-Vue-GO-Client/models"
)

func AddAnime(c *fiber.Ctx) error {
	var anime models.AnimeList

	if err := c.BodyParser(&anime); err != nil {
		c.Status(fiber.StatusBadRequest).SendString("Something went wrong")
		return err
	}

	userID := c.Params("userID")

	newUserID, err := strconv.Atoi(userID)
	if err == nil {
		anime.UserID = newUserID
	}

	var existingAnime models.AnimeList
	if err := initializers.DB.Where("user_id = ? AND title = ?", userID, anime.Title).First(&existingAnime).Error; err == nil {
		c.Status(fiber.StatusConflict).SendString("Anime already added")
		return err
	}

	if anime.Rating > 10 || anime.Rating < 0 {
		return c.Status(fiber.StatusBadRequest).SendString("Rating must be between 0 and 10")
	}

	if err := initializers.DB.Create(&anime).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
		return err
	}

	return c.JSON(anime)
}

func DeleteAnime(c *fiber.Ctx) error {
	id := c.Params("id")

	var anime models.AnimeList

	if err := initializers.DB.First(&anime, id).Error; err != nil {
		c.Status(fiber.StatusNotFound).SendString("Anime not found")
		return err
	}

	// if err := initializers.DB.Where("title = ?", anime.Title).Error; err != nil {
	// 	c.Status(fiber.StatusNotFound).SendString("Anime not found in user list")
	// 	return err
	// }

	if err := initializers.DB.Delete(&anime).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("Internal Server error")
		return err
	}

	return c.Status(fiber.StatusAccepted).SendString("Anime deleted successfully")
}

func UpdateAnime(c *fiber.Ctx) error {
	id := c.Params("id")

	var anime models.AnimeList

	if err := initializers.DB.First(&anime, id).Error; err != nil {
		c.Status(fiber.StatusNotFound).SendString("Anime not found")
		return err
	}

	var updatedAnime models.AnimeList

	if err := c.BodyParser(&updatedAnime); err != nil {
		c.Status(fiber.StatusBadRequest).SendString("Something went wrong")
		return err
	}

	if updatedAnime.Finished != anime.Finished {
		anime.Finished = updatedAnime.Finished
	}

	if updatedAnime.Rating != anime.Rating {
		anime.Rating = updatedAnime.Rating
	}

	if updatedAnime.Comment != anime.Comment {
		anime.Comment = updatedAnime.Comment
	}

	if updatedAnime.Episodes != anime.Episodes {
		anime.Episodes = updatedAnime.Episodes
	}

	if updatedAnime.PlanToWatch != anime.PlanToWatch {
		anime.PlanToWatch = updatedAnime.PlanToWatch
	}

	if updatedAnime.StartingDate != anime.StartingDate {
		anime.StartingDate = updatedAnime.StartingDate
	}

	if updatedAnime.EndingDate != anime.EndingDate {
		anime.EndingDate = updatedAnime.EndingDate
	}

	if err := initializers.DB.Save(&anime).Error; err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).SendString("Anime updated")
}

func GetAnime(c *fiber.Ctx) error {
	userID := c.Params("userID")

	var anime []models.AnimeList

	if err := initializers.DB.Where("user_id = ?", userID).Find(&anime).Error; err != nil {
		c.Status(fiber.StatusNotFound).SendString("Anime not found")
		return err
	}

	return c.JSON(anime)
}
