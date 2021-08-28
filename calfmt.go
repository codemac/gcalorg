package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func datesToOrg(start, end *time.Time) string {
	if start == nil { // this event has dates! hurrah!
		return "\n"
	}

	//	debug := fmt.Sprintf("s: %s e: %s Δ: %s", start.Format("2006-01-02 Mon 15:04"), end.Format("2006-01-02 Mon 15:04"), end.Sub(*start))
	debug := ""
	// Figuring out all day events:
	// - see that they start & end at 00:00 and 23:59 or 00:00
	if start.Hour() == 0 && start.Minute() == 0 && end.Hour() == 23 && end.Minute() == 59 {
		delta := end.Sub(*start)
		if delta <= time.Duration(time.Hour*24) && delta > (time.Duration(time.Hour*23)+time.Duration(time.Minute*59)) {
			return fmt.Sprintf("<%s>", start.Format("2006-01-02"))
		}

		// else this is whole days, just not 1 day
		return fmt.Sprintf("<%s>--<%s>", start.Format("2006-01-02"), end.Format("2006-01-02")) + " " + debug
	}

	final := fmt.Sprintf("<%s", start.Format("2006-01-02 Mon 15:04"))

	if end == nil {
		return final + "> " + debug
	}

	if end.Day() != start.Day() {
		return final + ">--<" + end.Format("2006-01-02 Mon 15:04") + "> " + debug
	}

	return final + "-" + end.Format("15:04") + "> " + debug
}

func datesToInactiveOrg(start, end *time.Time) string {
	datestr := datesToOrg(start, end)
	left := strings.ReplaceAll(datestr, "<", "[")
	return strings.ReplaceAll(left, ">", "]")
}

func noTodoKwds(s string) string {
	states := []string{"TODO", "NEXT", "STARTED", "WAITING", "PROJECT", "DONE", "NVM"}
	for _, state := range states {
		// We essentially reimplement TrimPrefix, so we don't waste time
		// running HasPrefix multiple times. We also don't remove the
		// space in the prefix.
		if strings.HasPrefix(s, state+" ") {
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

func fmtOrgHeader(e *Event) string {
	var buf string
	buf += fmt.Sprintf("** ")
	summary := e.Name
	if summary == "" {
		summary = "busy"
	}

	buf += fmt.Sprintf("%s\n", noTodoKwds(summary))
	buf += fmt.Sprintf(":PROPERTIES:\n")
	buf += fmt.Sprintf(":ID:       %s\n", e.UID)
	buf += fmt.Sprintf(":GCALLINK: %s\n", e.URL)
	buf += fmt.Sprintf(":ORGANIZER: [[mailto:%s][%s]]\n", e.OrganizerMail, cleanString(e.OrganizerMail))
	buf += fmt.Sprintf(":END:\n\n")

	return buf
}

func fmtOrgDate(e *Event) string {
	return fmt.Sprintf("%s\n", datesToOrg(e.Start, e.End))
}

func fmtInactiveOrgDate(e *Event) string {
	return fmt.Sprintf("%s\n", datesToInactiveOrg(e.Start, e.End))
}

func fmtOrgAttendees(e *Event) string {
	var buf string
	attendees := e.Attendees
	if len(attendees) == 0 {
		return ""
	}

	canonical_id := func(ea *Attendee) string {
		if ea.Email != "" {
			return ea.Email
		} else if ea.Name != "" {
			return cleanString(ea.Name)
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

		statuschar := " "
		switch a.ResponseStatus {
		case ResponseDeclined:
			statuschar = "✗"
		case ResponseTentative:
			statuschar = "☐"
		case ResponseAccepted:
			statuschar = "✓"
		}

		linkname := cleanString(a.Name)
		if linkname == "" {
			linkname = a.Email
		}
		buf += fmt.Sprintf(" %s [[mailto:%s][%s]]\n", statuschar, a.Email, linkname)
	}
	return buf
}

func fmtOrgBody(e *Event) string {
	var buf string
	to_p := fmt.Sprintf("\nSummary: %s\n%s\n", e.Name, e.Description)
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

func fmtEventGroup(calid string, events []*Event) string {
	var buf string

	// take the last header of the set, has the most recent summary info.
	buf = fmtOrgHeader(events[len(events)-1])

	// Put the dates from each event repeat
	unique_attendees := make(map[string]struct{})
	unique_dates := make(map[string]struct{})
	for _, i := range events {
		if attendingEvent(calid, *i) {
			formatted_date := fmtOrgDate(i)
			if _, ok := unique_dates[formatted_date]; !ok {
				unique_dates[formatted_date] = struct{}{}
				buf += formatted_date
			}
		} else {
			buf += fmtInactiveOrgDate(i)
		}
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

func attendingEvent(calid string, event Event) bool {
	for _, email := range attendeeFilters[calid] {
		for _, att := range event.Attendees {
			if (att.Email == email ||
				att.Name == email) &&
				att.ResponseStatus == ResponseDeclined {
				return false
			}
		}
	}
	// include by default rather than not include.
	return true
}

func fmtEvents(c *Calendar, events []*Event) {
	events_by_uid := make(map[string][]*Event)
	for _, v := range events {
		recur_id := strings.Split(v.UID, "_R")[0]
		events_by_uid[recur_id] = append(events_by_uid[recur_id], v)
	}

	type eventWithUid struct {
		uid    string
		events []*Event
	}

	// sorted events
	sorted_by_uid := make([]eventWithUid, 0, len(events_by_uid))
	for uid, events := range events_by_uid {
		sorted_by_uid = append(sorted_by_uid,
			eventWithUid{uid, events})
	}

	sort.Slice(sorted_by_uid, func(i, j int) bool {
		return sorted_by_uid[i].uid < sorted_by_uid[j].uid
	})

	for _, e := range sorted_by_uid {
		events := e.events
		if len(events) == 0 {
			continue
		}
		// skip things that are chatty (repeating calendar
		// events -> org-mode has been difficult, manually
		// manage those for now. There is probably a way of
		// getting them, but converting the ical format to the
		// org format would be a significant piece of logic)
		if filteredEvent(c.Id, events[0].Name) {
			continue
		}

		fmt.Println(fmtEventGroup(c.Id, events))
	}
}

func fmtCalendar(c *Calendar) {
	fmt.Printf("* %s :%s:\n", noTodoKwds(c.Name), c.Tag)
	fmt.Printf("  :PROPERTIES:\n")
	fmt.Printf("  :ID:         %s\n", c.Id)
	fmt.Printf("  :END:\n")
	fmt.Printf("\n%s\n\n", c.Description)
}

func fmtFile() {
	fmt.Printf("# -*- eval: (auto-revert-mode 1); -*-\n")
	fmt.Printf("#+category: cal\n")
}
