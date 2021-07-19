package main

import (
	"fmt"
	"time"
)

func main() {
	getGlobals()
	fetchEvents()
	fmt.Println("Upcoming Events:")
	for _, item := range events.Items {
		date := item.Start.DateTime
		fmt.Printf("%v (%v)\n", item.Summary, date)
	}

	go fetchEventsRoutine()
	for {
		for _, item := range events.Items {
			curTime := time.Now().Add(time.Minute * 3).Format(time.RFC3339)
			date := item.Start.DateTime
			if curTime == date {
				notify(item.Summary)
			}
		}
		time.Sleep(time.Second)
	}
}
