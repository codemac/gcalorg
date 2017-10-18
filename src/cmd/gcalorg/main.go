package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"google.golang.org/api/calendar/v3"
)

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(filename string, ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile(filename)
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile(filename string) (string, error) {
	tokname := filepath.Base(filename)
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape(fmt.Sprintf("calendar-api-quickstart.%s.json", tokname))), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func printOrgDate(start, end *calendar.EventDateTime) string {
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

// cleanString removes special characters for org-mode, as almost no one will be
// using org-mode formatting.
func cleanString(s string) string {
	s = strings.Replace(s, "[", "{", -1)
	s = strings.Replace(s, "]", "}", -1)
	s = strings.Replace(s, "\n*", "\n,*", -1)
	return s
}

func printOrg(e *calendar.Event) {
	var fullentry string
	print_entry := true
	fullentry += fmt.Sprintf("** ")
	if e.Status == "tenative" || e.Status == "cancelled" {
		fullentry += fmt.Sprintf("(%s) ", e.Status)
	}
	summary := e.Summary
	if summary == "" {
		summary = "busy"
	}
	fullentry += fmt.Sprintf("%s\n", summary)
	fullentry += fmt.Sprintf("   :PROPERTIES:\n")
	fullentry += fmt.Sprintf("   :ID:       %s\n", e.ICalUID)
	fullentry += fmt.Sprintf("   :GCALLINK: %s\n", e.HtmlLink)
	if e.Creator != nil {
		fullentry += fmt.Sprintf("   :CREATOR: [[mailto:%s][%s]]\n", e.Creator.Email, cleanString(e.Creator.DisplayName))
	}
	if e.Organizer != nil {
		fullentry += fmt.Sprintf("   :ORGANIZER: [[mailto:%s][%s]]\n", e.Organizer.Email, cleanString(e.Organizer.DisplayName))
	}
	fullentry += fmt.Sprintf("   :END:\n\n")
	fullentry += fmt.Sprintf("%s\n", printOrgDate(e.Start, e.End))
	attendees := e.Attendees
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
	if len(attendees) > 0 {
		fullentry += fmt.Sprintf("Attendees:\n")
	}
	if len(attendees) > 20 {
		fullentry += fmt.Sprintf("... Many\n")
	} else {
		for _, a := range attendees {
			if a != nil {

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
				fullentry += fmt.Sprintf(" %s [[mailto:%s][%s]]\n", statuschar, a.Email, linkname)

				// If the entire thing is actually declined, why
				// the fuck does google show it to me? this is
				// the most bullshit aspect of this calendar API
				// afaict. I really hope I've found the wrong
				// way of doing this.
				if a.Self && a.ResponseStatus == "declined" {
					print_entry = false
				}

			}
		}
	}

	to_p := fmt.Sprintf("\n%s\n", e.Description)
	esc_desc := cleanString(to_p)
	fullentry += fmt.Sprintf(esc_desc)
	fullentry += fmt.Sprintf("\n")
	fullentry += fmt.Sprintf("\nAttachments:\n")
	for _, a := range e.Attachments {
		if a == nil {
			continue
		}

		fullentry += fmt.Sprintf("- [[%s][%s]]\n", a.FileUrl, cleanString(a.Title))
	}
	if print_entry {
		fmt.Printf(fullentry)
	}
}

func printCalendars(client *http.Client, approvedCals []string) {

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
		fmt.Printf("* %s\n", c.Summary)
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

		// sort events by Id
		sort.SliceStable(event_list, func(i, j int) bool { return event_list[i].Id < event_list[j].Id })

	itemloop:
		for _, i := range event_list {
			// If the DateTime is an empty string the Event is an
			// all-day Event.  So only Date is available.

			// skip things that are chatty (repeating calendar
			// events -> org-mode has been difficult, manually
			// manage those for now. There is probably a way of
			// getting them, but converting the ical format to the
			// org format would be a significant piece of logic)
			for _, v := range titleFilters[c.Id] {
				if strings.Contains(i.Summary, v) {
					continue itemloop
				}
			}
			printOrg(i)
		}
	}
}

func genClient(filename string) *http.Client {
	ctx := context.Background()

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	return getClient(filename, ctx, config)
}

func main() {
	home := os.Getenv("HOME")
	type caldata struct {
		name string
		cals []string
	}
	secrets := []caldata{
		{home + "/src/gcalorg/jmickeygoogle_secret.json", workCals},
		{home + "/src/gcalorg/codemacgmail_secret.json", gmailCals},
	}

	// we need to sort before we do much of anything, so things show up in a
	// decent order.
	for _, v := range secrets {
		fmt.Fprintf(os.Stderr, "Getting client for: %s", v.name)
		cl := genClient(v.name)
		printCalendars(cl, v.cals)
	}
}
