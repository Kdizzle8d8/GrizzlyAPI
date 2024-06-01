package main

import (
	"os"
	"testing"

	"github.com/Kdizzle8d8/GrizzlyAPI/calendars"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	data, jsonErr := os.ReadFile("calendar.json")
	if jsonErr != nil {
		assert.Failf(t, "Error reading file: %v", jsonErr.Error())
	}

	calendar, err := calendars.LoadCalendar(data)
	assert.NoErrorf(t, err, "Error creating calendar: %v", err)

	calendarJSON, err := calendar.ToJSON()
	assert.NoErrorf(t, err, "Error converting calendar to JSON: %v", err)

	//0777 is the file permissions, similar to chmod in the terminal
	//0 means the number is in octal, and the 3 digits represent the permissions for the owner, group, and others
	//1 is execute, 2 is write, and 4 is read, so 7 is read, write, and execute (1+2+4=7)
	os.WriteFile("output.json", calendarJSON, 0777)
}
