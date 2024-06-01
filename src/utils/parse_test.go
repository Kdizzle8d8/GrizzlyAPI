package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	//In golang alot of functions return 2 values to indicate an error vs in javascript where they automatically throw an error
	jsonData, jsonErr := os.ReadFile("calendar.json")
	//Manual error checking
	if jsonErr != nil {
		fmt.Println("Error:", jsonErr)
		return
	}

	//We can't unmarshal the JSON data directly into our Calendar struct because the JSON data is structured differently
	//We could make a new struct that matches the JSON data, but that seems unnecary for this one use case
	var tempCalendar map[string]map[string]string
	//The first parameter is the JSON to parse, and the second is a pointer to the variable to store the parsed data
	//The default behavior in Go is to pass a copy of a variable to a function, so if we want to modify the original variable, we need to pass a pointer which is why we use the & operator
	err := json.Unmarshal(jsonData, &tempCalendar)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//This is our empty struct that we will populate with the converted JSON data
	var calendar Calendar

	//this is basically saying for(month in tempCalendar) and then assinging days=month.days and monthName=month.name

	calendarJSON, err := json.Marshal(calendar)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//0777 is the file permissions, similar to chmod in the terminal
	//0 means the number is in octal, and the 3 digits represent the permissions for the owner, group, and others
	//1 is execute, 2 is write, and 4 is read, so 7 is read, write, and execute (1+2+4=7)
	os.WriteFile("output.json", calendarJSON, 0777)
}
