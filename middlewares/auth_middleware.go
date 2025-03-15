package middlewares

import (
	"drop_ball/models"
	"drop_ball/types"
	"drop_ball/utils"
	"errors"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")

		if len(authorization) == 0 {
			c.AbortWithStatusJSON(401, types.Response{
				Code: 0,
				Msg:  "Authorization is required!",
				Data: nil,
			})
			return
		}

		key := os.Getenv("KEY")
		if len(key) == 0 {
			c.AbortWithStatusJSON(401, types.Response{
				Code: 0,
				Msg:  "Key not found!",
				Data: nil,
			})
			return
		}

		uin := utils.Decrypt(authorization, key)
		var user models.User
		if result := db.Where("uin = ?", uin).First(&user); result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				log.Println("[USER] not found, create new...")
				db.Create(&models.User{Uin: uin, BallCount: 10})
			} else {
				c.AbortWithStatusJSON(500, types.Response{
					Code: 0,
					Msg:  "Find user error!",
					Data: nil,
				})
				return
			}
		}

		c.Set("uin", uin)
		c.Next()
	}
}
