package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

var (
	icalBuddyBullet = "*icalBuddy-bullet*"
)

func icbSplitToEventStrings(buf string) []string {
	// find each "*icalbuddy-bullet* string and split on that
	res := strings.Split(buf, icalBuddyBullet)
	return res
}

func icalBuddyEvents(timeMin, timeMax time.Time) []*Event {
	time_fmt := "2006-01-02"
	time_min_str := timeMin.Format(time_fmt)
	time_max_str := timeMax.Format(time_fmt)
	e := exec.Command("icalBuddy",
		"-tf", "%H:%M",
		"-df", "%Y-%m-%d",
		"-b", icalBuddyBullet,
		"-ab", "*icalBuddy-alertBullet",
		"-ss", "\n*icalBuddy-sectionSeparator*\n",
		"-ps", "|\n|",
		"-nrd",
		"eventsFrom:"+time_min_str, "to:"+time_max_str)
	bout, err := e.CombinedOutput()
	if err != nil {
		panic(err)
	}
	out := string(bout[:])

	// now we convert it!
	event_strs := icbSplitToEventStrings(out)

	for _, s := range event_strs {
		icbStrToEvent(s)
	}
	return make([]*Event, 0)
}

func icbStrToEvent(src string) *Event {
	// the icb format is pretty straight forward.
	// event title (calendar name)
	// properties
	// date

	by_line := strings.Split(str, "\n")
	
	// first strip out calendar from first line
	title_src := by_line[0]
	calendar_idx := strings.LastIndex(title_src)
	title := strings.Trim(title_src[:calendar_idx])
	calendar := title_src[calendar_idx+1:len(calendar_idx)-1]

	date_str := by_line[len(by_line)]

	notes := strings.Join(by_line[1:len(by_line)-1], "\n")


	return &Event{
		UID: "random",
		Name: title,
		Description: 
		
	
}
