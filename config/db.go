package config

import (
	"drop_ball/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDefaultGifts(db *gorm.DB) {
	var giftsLength int64
	db.Model(&models.Gift{}).Count(&giftsLength)

	if giftsLength == 0 {
		log.Println("[GIFT] does not exists, create new...")

		db.Create(&models.Gift{Name: "Model", MaxCount: 3, Index: 1})
		db.Create(&models.Gift{Name: "Skin VIP", MaxCount: 10, Index: 2})
		db.Create(&models.Gift{Name: "Skin DIY", MaxCount: 30, Index: 3})
	}

	log.Println("[GIFT] has initialize successfully!")
}

// InitDB initializes the database connection
func InitDB(config *Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	if err := db.AutoMigrate(&models.User{}, &models.Gift{}, &models.Reward{}, &models.GiftCode{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	CreateDefaultGifts(db)
	log.Println("Database connection established")

	return db
}
