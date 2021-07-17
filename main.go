package main

import (
    "github.com/deckarep/gosx-notifier"
	"fmt"
	"time"
)

func notify(text string) {
	note := gosxnotifier.NewNotification("Join the google meet")
    note.Title = "Time for Meeting"
    note.Subtitle = text
    note.Sound = gosxnotifier.Default
	note.Sender = "com.apple.Safari"
	note.Push()	
}



func main() {
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
	}
}
