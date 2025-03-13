package config

import (
	"crypto/rand"
	"drop_ball/middlewares"
	"drop_ball/models"
	"drop_ball/types"
	"drop_ball/utils"
	"errors"
	"log"
	"math/big"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoadRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/random-uin", func(c *gin.Context) {
		key := os.Getenv("KEY")
		max := big.NewInt(1000000)
		randomNumber, err := rand.Int(rand.Reader, max)

		if err != nil {
			c.JSON(500, types.Response{
				Code: 0,
				Msg:  "Error generating random number:" + err.Error(),
			})
			return
		}

		encoding := utils.Encrypt("user_"+randomNumber.String(), key)

		c.JSON(200, types.Response{
			Code: 1,
			Msg:  "Random UIN successfully!",
			Data: gin.H{
				"uin": encoding,
			},
		})
	})

	r.Use(middlewares.AuthMiddleware())

	r.GET("/", func(c *gin.Context) {
		uin, ok := c.Get("uin")

		if !ok {
			c.JSON(401, types.Response{
				Code: 0,
				Msg:  "UIN is invalid!",
				Data: nil,
			})
		}

		var user models.User

		result := db.Preload("Missions").First(&user, "uin = ?", uin)

		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				uinValue, ok := uin.(string)
				if !ok {
					c.JSON(400, types.Response{
						Code: 0,
						Msg:  "Invalid UIN type",
					})
					return
				}

				log.Println("[USER] not found, create new...")

				db.Create(&models.User{
					Uin:       uinValue,
					BallCount: 10,
				})

				db.Preload("Missions").First(&user, "uin = ?", uinValue)
			} else {
				c.JSON(500, types.Response{
					Code: 0,
					Msg:  "Database error",
				})
				return
			}
		}

		missionMap := make(map[string]bool)
		for index, mission := range user.Missions {
			missionMap[string(rune(index))] = mission.IsComplete
		}

		c.JSON(200, types.Response{
			Code: 1,
			Msg:  "Get user info successfully!",
			Data: gin.H{
				"uin":        user.Uin,
				"ball_count": user.BallCount,
				"missions":   missionMap,
			},
		})
	})
}
