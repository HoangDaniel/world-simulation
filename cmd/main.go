package main

import (
	"fmt"
	"hoangdaniel/world/cmd/timeservice"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	startTime := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
	ts := timeservice.NewTimeService(startTime)
	ts.SetSpeed(60)
	fmt.Println("New Age started with date time: %v", ts.GetCurrentTime())

	for range ticker.C {
		fmt.Println(ts.GetCurrentTime())
	}
}
