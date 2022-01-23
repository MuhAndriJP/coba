package main

import (
	"coba/domain"

	"github.com/jasonlvhit/gocron"
)

func runCron(d *domain.Domain) {

	gocron.Every(10).Minute().Do(func() {
		// TODO: do cron/periodic action here
		//d.CalculateRank()
	})
	gocron.Start()
}
