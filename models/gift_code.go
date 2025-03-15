package models

import (
	"time"
)

type GiftCode struct {
	Model
	Code           string    `json:"code"`
	ExpiredDate    time.Time `json:"expired_date"`
	RemainingCount int       `json:"remaining_count"`
	Rewards        []Reward  `json:"rewards" gorm:"many2many:giftcodes_rewards"`
}
