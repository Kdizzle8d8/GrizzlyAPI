package calendars

import (
	"fmt"
	"time"
)
type dateResult struct {
	Date string `json:"date"`
	Type string `json:"type"`
}

func GetTypeFromDate(month int, day int) dateResult {
	c := NewCalendar()

	for _, m := range c.Months {
		if m.Number == month {
			for _, d := range m.Days {
				if d.Number == day {
					return dateResult{Date: fmt.Sprintf("%d/%d", month, day), Type: d.Type} 
				}
			}
		}
	}
	return dateResult{} // Return empty dateResult if no match is found
}
func GetTypeFromToday() dateResult  {
	day:=time.Now().Day()
	month:=int(time.Now().Month())
	return GetTypeFromDate(month,day)
}