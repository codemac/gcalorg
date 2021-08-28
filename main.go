package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/api/calendar/v3"
)

func main() {
	home := os.Getenv("HOME")
	type caldata struct {
		name    string
		cals    []string
		tagname string
	}

	secrets := []caldata{}
	icsDirs := []string{home + "/Library/Calendars"}

	curtime := time.Now().UTC().Add(24 * time.Hour).Truncate(24 * time.Hour)
	timeMin := curtime.AddDate(0, -9, 0)
	timeMax := curtime.AddDate(1, 0, 0)

	calendar_events := make(map[*Calendar][]*Event, 0)
	for _, v := range secrets {
		fmt.Fprintf(os.Stderr, "Getting client for: %s", v.name)
		cl := genClient(v.name)
		srv, err := calendar.New(cl)
		if err != nil {
			log.Fatalf("Unable to retrieve calendar Client %v", err)
		}

		gcals := gcalCals(srv, v.cals, v.tagname)
		for _, c := range gcals {
			calendar_events[c] = gcalEvents(srv, c, timeMin, timeMax)
		}

	}

	icscals := icsCals(icsDirs)
	for _, c := range icscals {
		calendar_events[c] = icsEvents(c, timeMin, timeMax)
	}

	// format calendar file output
	fmtFile()
	for c, es := range calendar_events {
		fmtCalendar(c)
		fmtEvents(c, es)
	}
}
