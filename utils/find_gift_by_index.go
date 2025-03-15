package utils

import "drop_ball/models"

func FindGiftByIndex(gifts []models.Gift, targetIndex int) models.Gift {
	for _, gift := range gifts {
		if gift.Index == targetIndex {
			return gift
		}

		continue
	}

	return models.Gift{}
}
