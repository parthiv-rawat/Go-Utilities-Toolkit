package timeutils

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func StartCronJob() {
	c := cron.New()
	c.AddFunc("@every 10s", func() {
		fmt.Println("Running every 10 seconds", time.Now())
	})
	c.Start()
}
