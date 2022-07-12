package main

import (
	"time"

	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

func Job(name string) {
	log.Info(name)
}

func runCronJob() {
	schedule := cron.New()
	schedule.AddFunc("*/1 * * * *", func() {
		Job("Hello Philia, Shree Calling you...")
	})
	log.Info("Start cron after every minute")
	schedule.Start()
	time.Sleep(4 * time.Minute)
}

func main() {
	runCronJob()
}
