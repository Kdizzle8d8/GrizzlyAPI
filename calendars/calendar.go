package calendars

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Define format of JSON data
type Month struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
	Days   []Day  `json:"days"`
}

type Day struct {
	Number int    `json:"number"`
	Type   string `json:"type"`
}

type Calendar struct {
	Months []Month `json:"months"`
}

// This method is defined on the Calendar struct, so it can be called on any Calendar object.
func (c *Calendar) ToJSON() ([]byte, error) {

	sort.Slice(c.Months, func(i, j int) bool {
		if c.Months[i].Number > 8 && c.Months[j].Number <= 8 {
			return true
		} else if c.Months[i].Number <= 8 && c.Months[j].Number > 8 {
			return false
		} else {
			return c.Months[i].Number < c.Months[j].Number
		}
	})

	for i := range c.Months {
		sort.Slice(c.Months[i].Days, func(j, k int) bool {
			return c.Months[i].Days[j].Number < c.Months[i].Days[k].Number
		})
	}

	return json.Marshal(c)
}

// Define map of month names to month numbers
// Maps work like a mix of an array and an object in JavaScript
// map[keyType]valueType, so if I wanted to see the number for january in this map, I would do monthNumbers["january"]
var monthNumbers = map[string]int{
	"january":   1,
	"february":  2,
	"march":     3,
	"april":     4,
	"may":       5,
	"june":      6,
	"july":      7,
	"august":    8,
	"september": 9,
	"october":   10,
	"november":  11,
	"december":  12,
}

func LoadCalendar(data []byte) (Calendar, error) {
	var tempCalendar map[string]map[string]string

	err := json.Unmarshal(data, &tempCalendar)
	if err != nil {
		fmt.Println("Error:", err)
		return Calendar{}, err
	}

	return ParseCalendar(tempCalendar)
}

func ParseCalendar(data map[string]map[string]string) (Calendar, error) {
	calendar := Calendar{}
	for monthName, days := range data {
		//Strings don't have methods by default without importing a package, so we need to use the strings package for split and tolower
		monthKey := strings.ToLower(strings.Split(monthName, " ")[0])
		fmt.Println(strings.Split(monthName, " ")[0])

		month := Month{
			Name:   monthName,
			Number: monthNumbers[monthKey],
		}

		for day, event := range days {
			//strconv is a package that allows us to convert strings to numbers similar to parseInt in JavaScript
			dayNumber, _ := strconv.Atoi(day)
			dayStruct := Day{
				Number: dayNumber,
				Type:   event,
			}
			//Once again golang tends to stay away from methods on objects, so we use the append function with a target array and the value to append
			month.Days = append(month.Days, dayStruct)
		}

		calendar.Months = append(calendar.Months, month)
	}

	return calendar, nil
}
func NewCalendar() Calendar {
	data,err:=os.ReadFile("output.json")
	if err!=nil{
		fmt.Print("Error:",err)
	}
	var calendar Calendar
	err = json.Unmarshal(data, &calendar)
	if err!=nil{
		fmt.Print("Error:",err)
	}
	return calendar
}
