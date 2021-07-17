package main

import (
    "github.com/deckarep/gosx-notifier"
	"fmt"
	"context"
	"time"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
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
	client := authGoogleApi()
	ctx := context.Background()
	s, _ := calendar.NewService(ctx, option.WithHTTPClient(client))
	t := time.Now().Format(time.RFC3339)
	te := time.Now().AddDate(0,0,3).Format(time.RFC3339)
	fmt.Println(t)
    events, _ := s.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(t).TimeMax(te).MaxResults(10).OrderBy("startTime").Do()
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
			fmt.Println("No upcoming events found.")
	} else {
			for _, item := range events.Items {
					date := item.Start.DateTime
					layout := "2006-01-02T15:04:05Z07:00"
					tt, _ := time.Parse(layout, date)
					fmt.Println(tt)
			}
	}
	for {
		for _, item := range events.Items {
			curTime := time.Now().Add(time.Minute * 3).Format(time.RFC3339)
			date := item.Start.DateTime
			if curTime == date {
				notify(item.Summary)
			}
		}
	}
	//TODO : Add multithreading to refresh events timely
}
