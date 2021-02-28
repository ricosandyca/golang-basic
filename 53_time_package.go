package main

import (
	"fmt"
	"time"
)

func main() {
	// get current time
	current := time.Now()
	year := current.Year()
	month := current.Month()
	day := current.Day()

	fmt.Println(current)
	fmt.Println(year, month, day)

	// create a new time
	newTime := time.Date(1998, time.November, 12, 0, 0, 0, 0, time.UTC)
	fmt.Println(newTime)

	// parsing time
	layout := "2006-01-02"
	parse, err := time.Parse(layout, "2020-11-12")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parse)
}
