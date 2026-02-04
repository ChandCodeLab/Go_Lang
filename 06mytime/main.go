package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time study golang")

	presentTime := time.Now()
	fmt.Println("Present time is:", presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("Monday")) // it's fixed date
	fmt.Println(presentTime.Format("15:04:05"))

	//create a specific date
	createdDate := time.Date(2026, time.March, 10, 23, 3, 0, 0, time.UTC)
	fmt.Println("Created date is:", createdDate)

	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
