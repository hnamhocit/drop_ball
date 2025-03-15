package handlers

import (
	"drop_ball/models"
	"drop_ball/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RewardRepo struct {
	DB *gorm.DB
}

type CreateRewardDTO struct {
	DisplayName string `json:"display_name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

func (r *RewardRepo) CreateReward(c *gin.Context) {
	uin, ok := c.Get("uin")
	if !ok {
		c.JSON(401, types.Response{
			Code: 0,
			Msg:  "UIN is required!",
			Data: nil,
		})
		return
	}

	var input CreateRewardDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, types.Response{
			Code: 0,
			Msg:  "Can not bind json data!",
			Data: nil,
		})
		return
	}

	reward := models.Reward{
		DisplayName: &input.DisplayName,
		Address:     &input.Address,
		PhoneNumber: &input.PhoneNumber,
		Uin:         uin.(string),
	}
	if result := r.DB.Create(&reward); result.Error != nil {
		c.JSON(500, types.Response{
			Code: 0,
			Msg:  "Create reward error!",
		})
		return
	}

	c.JSON(201, types.Response{
		Code: 1,
		Msg:  "Create reward successfully!",
		Data: reward,
	})
}
