package handlers

import (
	"drop_ball/models"
	"drop_ball/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (r *UserRepo) GetUser(c *gin.Context) {
	uin, ok := c.Get("uin")

	if !ok {
		c.JSON(401, types.Response{
			Code: 0,
			Msg:  "UIN is invalid!",
			Data: nil,
		})
	}

	var user models.User
	r.DB.First(&user, "uin = ?", uin)

	c.JSON(200, types.Response{
		Code: 1,
		Msg:  "Get user info successfully!",
		Data: user,
	})
}
