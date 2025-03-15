package handlers

import (
	"drop_ball/models"
	"drop_ball/types"
	"drop_ball/utils"
	"log"
	"math/rand"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RandomRepo struct {
	DB *gorm.DB
}

type RandomDTO struct {
	Gate []int `json:"gate"`
}

func (r *RandomRepo) Randoms(c *gin.Context) {
	uin, ok := c.Get("uin")
	if !ok {
		c.JSON(500, types.Response{
			Code: 0,
			Msg:  "UIN does not exist!",
			Data: nil,
		})
		return
	}

	var user models.User
	result := r.DB.First(&user, "uin = ?", uin)

	if result.Error != nil {
		c.JSON(404, types.Response{Code: 0, Msg: "User not found!", Data: nil})
		return
	}

	if user.BallCount <= 0 {
		c.JSON(500, types.Response{
			Code: 0,
			Msg:  "Insufficient ball count",
			Data: nil,
		})
		return
	}

	user.BallCount -= 6
	if err := r.DB.Save(&user).Error; err != nil {
		c.JSON(500, types.Response{
			Code: 0,
			Msg:  "Failed to update user ball count!",
			Data: nil,
		})
		return
	}

	var randomDto RandomDTO
	if err := c.ShouldBindJSON(&randomDto); err != nil {
		c.JSON(400, types.Response{
			Code: 0,
			Msg:  "Invalid JSON input: " + err.Error(),
			Data: nil,
		})
		return
	}

	isValidGate := true
	for _, gateValue := range randomDto.Gate {
		if gateValue < 1 || gateValue > 6 {
			log.Println("Invalid: ", gateValue)
			isValidGate = false
			break
		}
	}

	if !isValidGate {
		c.JSON(400, types.Response{
			Code: 0,
			Msg:  "Invalid gate values! Gate values must be between 1 and 6.",
			Data: nil,
		})
		return
	}

	results := map[int]interface{}{
		1: 0,
		2: 0,
		3: 0,
		4: []string{},
		5: 0,
	}

	gate := randomDto.Gate
	gateLength := len(gate)

	var gifts []models.Gift
	giftResult := r.DB.Find(&gifts)

	if giftResult.Error != nil {
		c.JSON(500, types.Response{
			Code: 0,
			Msg:  "Get gifts error!",
			Data: nil,
		})
		return
	}

	/**
	Gifts
	1. Mô hình, số lượng 3
	2. Skin VIP, số lượng 10
	3. Skin DIY, số lượng 30
	4. Giftcode, số lượng 1000
	5. Thêm bóng
	6. Chúc bạn may mắn lần sau

	6 results & 6 gates
	Gate: []gift
	1: [1, 2, 3, 4, 5, 6]
	2: [2, 3, 4, 5, 6]
	3: [3, 4, 5, 6],
	4: [4, 5, 6]
	5: [5, 6]
	6: [6] skip
	*/

	for i := range gateLength {
		if i == 6 {
			continue
		}

		// create weights per gate
		weights := make(map[int]int)
		for j := i; j < gateLength; j++ {
			// j == 6 => 0%, skipped for this case
			if j == 6 {
				break
			}

			if j == 1 {
				weights[j] = 10
			} else {
				weights[j] = 30
			}
		}

		selectedGate := utils.WeightedRandomSelector(weights)
		switch selectedGate {
		case 1:
		case 2:
		case 3:
			if value, ok := results[selectedGate].(int); ok {
				results[selectedGate] = value + 1
			}

			break
		case 4:
			var giftCodes []models.GiftCode
			result := r.DB.Find(&giftCodes)

			if result.Error != nil {
				c.JSON(500, types.Response{
					Code: 0,
					Msg:  "Get giftcodes error!",
					Data: nil,
				})
				return
			}

			giftCodesLen := len(giftCodes)
			if giftCodesLen > 0 {
				randomIndex := rand.Intn(giftCodesLen)
				if giftSlice, ok := results[selectedGate].([]string); ok {
					results[selectedGate] = append(giftSlice, giftCodes[randomIndex].Code)
				} else {
					log.Println("Type assertion failed for key", selectedGate)
				}
			}

			break
		case 5:
			min := 2
			max := 6
			randomBallCount := rand.Intn(max-min+1) + min

			if value, ok := results[5].(int); ok {
				results[5] = value + randomBallCount
			}

			user.BallCount += randomBallCount
			if result := r.DB.Save(&user); result.Error != nil {
				c.JSON(500, types.Response{
					Code: 0,
					Msg:  "Save user failed!",
					Data: nil,
				})
				return
			}
			break
		default:
			break
		}
	}

	// err := utils.CreateReward(r.DB, map[int]interface{}{
	// 	1: 0,
	// 	2: 0,
	// 	3: 0,
	// 	4: []string{"BOCLOTSUCLAODONG", "HELLO2026"},
	// 	5: 20,
	// }, gifts, &user)
	// if err != nil {
	// 	c.JSON(500, types.Response{
	// 		Code: 0,
	// 		Msg:  "Create reward error: " + err.Error(),
	// 		Data: nil,
	// 	})
	// 	return
	// }

	c.JSON(201, types.Response{
		Code: 1,
		Msg:  "Get random gift successfully!",
		Data: results,
	})
}
