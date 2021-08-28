package main

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/apognu/gocal"
)

func icsCals(dirs []string) []*Calendar {
	cals := make([]*Calendar, 0, len(dirs))
	for _, d := range dirs {
		cals = append(cals, NewCalendarFromIcs(d, ""))
	}

	return cals
}

func icsEvents(c *Calendar, timeMin time.Time, timeMax time.Time) []*Event {
	path := c.Id
	paths := icsWalkFiles(path)
	icsevents := icsParseAll(paths, timeMin, timeMax)

	events := make([]*Event, 0, len(icsevents))
	for _, e := range icsevents {
		events = append(events, NewEventFromIcs(&e))
	}

	return events
}

// Given a `dir` directory, return all .ics file locations
func icsWalkFiles(dir string) []string {
	paths := make([]string, 0, 100)
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			// continue
			return nil
		}
		if strings.HasSuffix(d.Name(), "ics") {
			paths = append(paths, path)
			return nil
		}
		// keep searching
		return nil
	})
	return paths
}

func icsParseAll(files []string, timeMin time.Time, timeMax time.Time) []gocal.Event {
	var fds []io.Reader
	var rawFds []*os.File
	for _, f := range files {
		f, _ := os.Open(f)
		fds = append(fds, f)
		rawFds = append(rawFds, f)
	}
	r := io.MultiReader(fds...)

	c := gocal.NewParser(r)
	c.Start = &timeMin
	c.End = &timeMax
	c.Parse()
	return c.Events
}
