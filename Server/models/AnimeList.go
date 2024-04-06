package models

type AnimeList struct {
	ID           uint    `gorm:"primaryKey"`
	Title        string  `json:"title"`
	Finished     bool    `json:"finished"`
	Episodes     int     `json:"episodes"`
	Comment      string  `json:"comment"`
	PlanToWatch  bool    `json:"plan_to_watch"`
	StartingDate string  `json:"starting_date"`
	EndingDate   string  `json:"ending_date"`
	Rating       float64 `json:"rating"`
	AnimeID      int     `json:"anime_id"`
	UserID       int     `json:"user_id"`
}
