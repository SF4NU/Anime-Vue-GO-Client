package initializers

import (
	"github.com/sf4nu/Anime-Vue-GO-Client/models"
)

func AutoMigrate() {
	DB.AutoMigrate(&models.User{}, &models.AnimeList{}, &models.MangaList{})
}
