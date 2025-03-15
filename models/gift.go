package models

type Gift struct {
	Model
	Name     string   `json:"name"`
	Index    int      `json:"index"`
	MaxCount int      `json:"max_count"`
	Rewards  []Reward `json:"rewards" gorm:"foreignKey:GiftId"`
}
