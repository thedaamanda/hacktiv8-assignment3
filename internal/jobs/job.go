package jobs

import (
	"assignment3/internal/handlers"
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("Starting scheduler...")
	s := gocron.NewScheduler()
	s.Every(15).Seconds().Do(handlers.GetDataStatus)
	<-s.Start()
}

func StartScheduler() {
	go task()
}
