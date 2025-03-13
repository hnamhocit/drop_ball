package middlewares

import (
	"drop_ball/types"
	"drop_ball/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")

		if len(authorization) == 0 {
			c.JSON(401, types.Response{
				Code: 0,
				Msg:  "Authorization is required!",
				Data: nil,
			})
			return
		}

		key := os.Getenv("KEY")
		if len(key) == 0 {
			c.JSON(401, types.Response{
				Code: 0,
				Msg:  "Key not found!",
				Data: nil,
			})
		}

		uin := utils.Decrypt(authorization, key)
		c.Set("uin", uin)

		c.Next()
	}
}
