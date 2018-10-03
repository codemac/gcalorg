package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"google.golang.org/api/calendar/v3"
)

func printCalendars(client *http.Client, approvedCals []string, tagname string) {
	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve calendar Client %v", err)
	}

	// find all calendars
	calendars, err := srv.CalendarList.List().ShowHidden(false).ShowDeleted(false).
		MaxResults(250).Do()
	if err != nil {
		log.Fatalf("Unable to list calendars! %v", err)
	}

	curtime := time.Now().UTC().Add(24 * time.Hour).Truncate(24 * time.Hour)
	timeMin := curtime.AddDate(0, -9, 0).Format("2006-01-02T15:04:05Z")
	timeMax := curtime.AddDate(1, 0, 0).Format("2006-01-02T15:04:05Z")

	receivedCals := make(map[string]*calendar.CalendarListEntry, 0)
	for _, c := range calendars.Items {
		receivedCals[c.Id] = c
	}
	fmt.Printf("# -*- eval: (auto-revert-mode 1); -*-\n")
	fmt.Printf("#+category: cal\n")
	for _, approvedCal := range approvedCals {

		c, ok := receivedCals[approvedCal]
		if !ok {
			continue
		}
		fmt.Printf("* %s :%s:\n", noTodoKwds(c.Summary), tagname)
		fmt.Printf("  :PROPERTIES:\n")
		fmt.Printf("  :ID:         %s\n", c.Id)
		fmt.Printf("  :END:\n")
		fmt.Printf("\n%s\n\n", c.Description)

		npt := ""
		notdone := true

		event_list := make([]*calendar.Event, 0, 250)
		for notdone {
			eventsReq := srv.Events.List(c.Id).ShowDeleted(false).
				SingleEvents(true).TimeMin(timeMin).TimeMax(timeMax).MaxResults(250)
			if npt != "" {
				eventsReq = eventsReq.PageToken(npt)
				npt = ""
			}

			events, err := eventsReq.Do()
			if err != nil {
				log.Fatalf("Unable to retrieve next ten of the user's events. %v", err)
			}

			notdone = events.NextPageToken != ""
			if notdone {
				npt = events.NextPageToken
			}

			event_list = append(event_list, events.Items...)
		}

		events_by_id := make(map[string][]*calendar.Event)
		for _, v := range event_list {
			recur_id := strings.Split(v.ICalUID, "_R")[0]
			events_by_id[recur_id] = append(events_by_id[recur_id], v)
		}

		type eventWithId struct {
			id string
			events []*calendar.Event
		}

		// sorted events
		sorted_by_id := make([]eventWithId, 0, len(events_by_id))
		for id, events := range events_by_id {
			sorted_by_id = append(sorted_by_id,
				eventWithId{id, events})
		}

		sort.Slice(sorted_by_id, func (i, j int) bool {
			return sorted_by_id[i].id < sorted_by_id[j].id
		})

		for _, e := range sorted_by_id {
			events := e.events
			if len(events) == 0 {
				continue
			}
			// skip things that are chatty (repeating calendar
			// events -> org-mode has been difficult, manually
			// manage those for now. There is probably a way of
			// getting them, but converting the ical format to the
			// org format would be a significant piece of logic)
			if filteredEvent(c.Id, events[0].Summary) {
				continue
			}

			fmt.Println(fmtEventGroup(events))
		}
	}
}

func main() {
	home := os.Getenv("HOME")
	type caldata struct {
		name    string
		cals    []string
		tagname string
	}
	secrets := []caldata{
		{
			home + "/go/src/github.com/codemac/gcalorg/jmickeygoogle_secret.json",
			workCals,
			"WORK",
		},
		{
			home + "/go/src/github.com/codemac/gcalorg/codemacgmail_secret.json",
			gmailCals,
			"HOME",
		},
	}

	// we need to sort before we do much of anything, so things show up in a
	// decent order.
	for _, v := range secrets {
		fmt.Fprintf(os.Stderr, "Getting client for: %s", v.name)
		cl := genClient(v.name)
		printCalendars(cl, v.cals, v.tagname)
	}
}
