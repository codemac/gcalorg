package main

import (
	"log"
	"strings"
	"time"

	"github.com/apognu/gocal"
	"google.golang.org/api/calendar/v3"
)

type Calendar struct {
	Id          string
	Name        string
	Description string
	Tag         string
}

func NewCalendarFromGcal(e *calendar.CalendarListEntry, tag string) *Calendar {
	return &Calendar{
		Id:          e.Id,
		Name:        e.Summary,
		Description: e.Description,
		Tag:         tag,
	}
}

func NewCalendarFromIcs(filename string, tag string) *Calendar {
	return &Calendar{
		Id:          filename,
		Name:        "ICS: " + filename,
		Description: "",
		Tag:         tag,
	}
}

type Attachment struct {
	Title   string
	FileUrl string
}

type ResponseStatus int

const (
	ResponseUnknown = iota
	ResponseAccepted
	ResponseTentative
	ResponseDeclined
)

type Attendee struct {
	ResponseStatus ResponseStatus
	Email          string
	Name           string
}

type Event struct {
	UID           string
	Name          string
	Description   string
	Attachments   []*Attachment
	Attendees     []*Attendee
	URL           string
	OrganizerMail string
	Start         *time.Time
	End           *time.Time
}

func convertAttendeeResponse(response string) ResponseStatus {
	switch strings.ToLower(response) {
	case "accepted":
		return ResponseAccepted

	case "declined":
		return ResponseDeclined

	case "tentative":
		return ResponseTentative

	case "needs-action":
		fallthrough
	case "needsaction":
		fallthrough
	default:
		return ResponseUnknown
	}
}

func NewEventFromGcal(e *calendar.Event) *Event {

	// parse attachments
	attachments := make([]*Attachment, 0)
	for _, a := range e.Attachments {
		attachments = append(attachments, &Attachment{
			Title:   a.Title,
			FileUrl: a.FileUrl,
		})
	}

	// parse attendees
	attendees := make([]*Attendee, 0, len(e.Attendees))
	for _, a := range e.Attendees {
		attendees = append(attendees, &Attendee{
			ResponseStatus: convertAttendeeResponse(a.ResponseStatus),
			Email:          a.Email,
			Name:           a.DisplayName,
		})
	}

	// parse times
	var start, end time.Time
	if e.Start.Date != "" {
		var err error
		start, err = time.Parse("2006-01-02", e.Start.Date)
		if err != nil {
			log.Fatalf("could not parse time %s", e.Start.Date)
		}
		end, err = time.Parse("2006-01-02", e.End.Date)
		if err != nil {
			log.Fatalf("could not parse time %s", e.End.Date)
		}
	} else {
		ts, err := time.Parse(time.RFC3339, e.Start.DateTime)
		if err != nil {
			log.Fatalf("could not parse time %s", e.Start.DateTime)
		}

		start = ts.In(time.Local)
		if e.End == nil {
			end = start
		} else {
			te, err := time.Parse(time.RFC3339, e.End.DateTime)
			if err != nil {
				log.Fatalf("could not parse time %s", e.Start.DateTime)
			}

			end = te.In(time.Local)
		}
	}

	return &Event{
		UID:           e.ICalUID,
		Name:          e.Summary,
		Description:   e.Description,
		Attachments:   attachments,
		Attendees:     attendees,
		URL:           e.HtmlLink,
		OrganizerMail: e.Organizer.Email,
		Start:         &start,
		End:           &end,
	}
}

func NewEventFromIcs(e *gocal.Event) *Event {
	attachments := make([]*Attachment, 0, len(e.Attachments))
	for _, a := range e.Attachments {
		attachments = append(attachments, &Attachment{
			Title:   a.Filename,
			FileUrl: a.Value,
		})
	}

	attendees := make([]*Attendee, 0, len(e.Attendees))
	for _, a := range e.Attendees {
		attendees = append(attendees, &Attendee{
			ResponseStatus: convertAttendeeResponse(a.Status),
			Name:           a.Cn,
			Email:          a.Value,
		})
	}
	var orgmail string
	if e.Organizer != nil {
		orgmail = strings.TrimPrefix(e.Organizer.Value, "mailto:")
	}

	description := strings.ReplaceAll(e.Description, "\\n", "\n")

	return &Event{
		UID:           e.Uid,
		Name:          e.Summary,
		Description:   description,
		Attachments:   attachments,
		Attendees:     attendees,
		URL:           e.URL,
		OrganizerMail: orgmail,
		Start:         e.Start,
		End:           e.End,
	}
}
