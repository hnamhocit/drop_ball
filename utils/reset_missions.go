package utils

import (
	"drop_ball/models"
	"log"
	"time"

	"gorm.io/gorm"
)

func ResetMissions(db *gorm.DB) {
	log.Println("Running mission reset at:", time.Now())

	// Get all users
	var users []models.User
	if err := db.Preload("Missions", "deleted_at IS NULL").Find(&users).Error; err != nil {
		log.Printf("failed to fetch users: %v", err)
		return
	}

	for _, user := range users {
		// If the user has no missions, create 5 default missions
		if len(user.Missions) == 0 {
			log.Printf("No missions found for user %s. Creating 5 default missions...\n", user.Uin)
			CreateDefaultMissions(db, user)
			continue
		}

		// Sort missions by CreatedAt (ascending order)
		missions := user.Missions
		SortMissionsByCreatedAt(missions)

		// Delete all missions except the last one
		for i := range len(missions) {
			if err := db.Delete(&missions[i]).Error; err != nil {
				log.Printf("failed to delete mission %d: %v", missions[i].Id, err)
			} else {
				log.Printf("Deleted mission %d for user %s\n", missions[i].Id, user.Uin)
			}
		}
	}
}

// createDefaultMissions creates 5 default missions for the given user
func CreateDefaultMissions(db *gorm.DB, user models.User) {
	defaultMissionNames := []string{"Task 1", "Task 2", "Task 3", "Task 4", "Task 5"}
	for _, name := range defaultMissionNames {
		mission := models.Mission{
			Name:       name,
			IsComplete: false,
			BallCount:  1, // Default ball count
			UserUin:    user.Uin,
		}
		if err := db.Create(&mission).Error; err != nil {
			log.Printf("failed to create mission %s for user %s: %v", name, user.Uin, err)
		} else {
			log.Printf("Created mission %s for user %s\n", name, user.Uin)
		}
	}
}

// sortMissionsByCreatedAt sorts missions by CreatedAt in ascending order
func SortMissionsByCreatedAt(missions []models.Mission) {
	for i := range len(missions) {
		for j := i + 1; j < len(missions); j++ {
			if missions[i].CreatedAt.After(missions[j].CreatedAt) {
				missions[i], missions[j] = missions[j], missions[i]
			}
		}
	}
}
