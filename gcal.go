package main

import (
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

func gcalCals(srv *calendar.Service, approvedCals []string, tagname string) []*Calendar {
	// find all calendars
	calendars, err := srv.CalendarList.List().ShowHidden(false).ShowDeleted(false).
		MaxResults(250).Do()
	if err != nil {
		log.Fatalf("Unable to list calendars! %v", err)
	}

	receivedCals := make(map[string]*Calendar, 0)
	for _, c := range calendars.Items {
		receivedCals[c.Id] = NewCalendarFromGcal(c, tagname)
	}

	calResults := make([]*Calendar, 0)
	for _, approvedCal := range approvedCals {
		c, ok := receivedCals[approvedCal]
		if ok {
			calResults = append(calResults, c)
		}
	}

	return calResults
}

func gcalEvents(srv *calendar.Service, c *Calendar, timeMin time.Time, timeMax time.Time) []*Event {
	event_list := make([]*calendar.Event, 0, 250)
	notdone := true
	npt := ""
	for notdone {
		eventsReq := srv.Events.List(c.Id).ShowDeleted(false).
			SingleEvents(true).TimeMin(timeMin.Format("2006-01-02T15:04:05Z")).
			TimeMax(timeMax.Format("2006-01-02T15:04:05Z")).MaxResults(250)
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

	events := make([]*Event, 0, 250)
	for _, v := range event_list {
		events = append(events, NewEventFromGcal(v))
	}

	return events
}
