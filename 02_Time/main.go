package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("lets work on time")

	presentime := time.Now()
	fmt.Println(presentime)

	newPresentTime := presentime.Format("02-01-2006 Monday") //month-date-year
	fmt.Println(newPresentTime)

	createdDate := time.Date(2023, time.July, 26, 23, 15, 15, 15, time.UTC)
	fmt.Println(createdDate)
	//we can also the fromat date in created date also to format the date according to our needs.
}
