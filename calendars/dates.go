package calendars

import (
	"fmt"
	"time"
)

type dateResult struct {
	Date string `json:"date"`
	Type string `json:"type"`
}

func GetTypeFromDate(month int, day int) (dateResult, error) {
	c := NewCalendar()

	for _, m := range c.Months {
		if m.Number == month {
			for _, d := range m.Days {
				if d.Number == day {
					return dateResult{Date: fmt.Sprintf("%d/%d", month, day), Type: d.Type}, nil
				}
			}
		}
	}
	return dateResult{}, fmt.Errorf("date not found")
}
func GetTypeFromToday() (dateResult, error) {
	day := time.Now().Day()
	month := int(time.Now().Month())
	return GetTypeFromDate(month, day)
}
