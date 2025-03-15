package models

type Reward struct {
	Model
	Count       *int        `json:"count"`
	PhoneNumber *string     `json:"phone_number"`
	Address     *string     `json:"address"`
	DisplayName *string     `json:"display_name"`
	GiftId      *uint       `json:"gift_id"`
	Uin         string      `json:"uin"`
	GiftCodes   []*GiftCode `json:"gift_codes" gorm:"many2many:giftcodes_rewards"`
}
