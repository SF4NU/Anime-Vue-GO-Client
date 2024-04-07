package models

type MangaList struct {
	ID           uint    `gorm:"primaryKey"`
	Title        string  `json:"title"`
	Finished     bool    `json:"finished"`
	Chapters     int     `json:"chapters"`
	Comment      string  `json:"comment"`
	PlanToRead   bool    `json:"plan_to_read"`
	StartingDate string  `json:"starting_date"`
	EndingDate   string  `json:"ending_date"`
	Rating       float64 `json:"rating"`
	MangaID      int     `json:"manga_id"`
	UserID       int     `json:"user_id"`
}
