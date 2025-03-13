package config

import (
	"drop_ball/utils"
	"log"

	"github.com/go-co-op/gocron/v2"
	"gorm.io/gorm"
)

func SetResetSchedule(db *gorm.DB) {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		log.Fatalf("failed to create scheduler: %v", err)
	}

	job, err := scheduler.NewJob(
		gocron.CronJob("0 0 * * *", false),
		gocron.NewTask(utils.ResetMissions, db),
	)
	if err != nil {
		log.Fatalf("failed to schedule job: %v", err)
	}

	log.Printf("Scheduled job with ID: %s\n", job.ID())

	scheduler.Start()
}
