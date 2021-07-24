package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var credLoc string
var tokLoc string
var fetchSleep int
var fetchEventSuccess bool
var dayRange, minRange *int

func getGlobals() {
	file, err := os.Open("config.cfg")
	if err != nil {
		fmt.Println("Unable to open config")
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	m := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		m[strings.Split(line, "=")[0]] = strings.Split(line, "=")[1]
	}
	credLoc = m["CRED_LOC"]
	tokLoc = m["TOK_LOC"]
	fetchSleep, _ = strconv.Atoi(m["FETCH_SLEEP"])
}
