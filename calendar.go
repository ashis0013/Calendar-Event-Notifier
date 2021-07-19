package main

import (
	"io/ioutil"
	"net/http"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
	"fmt"
	"context"
	"log"
	"os"
	"encoding/json"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"time"
)

func getClient(config *oauth2.Config) *http.Client {
	tokFile := "/Users/ashis.paul/go/src/calendar-notifier/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
			tok = getTokenFromWeb(config)
			saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
			"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
			log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
			log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
			return nil, err
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
			log.Fatalf("Unable to cache OAuth token: %v", err)
	}
	json.NewEncoder(f).Encode(token)
}

func authGoogleApi() *http.Client{
	b, err := ioutil.ReadFile("/Users/ashis.paul/go/src/calendar-notifier/credentials.json")
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
			fmt.Println("Unable to parse client secret file to config:")
	}
	return getClient(config)
}

var events *calendar.Events

func fetchEvents() {
	client := authGoogleApi()
	ctx := context.Background()
	s, _ := calendar.NewService(ctx, option.WithHTTPClient(client))
	t := time.Now().Format(time.RFC3339)
	te := time.Now().AddDate(0,0,1).Format(time.RFC3339)
    events,_ = s.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(t).TimeMax(te).MaxResults(10).OrderBy("startTime").Do()
}

func fetchEventsRoutine() {
	for {
		fetchEvents()
		time.Sleep(time.Second)
	}
}
