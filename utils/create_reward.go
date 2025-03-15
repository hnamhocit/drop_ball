package utils

import (
	"errors"
	"log"

	"drop_ball/models"

	"gorm.io/gorm"
)

func GetGiftByIndex(gifts []models.Gift, targetIndex int) *models.Gift {
	for _, gift := range gifts {
		if gift.Index == targetIndex {
			return &gift
		}
	}
	return nil
}

func GetRewardCountByGiftId(rewards []models.Reward, giftId uint) int {
	count := 0
	for _, reward := range rewards {
		if *reward.GiftId == giftId {
			count += *reward.Count
		}
	}
	return count
}

// CreateReward creates rewards based on the results map
func CreateReward(db *gorm.DB, results map[int]interface{}, gifts []models.Gift, user *models.User) error {
	var rewards []models.Reward
	if result := db.Find(&rewards); result.Error != nil {
		return errors.New("Failed to fetch rewards: " + result.Error.Error())
	}

	for key, value := range results {
		switch key {
		// case 1, 2, 3:
		// 	quantity, ok := value.(int)
		// 	if !ok {
		// 		log.Println("Type assertion failed for key", key)
		// 		continue
		// 	}

		// 	if quantity > 0 {
		// 		gift := GetGiftByIndex(gifts, key)
		// 		if gift == nil {
		// 			log.Printf("Gift with index %d not found\n", key)
		// 			continue
		// 		}

		// 		rewardCount := GetRewardCountByGiftId(rewards, uint(gift.Id))
		// 		if rewardCount > gift.MaxCount {
		// 			log.Printf("Exceeding max count for gift with index %d\n", key)
		// 			continue
		// 		}

		// 		newReward := models.Reward{
		// 			UserUin: user.Uin,
		// 			Count:   quantity,
		// 			GiftId:  gift.Id,
		// 		}

		// 		if err := db.Create(&newReward).Error; err != nil {
		// 			return errors.New("Failed to create reward: " + err.Error())
		// 		}
		// 	}

		// 	break
		case 4:
			giftCodes := results[4].([]string)
			giftCodeSlice := []*models.GiftCode{}

			for _, code := range giftCodes {
				giftCodeSlice = append(giftCodeSlice, &models.GiftCode{
					Code: code,
				})
			}

			reward := models.Reward{
				Uin:       user.Uin,
				GiftCodes: giftCodeSlice,
			}

			if err := db.Create(&reward).Error; err != nil {
				log.Println(err.Error())
				return errors.New("Failed to create reward: " + err.Error())
			}

		case 5:
			ballCount, ok := value.(int)
			if !ok {
				log.Println("Type assertion failed for key 5")
				continue
			}

			user.BallCount += ballCount
			if err := db.Save(&user).Error; err != nil {
				return errors.New("Failed to update user ball count: " + err.Error())
			}
		default:

		}
	}

	return nil
}
