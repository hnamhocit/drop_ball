package config

import (
	"crypto/rand"
	"drop_ball/handlers"
	"drop_ball/middlewares"
	"drop_ball/types"
	"drop_ball/utils"
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

	r.Use(middlewares.AuthMiddleware(db))

	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			userRepo := handlers.UserRepo{DB: db}
			users.GET("/me", userRepo.GetUser)
		}

		randoms := api.Group("/randoms")
		{
			randomRepo := handlers.RandomRepo{DB: db}
			randoms.POST("", randomRepo.Randoms)
		}

		giftCodes := api.Group("/giftcodes")
		{
			giftCodeRepo := handlers.GiftCodeRepo{DB: db}
			giftCodes.POST("", giftCodeRepo.CreateGiftCode)
			giftCodes.GET("", giftCodeRepo.GetGiftCodes)
		}

		rewards := api.Group("/rewards")
		{
			rewardRepo := handlers.RewardRepo{DB: db}
			rewards.POST("", rewardRepo.CreateReward)
		}
	}
}
