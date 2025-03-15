package models

type User struct {
	Model
	Uin       string    `json:"uin" gorm:"unique;not null"`
	BallCount int       `json:"ball_count"`
	Rewards   []*Reward `json:"rewards" gorm:"foreignKey:Uin"`
}
