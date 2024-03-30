package models

type AnimeList struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `json:"title"`
	Finished bool   `json:"finished"`
	Rating   int    `json:"rating"`
}
