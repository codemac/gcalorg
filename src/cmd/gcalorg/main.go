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
			return final + fmt.Sprintf(">")
		}

		te, _ := time.Parse("2006-01-02", end.Date)
		tef := te.Format("2006-01-02")
		return final + fmt.Sprintf(">--<%s>", tef)
	}

	ts, _ := time.Parse(time.RFC3339, start.DateTime)
	tsf := ts.Format("2006-01-02 Mon 15:04")
	final = final + fmt.Sprintf("<%s", tsf)

	if end == nil {
		return final + fmt.Sprintf(">")
	}

	te, _ := time.Parse(time.RFC3339, end.DateTime)
	if te.Day() != ts.Day() { // event spans days
		tef := te.Format("2006-01-02 Mon 15:04")
		return final + fmt.Sprintf(">--<%s>", tef)
	}

	tef := te.Format("15:04")
	return final + fmt.Sprintf("-%s>", tef)
}
func printOrg(e *calendar.Event) {
	fmt.Printf("** ")
	if e.Status == "tenative" || e.Status == "cancelled" {
		fmt.Printf("(%s) ", e.Status)
	}
	fmt.Printf("%s\n", e.Summary)
	fmt.Printf("   :PROPERTIES:\n")
	fmt.Printf("   :ID:       %s\n", e.ICalUID)
	fmt.Printf("   :GCALLINK: %s\n", e.HtmlLink)
	if e.Creator != nil {
		fmt.Printf("   :CREATOR: [[mailto:%s][%s]]\n", e.Creator.Email, e.Creator.DisplayName)
	}
	if e.Organizer != nil {
		fmt.Printf("   :ORGANIZER: [[mailto:%s][%s]]\n", e.Organizer.Email, e.Organizer.DisplayName)
	}
	fmt.Printf("   :END:\n")
	fmt.Printf("\n")
	fmt.Printf("%s\n", printOrgDate(e.Start, e.End))
	if len(e.Attendees) > 0 {
		fmt.Printf("Attendees:\n")
	}
	for _, a := range e.Attendees {
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

			fmt.Printf(" %s [[mailto:%s][%s]]\n", statuschar, a.Email, a.DisplayName)
		}

	}
	fmt.Printf("\n%s\n", e.Description)
	fmt.Printf("\n")
}

func PrintCalendars(client *http.Client, approved_cals map[string]struct{}) {

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
	timeMin := curtime.AddDate(0, -1, 0).Format("2006-01-02T15:04:05Z")
	timeMax := curtime.AddDate(1, 0, 0).Format("2006-01-02T15:04:05Z")
	for _, c := range calendars.Items {

		// this is a map[string]struct{} to check for
		// calendars to print. Remove this or add your own
		// secrets.go in the same package with your "approved
		// calendars" Id's to use this.
		if _, ok := approved_cals[c.Id]; !ok {
			continue
		}
		fmt.Printf("* %s\n", c.Summary)
		fmt.Printf("  :PROPERTIES:\n")
		fmt.Printf("  :ID:         %s\n", c.Id)
		fmt.Printf("  :END:\n")
		fmt.Printf("\n%s\n\n", c.Description)

		npt := ""
		notdone := true

		for notdone {
			events_notdone := srv.Events.List(c.Id).ShowDeleted(false).
				SingleEvents(true).TimeMin(timeMin).TimeMax(timeMax).MaxResults(250)
			if npt != "" {
				events_notdone = events_notdone.PageToken(npt)
				npt = ""
			}

			events, err := events_notdone.Do()
			if err != nil {
				log.Fatalf("Unable to retrieve next ten of the user's events. %v", err)
			}

			notdone = events.NextPageToken != ""
			if notdone {
				npt = events.NextPageToken
			}

			for _, i := range events.Items {
				// If the DateTime is an empty string the Event is an all-day Event.
				// So only Date is available.
				printOrg(i)
			}
		}
	}
}

func genClient(file_secrets string) *http.Client {
	ctx := context.Background()

	b, err := ioutil.ReadFile(file_secrets)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	return getClient(file_secrets, ctx, config)
}

func main() {
	secrets := map[string]map[string]struct{}{
		"/home/codemac/code/gcalorg/codemacgmail_secret.json": gmail_approved_cals,
		"/home/codemac/code/gcalorg/igneous_secret.json":      igneous_approved_cals,
	}

	for k, v := range secrets {
		cl := genClient(k)
		PrintCalendars(cl, v)
	}
}
