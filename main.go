package main

import (
	"flag"
	"fmt"
	"time"
)

func manageFlags() {
	dayRange = flag.Int("d", 1, "Number of days upto which events will be fetched")
	minRange = flag.Int("m", 0, "Number of minutes upto which events will be fetched. You can only use one flag.")
	flag.Parse()
}

func main() {
	manageFlags()
	getGlobals()
	fetchEvents()
	if !fetchEventSuccess {
		return
	}
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
