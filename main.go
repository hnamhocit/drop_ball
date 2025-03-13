package main

import (
	"drop_ball/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	cfg := config.LoadConfig()
	db := config.InitDB(cfg)

	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	config.LoadRoutes(r, db)
	config.SetResetSchedule(db)

	r.Run()
}
