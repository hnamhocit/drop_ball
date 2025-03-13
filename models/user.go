package models

type User struct {
	Model
	Uin       string    `json:"uin" gorm:"unique;not null"` // Ensure Uin is unique
	BallCount int       `json:"ball_count"`
	Missions  []Mission `json:"missions" gorm:"foreignKey:UserUin"`
}
