package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
