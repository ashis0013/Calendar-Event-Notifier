package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
)

var credLoc string
var tokLoc string
var fetchSleep int

func getGlobals() {
	file, _ := os.Open("/Users/ashis.paul/go/src/calendar-notifier/config.cfg")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	m := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		m[strings.Split(line,"=")[0]] = strings.Split(line,"=")[1]
	}
	credLoc = m["CRED_LOC"]
	tokLoc = m["TOK_LOC"]
	fetchSleep, _ = strconv.Atoi(m["FETCH_SLEEP"])
}