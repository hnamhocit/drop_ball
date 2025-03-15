package handlers

import (
	"drop_ball/models"
	"drop_ball/types"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GiftCodeRepo struct {
	DB *gorm.DB
}

type CreateGiftCodeDTO struct {
	Code           string `json:"code"`
	Days           int    `json:"days"`
	RemainingCount int    `json:"remaining_count"`
}

func (r *GiftCodeRepo) CreateGiftCode(c *gin.Context) {
	var input CreateGiftCodeDTO

	err := c.Bind(&input)
	if err != nil {
		c.JSON(500, types.Response{Code: 0, Msg: "Bind form data error: " + err.Error()})
		return
	}

	var count int64
	existingGiftCode := r.DB.Model(&models.GiftCode{}).Where("code = ?", input.Code).Count(&count)

	if existingGiftCode.Error != nil {
		c.JSON(500, types.Response{
			Code: 0,
			Msg:  "Database error!",
			Data: nil,
		})
		return
	}

	if count > 0 {
		c.JSON(409, types.Response{
			Code: 0,
			Msg:  "Giftcode already exists!",
		})
		return
	}

	giftCode := models.GiftCode{
		Code:           input.Code,
		ExpiredDate:    time.Now().Add(time.Duration(input.Days) * 24 * time.Hour),
		RemainingCount: input.RemainingCount,
	}

	r.DB.Create(&giftCode)

	c.JSON(201, types.Response{
		Code: 1,
		Msg:  "Create gift code successfully!",
		Data: giftCode,
	})
}

func (r *GiftCodeRepo) GetGiftCodes(c *gin.Context) {
	var giftCodes []models.GiftCode

	result := r.DB.Find(&giftCodes)
	if result.Error != nil {
		c.JSON(500, types.Response{
			Code: 0,
			Msg:  "Get giftcodes failed!",
			Data: nil,
		})
	}

	c.JSON(200, types.Response{
		Code: 1,
		Msg:  "Get giftcodes successfully!",
		Data: giftCodes,
	})
}
