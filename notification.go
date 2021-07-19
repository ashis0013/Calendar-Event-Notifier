package main

import "github.com/deckarep/gosx-notifier"

func notify(text string) {
	note := gosxnotifier.NewNotification("Join the google meet")
    note.Title = "Time for Meeting"
    note.Subtitle = text
    note.Sound = gosxnotifier.Default
	note.Sender = "com.apple.Safari"
	note.Push()	
}