package main

import "time"

func main() {
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		{
			println("Weekend")
		}
	default:
		{
			println("weekday")
		}
	}
}
