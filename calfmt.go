package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"google.golang.org/api/calendar/v3"
)

func datesToOrg(start, end *calendar.EventDateTime) string {
	final := ""
	if start == nil { // this event has dates! hurrah!
		return "\n"
	}

	if start.Date != "" { // all day event!
		ts, _ := time.Parse("2006-01-02", start.Date)
		tsf := ts.Format("2006-01-02")
		final = final + fmt.Sprintf("<%s", tsf)
		if end == nil {
			return final + ">"
		}

		te, _ := time.Parse("2006-01-02", end.Date)
		te = te.AddDate(0, 0, -1)
		// The end date is "exclusive", so we should subtract a day, and
		// if the day is equivalent to start, then we should just print
		// start.
		if te.Equal(ts) {
			return final + ">"
		}
		tef := te.Format("2006-01-02")
		return final + fmt.Sprintf(">--<%s>", tef)
	}

	ts, _ := time.Parse(time.RFC3339, start.DateTime)
	ts = ts.In(time.Local)
	tsf := ts.Format("2006-01-02 Mon 15:04")
	final = final + fmt.Sprintf("<%s", tsf)

	if end == nil {
		return final + fmt.Sprintf(">")
	}

	te, _ := time.Parse(time.RFC3339, end.DateTime)
	te = te.In(time.Local)
	if te.Day() != ts.Day() {
		tef := te.Format("2006-01-02 Mon 15:04")
		return final + fmt.Sprintf(">--<%s>", tef)
	}
	tef := te.Format("15:04")
	return final + fmt.Sprintf("-%s>", tef)
}

func noTodoKwds(s string) string {
	states := []string{"TODO", "NEXT", "STARTED", "WAITING", "PROJECT", "DONE", "NVM"}
	for _, state := range states {
		// We essentially reimplement TrimPrefix, so we don't waste time
		// running HasPrefix multiple times. We also don't remove the
		// space in the prefix.
		if strings.HasPrefix(s, state + " ") {
			s = s[len(state):]
			s = "/" + state + "/" + s

			// Only the actual prefix matters, we don't need to
			// check the remaining keywords.
			return s
		}
	}
	return s
}

// cleanString removes special characters for org-mode, as almost no one will be
// using org-mode formatting.
func cleanString(s string) string {
	s = strings.Replace(s, "[", "{", -1)
	s = strings.Replace(s, "]", "}", -1)
	s = strings.Replace(s, "\n*", "\n,*", -1)
	return s
}

func isDeclinedByMe(e *calendar.Event) bool {
	for _, a := range e.Attendees {
		if a != nil && a.Self && a.ResponseStatus == "declined" {
			return true
		}
	}
	return false
}

func fmtOrgHeader(e *calendar.Event) string {
	var buf string
	buf += fmt.Sprintf("** ")
	if e.Status == "tenative" || e.Status == "cancelled" {
		buf += fmt.Sprintf("(%s) ", e.Status)
	}
	summary := e.Summary
	if summary == "" {
		summary = "busy"
	}

	buf += fmt.Sprintf("%s\n", noTodoKwds(summary))
	buf += fmt.Sprintf(":PROPERTIES:\n")
	buf += fmt.Sprintf(":ID:       %s\n", e.ICalUID)
	buf += fmt.Sprintf(":GCALLINK: %s\n", e.HtmlLink)
	if e.Creator != nil {
		buf += fmt.Sprintf(":CREATOR: [[mailto:%s][%s]]\n", e.Creator.Email, cleanString(e.Creator.DisplayName))
	}
	if e.Organizer != nil {
		buf += fmt.Sprintf(":ORGANIZER: [[mailto:%s][%s]]\n", e.Organizer.Email, cleanString(e.Organizer.DisplayName))
	}
	buf += fmt.Sprintf(":END:\n\n")

	return buf
}

func fmtOrgDate(e *calendar.Event) string {
	return fmt.Sprintf("%s\n", datesToOrg(e.Start, e.End))
}

func fmtOrgAttendees(e *calendar.Event) string {
	var buf string
	attendees := e.Attendees
	if len(attendees) == 0 {
		return ""
	}

	canonical_id := func(ea *calendar.EventAttendee) string {
		if ea.Id != "" {
			return ea.Id
		} else if ea.Email != "" {
			return ea.Email
		} else if ea.DisplayName != "" {
			return cleanString(ea.DisplayName)
		}
		return "sadness"
	}

	sort.SliceStable(attendees, func(i, j int) bool {
		return canonical_id(attendees[i]) < canonical_id(attendees[j])
	})

	if len(attendees) > 20 {
		buf += fmt.Sprintf("Attendees: ... Many\n")
		return buf
	}

	buf += fmt.Sprintf("Attendees:\n")
	for _, a := range attendees {
		if a == nil {
			continue
		}

		// ResponseStatus: The attendee's response status. Possible values are:
		//
		// - "needsAction" - The attendee has not responded to the invitation.
		//
		// - "declined" - The attendee has declined the invitation.
		// - "tentative" - The attendee has tentatively accepted the invitation.
		//
		// - "accepted" - The attendee has accepted the invitation.
		//  ResponseStatus string `json:"responseStatus,omitempty"`
		statuschar := " "
		switch a.ResponseStatus {
		case "":
		case "NeedsAction":
		case "declined":
			statuschar = "✗"
		case "tenative":
			statuschar = "☐"
		case "accepted":
			statuschar = "✓"
		}

		linkname := cleanString(a.DisplayName)
		if linkname == "" {
			linkname = a.Email
		}
		buf += fmt.Sprintf(" %s [[mailto:%s][%s]]\n", statuschar, a.Email, linkname)
	}
	return buf
}

func fmtOrgBody(e *calendar.Event) string {
	var buf string
	to_p := fmt.Sprintf("\nSummary: %s\n%s\n", e.Summary, e.Description)
	buf += cleanString(to_p)
	buf += "\n"
	attachment_title := "\nAttachments:\n"
	attachment_entries := ""
	for _, a := range e.Attachments {
		if a == nil {
			continue
		}

		attachment_entries += fmt.Sprintf("- [[%s][%s]]\n", a.FileUrl,
			cleanString(a.Title))
	}

	if len(attachment_entries) > 0 {
		buf += attachment_title + attachment_entries
	}

	return buf
}

func fmtEventGroup(events []*calendar.Event) string {
	var buf string

	// take the last header of the set, has the most recent summary info.
	buf = fmtOrgHeader(events[len(events) - 1])

	// Put the dates from each event repeat
	unique_attendees := make(map[string]struct{})
	for _, i := range events {
		buf += fmtOrgDate(i)
		attendee := fmtOrgAttendees(i)
		if _, ok := unique_attendees[attendee]; !ok {
			unique_attendees[attendee] = struct{}{}
			buf += attendee
		}
	}

	unique_bodies := make(map[string]struct{})
	// Remove duplicate bodies
	for _, i := range events {
		body := fmtOrgBody(i)
		if _, ok := unique_bodies[body]; !ok {
			unique_bodies[body] = struct{}{}
			buf += body
		}
	}

	return buf
}

func filteredEvent(calid, summary string) bool {
	for _, title := range titleFilters[calid] {
		if strings.Contains(summary, title) {
			return true
		}
	}
	return false
}
